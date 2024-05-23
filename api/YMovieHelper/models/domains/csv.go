package domains

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type CSVDomain struct {
	ruleRepo IRuleRepository
}

type IRuleRepository interface {
	GetDynamicItemNames(ruleID int) ([]string, error)
}

func NewCSVDomain(ruleRepo IRuleRepository) CSVDomain {
	domain := CSVDomain{ruleRepo: ruleRepo}
	return domain
}

func (domain CSVDomain) CreateCSV(ruleID int) (data io.Reader, err error) {
	dynamicItemNames, err := domain.ruleRepo.GetDynamicItemNames(ruleID)
	if err != nil {
		return nil, fmt.Errorf("CSVDomain.MakeCSV: %v", err)
	}

	var buffer bytes.Buffer
	writer := csv.NewWriter(transform.NewWriter(&buffer, japanese.ShiftJIS.NewEncoder()))
	firstRow := []string{"キャラクターの名前", "セリフ", "表情"}
	firstRow = append(firstRow, dynamicItemNames...)

	err = writer.Write(firstRow)
	if err != nil {
		return nil, fmt.Errorf("CSVDomain.MakeCSV: %v", err)
	}

	writer.Flush()
	return strings.NewReader(buffer.String()), err
}

func (domain CSVDomain) ReadCSV(data io.Reader) (csvData [][]string, err error) {
	reader := csv.NewReader(data)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("CSVDomain.ReadCSV: %v", err)
	}
	return records, nil
}
