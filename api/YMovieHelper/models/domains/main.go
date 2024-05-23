package domains

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/itkmaingit/YMovieHelper/utils"
)

type Start struct {
	InsertPlace     *string `db:"insert_place"`     //Fixed
	AdjustmentValue *int    `db:"adjustment_value"` //Flexible
	CharacterName   *string `db:"name"`             //Flexible
}

type End struct {
	Length          *int `db:"length"`           //Fixed
	HowManyAheads   *int `db:"how_many_aheads"`  //Flexible
	AdjustmentValue *int `db:"adjustment_value"` //Flexible
}

type EmptyItem struct {
	Sentence      string `db:"sentence"`
	CharacterName string `db:"name"`
}

type DynamicItem struct {
	ItemPath string `db:"item_path_in_pc"`
	ItemUrl  string `db:"item_url"`
	Layer    int    `db:"layer"`
	Name     string `db:"name"`
	Item     *ItemOnYMMP
}

type SingleItem struct {
	StaticItemID int    `db:"id"`
	ItemUrl      string `db:"item_path"`
	Layer        int    `db:"layer"`
	IsFixedStart bool   `db:"is_fixed_start"`
	IsFixedEnd   bool   `db:"is_fixed_end"`
	Start        Start
	End          End
	Item         *ItemOnYMMP
}

type MultipleItem struct {
	StaticItemID int    `db:"id"`
	ItemUrl      string `db:"item_path"`
	Layer        int    `db:"layer"`
	IsFixedStart bool   `db:"is_fixed_start"`
	Start        Start
	Items        []*ItemOnYMMP
}

type Rules struct {
	VoicelineLayer int
	MovieName      string
	EmotionData    map[string]map[string]*TachieFaceParameter
	EmptyItems     []EmptyItem
	DynamicItems   []DynamicItem
	SingleItems    []SingleItem
	MultipleItems  []MultipleItem
}

func CreateYMMP(data *YMMP, csvData *[][]string, rules *Rules) (*YMMP, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	errc := make(chan error)

	var voicelineRow []string
	for _, row := range (*csvData)[1:] {
		voicelineRow = append(voicelineRow, row[0])
	}

	voiceline := &data.Timeline.Items
	var createdDynamicItems, createdSingleItems, createdMultipleItems []*ItemOnYMMP

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, item := range rules.DynamicItems {
			var dynamicItemColumn int
			var dynamicItemRow []string

			//Dynamic Itemの列が何番目かを確認
		innerFor:
			for index, headerName := range (*csvData)[0] {
				if headerName == item.Name {
					dynamicItemColumn = index
					break innerFor
				}
			}

			for _, row := range (*csvData)[1:] {
				dynamicItemRow = append(dynamicItemRow, row[dynamicItemColumn])
			}

			items, err := CreateDynamicItems(*voiceline, &item, dynamicItemRow, rules.MovieName)
			if err != nil {
				errc <- fmt.Errorf("goroutine: domains.CreateYMMP in create dynamic items: %w", err)
				return
			}
			mu.Lock()
			createdDynamicItems = append(createdDynamicItems, items...)
			mu.Unlock()
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, item := range rules.SingleItems {
			items, err := CreateSingleItems(*voiceline, &item, voicelineRow)
			if err != nil {
				errc <- fmt.Errorf("goroutine: domains.CreateYMMP in create single items: %w", err)
				return
			}
			mu.Lock()
			createdSingleItems = append(createdSingleItems, items...)
			mu.Unlock()
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, item := range rules.MultipleItems {
			items, err := CreateMultipleItems(*voiceline, &item, voicelineRow)
			if err != nil {
				errc <- fmt.Errorf("goroutine: domains.CreateYMMP in create multiple items: %w", err)
				return
			}
			mu.Lock()
			createdMultipleItems = append(createdMultipleItems, items...)
			mu.Unlock()
		}

	}()

	go func() {
		wg.Wait()
		close(errc)
	}()

	var errMsg strings.Builder
	for err := range errc {
		if err != nil {
			errMsg.WriteString(err.Error() + "\n")
		}
	}

	if errMsg.Len() > 0 {
		return nil, errors.New(errMsg.String())
	}

	var voicelineEmotionCSV [][]string
	for _, row := range (*csvData)[1:] {
		selectedRow := []string{row[0], row[2]}
		voicelineEmotionCSV = append(voicelineEmotionCSV, selectedRow)
	}

	AttachEmotions(*voiceline, voicelineEmotionCSV, rules.EmotionData)
	DeleteEmptyItems(voiceline, rules.EmptyItems)
	MoveToVoicelineLayer(*voiceline, rules.VoicelineLayer)

	// 要素の追加なので、アドレスを渡している必要がある
	*voiceline = append(*voiceline, createdDynamicItems...)
	*voiceline = append(*voiceline, createdSingleItems...)
	*voiceline = append(*voiceline, createdMultipleItems...)

	return data, nil
}

func CreateScenarioTxt(csvData [][]string, emptyItems *[]EmptyItem) string {
	scenario := ""
OuterLoop:
	for _, row := range csvData[1:] {
		for _, item := range *emptyItems {
			if item.CharacterName == row[0] {
				serif := strings.Replace(strings.Replace(item.Sentence, "「", `\「`, -1), "」", `\」`, -1)
				scenario = fmt.Sprintf("%s%s「%s」\n", scenario, row[0], serif)
				continue OuterLoop
			}
		}
		serif := strings.Replace(strings.Replace(row[1], "「", `\「`, -1), "」", `\」`, -1)
		scenario = fmt.Sprintf("%s%s「%s」\n", scenario, row[0], serif)
	}
	return scenario
}

func DeleteEmptyItems(voiceline *[]*ItemOnYMMP, emtpyItems []EmptyItem) {
	j := 0
voiceline:
	for _, item := range *voiceline {

		for _, emptyItem := range emtpyItems {
			if *item.CharacterName == emptyItem.CharacterName {
				continue voiceline
			}
		}
		(*voiceline)[j] = item
		j++
	}
	*voiceline = (*voiceline)[:j]
}

func AttachEmotions(voiceline []*ItemOnYMMP, voicelineRow [][]string, emotionData map[string]map[string]*TachieFaceParameter) {
	for index, emotionCSV := range voicelineRow {
		if emotionCSV[1] == "" {
			continue
		}

		// Access the first level of map.
		if firstLevelMap, ok := emotionData[emotionCSV[0]]; ok {
			// Access the second level of map.
			if tachieFaceParameter, ok := firstLevelMap[emotionCSV[1]]; ok && tachieFaceParameter != nil {
				voiceline[index].TachieFaceParameter = tachieFaceParameter
			}
		}
	}
}

func MoveToVoicelineLayer(voiceline []*ItemOnYMMP, voicelineLayer int) {
	for index := range voiceline {
		voiceline[index].Layer = voicelineLayer
	}
}

// ボイスラインと、該当するDynamicItemの列を受け取ることでItemsを作成する
func CreateDynamicItems(voiceline []*ItemOnYMMP, dynamicItem *DynamicItem, dynamicItemRow []string, movieName string) ([]*ItemOnYMMP, error) {
	previousIndex := -1
	currentIndex := -1
	var previousFileName, currentFileName string
	var frame, length int
	var filePath string
	var ReturnItems []*ItemOnYMMP
	layer := dynamicItem.Layer
	lastFrame := GetLastFrame(voiceline)
	originalItem := dynamicItem.Item
	for i, fileName := range dynamicItemRow {
		if fileName != "" {
			// 最後の行は特別な操作を行う
			if i == len(dynamicItemRow)-1 {
				if currentFileName == "/" {
					currentIndex = i
					currentFileName = fileName
					frame = voiceline[currentIndex].Frame
					length = voiceline[currentIndex].Length
					filePath = utils.PathJoinForWindows(dynamicItem.ItemPath, movieName, dynamicItem.Name, currentFileName)
					newItem := NewDynamicItem(originalItem, frame, length, layer, filePath)
					ReturnItems = append(ReturnItems, newItem)
					continue

				} else if currentFileName != "/" {
					//通常の動作
					previousIndex = currentIndex
					currentIndex = i
					previousFileName = currentFileName
					currentFileName = fileName
					frame = voiceline[previousIndex].Frame
					length = voiceline[currentIndex].Frame - voiceline[previousIndex].Frame
					filePath = utils.PathJoinForWindows(dynamicItem.ItemPath, movieName, dynamicItem.Name, previousFileName)
					newItem := NewDynamicItem(originalItem, frame, length, layer, filePath)
					ReturnItems = append(ReturnItems, newItem)

					// 最後のアイテムは、最後のボイスアイテムの長さだけ獲得
					// ただし、最後のアイテムが/で終わっていれば、アイテムを挿入しない
					if currentFileName != "/" {
						frame = voiceline[currentIndex].Frame
						length = voiceline[currentIndex].Length
						filePath = utils.PathJoinForWindows(dynamicItem.ItemPath, movieName, dynamicItem.Name, currentFileName)
						newItem = NewDynamicItem(originalItem, frame, length, layer, filePath)
						ReturnItems = append(ReturnItems, newItem)
						continue
					}
					continue
				}
			}

			// 通常は、csvからまずファイル名がないかを走査する
			// ファイル名を見つけたらそれを一度保存しておく
			// 次に、「/」かファイル名を見つけた時、保存しておいたファイル名と、そのアイテムまでの距離を計算し、frameとlengthを作成する.
			previousIndex = currentIndex
			currentIndex = i
			previousFileName = currentFileName
			currentFileName = fileName
			if previousIndex == -1 {
				continue
			}
			if previousFileName != "/" {
				frame = voiceline[previousIndex].Frame
				length = voiceline[currentIndex].Frame - voiceline[previousIndex].Frame
				filePath = utils.PathJoinForWindows(dynamicItem.ItemPath, movieName, dynamicItem.Name, previousFileName)
				newItem := NewDynamicItem(originalItem, frame, length, layer, filePath)
				ReturnItems = append(ReturnItems, newItem)

			} else if previousFileName == "/" {
				if currentFileName != "/" {
					continue
				} else if currentFileName == "/" {
					return nil, errors.New("invalid csv file: '/' was followed twice in a row. ")
				}
			}

			// 基本的に、csvの中でファイル名が空の時はスキップする
			// しかし、最後の行の時は、直前にストックしておいたファイル名を最後のフレームまで伸ばす
		} else if fileName == "" {
			if i != len(dynamicItemRow)-1 {
				continue
			} else if i == len(dynamicItemRow)-1 {
				if currentFileName == "/" {
					continue
				} else if currentFileName != "" {
					frame = voiceline[currentIndex].Frame
					length = lastFrame - voiceline[currentIndex].Frame
					filePath = utils.PathJoinForWindows(dynamicItem.ItemPath, movieName, dynamicItem.Name, currentFileName)
					newItem := NewDynamicItem(originalItem, frame, length, layer, filePath)
					ReturnItems = append(ReturnItems, newItem)
				}
			}
		}
	}

	return ReturnItems, nil
}

func CreateSingleItems(voiceline []*ItemOnYMMP, singleItem *SingleItem, voicelineRow []string) ([]*ItemOnYMMP, error) {
	var frame, length int
	var ReturnItems []*ItemOnYMMP
	maxIndex := len(voiceline) - 1
	lastFrame := GetLastFrame(voiceline)
	layer := singleItem.Layer

	// ポインタなのでnilの可能性
	var howManyAheads, startAdjustmentValue, endAdjustmentValue int
	if singleItem.End.HowManyAheads != nil {
		howManyAheads = *singleItem.End.HowManyAheads
	}

	if singleItem.Start.AdjustmentValue != nil {
		startAdjustmentValue = *singleItem.Start.AdjustmentValue
	}

	if singleItem.End.AdjustmentValue != nil {
		endAdjustmentValue = *singleItem.End.AdjustmentValue
	}

	if singleItem.IsFixedStart {
		// 固定フレームなので先に取得
		if *singleItem.Start.InsertPlace == "最初" {
			frame = 0
		} else if *singleItem.Start.InsertPlace == "最後" {
			frame = lastFrame
		}

		// 固定長
		if singleItem.IsFixedEnd {
			length = *singleItem.End.Length
			appendItem := NewSingleItem(singleItem.Item, frame, length, layer)
			ReturnItems = append(ReturnItems, appendItem)

			// 動的長さの場合は、「最初」のhowmanyAheads&adjustmentValueしかありえない
		} else if !singleItem.IsFixedEnd {
			if howManyAheads <= maxIndex {
				length = endAdjustmentValue + (voiceline[howManyAheads].Frame)
				appendItem := NewSingleItem(singleItem.Item, frame, length, layer)
				ReturnItems = append(ReturnItems, appendItem)
			} else if (howManyAheads) > maxIndex {
				length = endAdjustmentValue + (lastFrame)
				appendItem := NewSingleItem(singleItem.Item, frame, length, layer)
				ReturnItems = append(ReturnItems, appendItem)
			}
		}

		return ReturnItems, nil

		// 動的フレームの場合はvoicelineを走査
	} else if !singleItem.IsFixedStart {
		for index, characterName := range voicelineRow {
			if characterName == *singleItem.Start.CharacterName {
				frame = voiceline[index].Frame + startAdjustmentValue

				if singleItem.IsFixedEnd {
					length = *singleItem.End.Length
					appendItem := NewSingleItem(singleItem.Item, frame, length, layer)
					ReturnItems = append(ReturnItems, appendItem)
					// 動的長さの時は、voicelineの長さによるバリデーションを行う、超えていたらLastFrame-現在のFrameにしておく(最後まで伸ばす)
				} else if !singleItem.IsFixedEnd {
					if howManyAheads+index <= maxIndex {
						length = endAdjustmentValue + (voiceline[index+howManyAheads].Frame - voiceline[index].Frame)
						appendItem := NewSingleItem(singleItem.Item, frame, length, layer)
						ReturnItems = append(ReturnItems, appendItem)
					} else if (howManyAheads + index) > maxIndex {
						length = endAdjustmentValue + (lastFrame - voiceline[index].Frame)
						appendItem := NewSingleItem(singleItem.Item, frame, length, layer)
						ReturnItems = append(ReturnItems, appendItem)
					}
				}
			}
		}
		return ReturnItems, nil
	}
	return ReturnItems, errors.New("domains.CreateSingleItems: This Single Item is neither IsFixedStart nor NotIsFixedStart.  ")
}

func CreateMultipleItems(voiceline []*ItemOnYMMP, multipleItem *MultipleItem, voicelineRow []string) ([]*ItemOnYMMP, error) {
	var frame int
	var ReturnItems []*ItemOnYMMP
	lastFrame := GetLastFrame(voiceline)
	layer := multipleItem.Layer
	var startAdjustmentValue int
	if multipleItem.Start.AdjustmentValue != nil {
		startAdjustmentValue = *multipleItem.Start.AdjustmentValue
	}

	if multipleItem.IsFixedStart {
		// 固定フレーム
		if *multipleItem.Start.InsertPlace == "最初" {
			frame = 0
		} else if *multipleItem.Start.InsertPlace == "最後" {
			frame = lastFrame
		}

		appendItem := NewMultipleItem(multipleItem.Items, frame, layer)
		ReturnItems = append(ReturnItems, appendItem...)
		return ReturnItems, nil

		// 動的フレームの場合はvoicelineを走査
	} else if !multipleItem.IsFixedStart {
		for index, characterName := range voicelineRow {
			if characterName == *multipleItem.Start.CharacterName {
				frame = voiceline[index].Frame + startAdjustmentValue

				appendItem := NewMultipleItem(multipleItem.Items, frame, layer)
				ReturnItems = append(ReturnItems, appendItem...)
			}
		}
		return ReturnItems, nil
	}
	return ReturnItems, errors.New("domains.CreateMultipleItems:  This Multiple Item is neither IsFixedStart nor NotIsFixedStart. ")
}

func GetLastFrame(voiceline []*ItemOnYMMP) int {
	item := voiceline[len(voiceline)-1]
	lastFrame := item.Frame + item.Length
	return lastFrame
}
