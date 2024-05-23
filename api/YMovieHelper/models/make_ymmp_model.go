package models

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/utils"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type MakeYMMPModel struct {
	charaRepo ICharacterRepository
	ruleRepo  IRuleRepository
	ymmpRepo  IMakeYMMPRepository
	csvDomain domains.CSVDomain
}

type IMakeYMMPRepository interface {
	ExistsCharacter(ruleID int) (bool, error)
	GetDynamicItemsInRule(ruleID int) ([]domains.DynamicItem, error)
	GetSingleItemsInRule(ruleID int) ([]domains.SingleItem, error)
	GetMultipleItemsInRule(ruleID int) ([]domains.MultipleItem, error)
	GetEmptyItems(ruleID int) ([]domains.EmptyItem, error)
	GetCharacterNamesInRule(ruleID int) ([]string, error)
	GetEmotionMap(ruleID int) (map[string]map[string]string, error)
}

func NewMakeYMMPModel(charaRepo ICharacterRepository, ruleRepo IRuleRepository, ymmpRepo IMakeYMMPRepository, csvDomain domains.CSVDomain) MakeYMMPModel {
	model := MakeYMMPModel{charaRepo: charaRepo, ruleRepo: ruleRepo, ymmpRepo: ymmpRepo, csvDomain: csvDomain}
	return model
}

// ➀まず、ymmpを作成できるかどうかを確かめる
// ➁作成できて初めてtemplate.csvを作成し、S3にアップロードしておく
func (model MakeYMMPModel) CheckRules(ctx context.Context, softwareID int, ruleID int) (bool, error) {
	existsRule, err := model.ruleRepo.ExistsRule(ruleID)
	if err != nil {
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

	existsCharacter, err := model.ymmpRepo.ExistsCharacter(ruleID)
	if err != nil {
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

	// emptyItemとcharacterがルールとcharactersの定義と一致しているかを確かめる
	characters, err := model.charaRepo.GetCharacters(softwareID)
	if err != nil {
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

	characterIDs, err := model.ruleRepo.GetCharacterItemIDsInRule(ruleID)
	if err != nil {
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

characterLoop:
	for _, characterID := range characterIDs {
		for _, character := range characters {
			if character.ID == characterID {
				if !character.IsEmpty {
					continue characterLoop
				} else if character.IsEmpty {
					return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
				}
			}
		}
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

	emptyItemIDs, err := model.ruleRepo.GetEmptyItemIDsInRule(ruleID)
	if err != nil {
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

emptyItemLoop:
	for _, emptyItemID := range emptyItemIDs {
		for _, character := range characters {
			if character.ID == emptyItemID {
				if character.IsEmpty {
					continue emptyItemLoop
				} else if !character.IsEmpty {
					return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
				}
			}
		}
		return false, fmt.Errorf("MakeYMMPModel.CheckRules: %w", err)
	}

	csvData, err := model.csvDomain.CreateCSV(ruleID)
	if err != nil {
		return false, fmt.Errorf("RuleModel.CreateRule: %w", err)
	}
	csvPath := fmt.Sprintf("./temp/%d/template.csv", ruleID)
	err = utils.SaveFile(csvData, csvPath)
	if err != nil {
		return false, fmt.Errorf("RuleModel.UploadRule: %w", err)
	}

	return existsRule && existsCharacter, err
}

func (model MakeYMMPModel) ResolveScenario(ctx *gin.Context, csvData io.Reader, ruleID int) (string, error) {
	csvBytes, err := io.ReadAll(csvData)
	if err != nil {
		customError := &utils.CustomError{
			FrontMsg: "CSVの形式が不正です！",
			BackMsg:  fmt.Sprintf("models.MakeYMMPModel.ResolveScenario: %v", err),
		}
		return "", fmt.Errorf("ResolveScenario: %w", customError)
	}

	reader := csv.NewReader(transform.NewReader(bytes.NewReader(csvBytes), japanese.ShiftJIS.NewDecoder()))
	var records [][]string

	// データの格納
	for {
		record, err := reader.Read()
		// CSVの終わりに達した場合
		if err == io.EOF {
			break
		}
		// 別のエラーが発生した場合
		if err != nil {
			customError := &utils.CustomError{
				FrontMsg: "CSVの形式が不正です！",
				BackMsg:  fmt.Sprintf("models.MakeYMMPModel.ResolveScenario: %v", err),
			}
			return "", fmt.Errorf("ResolveScenario: %w", customError)
		}
		records = append(records, record)
	}

	transposeRecords := utils.Transpose(records)
	charaNames, err := model.ymmpRepo.GetCharacterNamesInRule(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.ResolveScenario: %w", err)
	}

	emptyItems, err := model.ymmpRepo.GetEmptyItems(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.ResolveScenario: %w", err)
	}

	dynamicItems, err := model.ymmpRepo.GetDynamicItemsInRule(ruleID)
	if err != nil {
		return "", fmt.Errorf("ResolveScenario:%w", err)
	}

	totalColumns := len(dynamicItems)
	// EmptyItemと合わせて、ruleの中に存在するキャラクターの名前を全て格納する
	for _, emptyItem := range emptyItems {
		charaNames = append(charaNames, emptyItem.CharacterName)
	}

	//DynamicItemの数 + CharacterName + Serif + Emotionがヘッダー行の列数と一致しないとエラー
	if len(records[0]) != totalColumns+3 {
		customError := &utils.CustomError{
			FrontMsg: "不正なCSVです！ Step1に戻って正常なCSVから台本を作成してください！",
			BackMsg:  fmt.Sprintf("models.MakeYMMPModel.ResolveScenario: %v", err),
		}
		return "", fmt.Errorf("ResolveScenario: %w", customError)
	}

	//CSVのキャラクター行にRuleに含まれていないキャラクターが入っているとエラー
characterLoop:
	for _, record := range records[1:] {
		for _, charaName := range charaNames {
			if record[0] == charaName {
				continue characterLoop
			}
		}
		err := &utils.CustomError{
			FrontMsg: "CSVのキャラクター行にRuleに含まれていないキャラクターが入っています！",
			BackMsg:  "models.MakeYMMPModel.ResolveScenario: Include invalid character (not included in rule) ",
		}
		return "", err
	}

	//セリフ行に空白が含まれるとき、それがEmptyItemでない時はエラーを発生させる
emptyLoop:
	for _, record := range records[1:] {
		if utils.IsEmptyOrWhitespace(record[1]) {
			for _, emptyItem := range emptyItems {
				if record[0] == emptyItem.CharacterName {
					continue emptyLoop
				}
			}
			customError := &utils.CustomError{
				FrontMsg: "セリフ行に空白の文章が入っています！",
				BackMsg:  fmt.Sprintf("models.MakeYMMPModel.ResolveScenario: %v", err),
			}
			return "", fmt.Errorf("ResolveScenario: %w", customError)
		}
	}

	// CSVが空白を挟んで2つ以上"/"が並んだらエラー
	for i := 3; i < len(records[0]); i++ {
		previousFileName := ""
		for _, fileName := range transposeRecords[i][1:] {
			if fileName == "/" {
				if previousFileName == "/" {
					customError := &utils.CustomError{
						FrontMsg: "CSV内のDynamic Itemにおいて、連続して「/」が並んでいます。",
						BackMsg:  fmt.Sprintf("models.MakeYMMPModel.ResolveScenario: %v", err),
					}
					return "", fmt.Errorf("ResolveScenario: %w", customError)
				}
				previousFileName = fileName

			} else if fileName != "" {
				previousFileName = fileName
			}

		}
	}

	// CSVのヘッダー行にて、含まれていないDynamicItemがあったらエラー
outerDynamicLoop:
	for _, record := range records[0][3:] {
		for _, dynamicItem := range dynamicItems {
			if record == dynamicItem.Name {
				continue outerDynamicLoop
			}
		}
		customError := &utils.CustomError{
			FrontMsg: "CSVのヘッダー行に、ルールに含まれていないDynamic Itemが含まれています！Step1に戻って正常なCSVを取得してください！",
			BackMsg:  fmt.Sprintf("models.MakeYMMPModel.ResolveScenario: %v", err),
		}
		return "", fmt.Errorf("ResolveScenario: %w", customError)
	}

	scenarioTxt := domains.CreateScenarioTxt(records, &emptyItems)
	fileTxtPath := fmt.Sprintf("./temp/%d/scenario.txt", ruleID)
	err = utils.SaveFile(bytes.NewReader([]byte(scenarioTxt)), fileTxtPath)
	if err != nil {
		return "", fmt.Errorf("ResolveScenario: %v", err)
	}

	// CSV形式にエンコードするためのバッファを作成
	buf := &bytes.Buffer{}
	csvWriter := csv.NewWriter(buf)

	// recordsの全てのデータをCSV形式にエンコード
	err = csvWriter.WriteAll(records)
	if err != nil {
		return "", fmt.Errorf("ResolveScenario:  %w", err)
	}

	fileCSVPath := fmt.Sprintf("./temp/scenario.csv")
	err = utils.SaveFile(bytes.NewReader(buf.Bytes()), fileCSVPath)
	if err != nil {
		return "", fmt.Errorf("ResolveScenario: %v", err)
	}

	return fileCSVPath, nil
}

func (model MakeYMMPModel) MakeYMMP(ctx context.Context, data *domains.YMMP, ruleID int, movieName string) (fileUrl string, err error) {
	// ファイルにボイスアイテムでないもの、あるいはVoiceCacheが含まれていたらエラー
	for _, item := range data.Timeline.Items {
		if *item.Type != "YukkuriMovieMaker.Project.Items.VoiceItem, YukkuriMovieMaker" {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイル内にボイスアイテムでないものが含まれているようです！",
				BackMsg:  "Include not VoiceItem ",
			}
			return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", customError)
		}
		if item.VoiceCache != nil {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイル内にボイスキャッシュが含まれているようです！（詳しくはHow toページをご覧ください。）",
				BackMsg:  "Include VoiceCache ",
			}
			return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", customError)
		}
	}

	csvPath := fmt.Sprintf("./temp/scenario.csv")

	// ShiftJISでエンコードされたCSVをダウンロード
	body, err := utils.ReadFile(csvPath)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	// CSVリーダーを作成
	csvReader := csv.NewReader(strings.NewReader(string(body)))

	// 全てのレコードを読み込む
	records, err := csvReader.ReadAll()
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	// タイムラインアイテムの数と、csvの行数が一致していなければエラー
	if len(data.Timeline.Items) != len(records)-1 {
		customError := &utils.CustomError{
			FrontMsg: "タイムラインアイテムの数とCSVの行数が一致していません！（scenario.txtを読み込んだymmpファイルをそのままアップロードしてください。）",
			BackMsg:  "The number of items != the number of csv columns",
		}
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", customError)
	}

	voicelineLayer, err := model.ruleRepo.GetVoicelineLayer(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	dynamicItems, err := model.ymmpRepo.GetDynamicItemsInRule(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	singleItems, err := model.ymmpRepo.GetSingleItemsInRule(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	multipleItems, err := model.ymmpRepo.GetMultipleItemsInRule(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	emptyItems, err := model.ymmpRepo.GetEmptyItems(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	emotionMap, err := model.ymmpRepo.GetEmotionMap(ruleID)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	for index := range dynamicItems {
		dynamicItems[index].FetchData()
		if err != nil {
			return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
		}
	}

	for index := range singleItems {
		err := singleItems[index].FetchData()
		if err != nil {
			return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
		}
	}

	for index := range multipleItems {
		err := multipleItems[index].FetchData()
		if err != nil {
			return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
		}
	}

	emotionData, err := domains.FetchEmotionData(emotionMap)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	rules := domains.Rules{
		VoicelineLayer: voicelineLayer,
		MovieName:      movieName,
		EmptyItems:     emptyItems,
		EmotionData:    emotionData,
		DynamicItems:   dynamicItems,
		SingleItems:    singleItems,
		MultipleItems:  multipleItems}

	data, err = domains.CreateYMMP(data, &records, &rules)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}
	ymmpFile, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	filePath := fmt.Sprintf("./temp/%d/complete.ymmp", ruleID)
	err = utils.SaveFile(bytes.NewReader(ymmpFile), filePath)
	if err != nil {
		return "", fmt.Errorf("MakeYMMPModel.MakeYMMP: %w", err)
	}

	return filePath, nil

}
