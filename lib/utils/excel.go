package utils

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

func ReadExcel(fileName string, sheetsName string) ([]string, []map[string]string, error) {
	f, err := excelize.OpenFile(fileName)

	if err != nil {
		return nil, []map[string]string{}, err
	}
	defer f.Close()

	var (
		headers   []string
		finalData []map[string]string
	)
	rows, err := f.GetRows(sheetsName)
	if err != nil {
		return []string{}, []map[string]string{}, err
	}

	for rowIndex, row := range rows {
		tmpData := make(map[string]string)

		for colIndex, colCell := range row {
			if rowIndex == 0 {
				headers = append(headers, strings.TrimSpace(colCell))
			} else {
				if colIndex < len(headers) {
					tmpData[headers[colIndex]] = colCell
				}
			}

		}

		for i, d := range tmpData {
			if d == "" {
				delete(tmpData, i)
			}
		}

		if len(tmpData) != 0 {
			finalData = append(finalData, tmpData)
		}
	}
	return headers, finalData, nil
}
