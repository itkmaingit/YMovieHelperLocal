package models

import (
	"fmt"
)

type DownloadModel struct {
}

func NewDownloadModel() DownloadModel {
	model := DownloadModel{}
	return model
}

// ➀まず、ymmpを作成できるかどうかを確かめる
// ➁作成できて初めてtemplate.csvを作成し、S3にアップロードしておく
func (model DownloadModel) ConstructDownloadScenarioCSV(ruleID int) string {
	filePath := fmt.Sprintf("./temp/%d/template.csv", ruleID)
	return filePath
}
func (model DownloadModel) ConstructScenarioTXT(ruleID int) string {
	filePath := fmt.Sprintf("./temp/%d/scenario.txt", ruleID)
	return filePath
}
func (model DownloadModel) ConstructDownloadCompleteYMMP(ruleID int) string {
	filePath := fmt.Sprintf("./temp/%d/complete.ymmp", ruleID)
	return filePath
}
