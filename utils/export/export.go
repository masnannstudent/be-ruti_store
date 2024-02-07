package export

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

func setColumnWidths(sheetName string, file *excelize.File, headers []string, data [][]interface{}) {
	for col, header := range headers {
		maxWidth := len(header)
		for _, rowData := range data {
			if col < len(rowData) {
				cellValue := fmt.Sprintf("%v", rowData[col])
				if len(cellValue) > maxWidth {
					maxWidth = len(cellValue)
				}
			}
		}

		file.SetColWidth(sheetName, fmt.Sprintf("%c", 'A'+col), fmt.Sprintf("%c", 'A'+col), float64(maxWidth+2)) // Add some extra space
	}
}

func ExportXlsx(c *fiber.Ctx, data [][]interface{}, headers []string, title, fileName string) error {
	file := excelize.NewFile()
	sheetName := "Sheet1"

	file.SetCellValue(sheetName, "A1", title)
	titleStyle, _ := file.NewStyle(`{"font":{"bold":true,"size":16},"alignment":{"horizontal":"center"}}`)
	file.SetCellStyle(sheetName, "A1", fmt.Sprintf("%c%d", 'A'+len(headers)-1, 1), titleStyle)
	file.MergeCell(sheetName, "A1", fmt.Sprintf("%c%d", 'A'+len(headers)-1, 1))

	for col, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+col, 3)
		file.SetCellValue(sheetName, cell, header)
		headerStyle, _ := file.NewStyle(`{"alignment":{"horizontal":"center"},"fill":{"type":"pattern","color":["#ecf0f1"],"pattern":1}}`)
		file.SetCellStyle(sheetName, cell, cell, headerStyle)
	}

	for row, rowData := range data {
		for col, value := range rowData {
			cell := fmt.Sprintf("%c%d", 'A'+col, row+4)
			file.SetCellValue(sheetName, cell, value)
			cellStyle, _ := file.NewStyle(`{"alignment":{"horizontal":"left"}`)
			file.SetCellStyle(sheetName, cell, cell, cellStyle)
		}
	}

	setColumnWidths(sheetName, file, headers, data)
	err := file.SaveAs(fileName)
	if err != nil {
		return err
	}

	c.Set("Content-Disposition", "attachment; filename="+fileName)
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return c.SendFile(fileName)
}
