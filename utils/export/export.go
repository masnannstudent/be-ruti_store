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

func ExportXlsx(c *fiber.Ctx, data [][]interface{}, headers []string, title, companyName, companyAddress, dateRange, fileName string) error {
	file := excelize.NewFile()
	sheetName := "Sheet1"

	// Judul Laporan
	titleCell := fmt.Sprintf("A%d", 1)
	file.SetCellValue(sheetName, titleCell, title)
	titleStyle, _ := file.NewStyle(`{"font":{"bold":true,"size":16},"alignment":{"horizontal":"center"}}`)
	file.SetCellStyle(sheetName, titleCell, fmt.Sprintf("%c%d", 'A'+len(headers)-1, 1), titleStyle)
	file.MergeCell(sheetName, titleCell, fmt.Sprintf("%c%d", 'A'+len(headers)-1, 1))

	// Informasi Perusahaan (Satu baris)
	companyInfo := fmt.Sprintf("%s, %s", companyName, companyAddress)
	companyInfoCell := fmt.Sprintf("A%d", 3)
	file.SetCellValue(sheetName, companyInfoCell, companyInfo)
	companyInfoStyle, _ := file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center"}}`)
	file.SetCellStyle(sheetName, companyInfoCell, fmt.Sprintf("%c%d", 'A'+len(headers)-1, 3), companyInfoStyle)
	file.MergeCell(sheetName, companyInfoCell, fmt.Sprintf("%c%d", 'A'+len(headers)-1, 3))

	// Tanggal
	dateCell := fmt.Sprintf("A%d", 2)
	file.SetCellValue(sheetName, dateCell, dateRange)
	dateStyle, _ := file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center"}}`)
	file.SetCellStyle(sheetName, dateCell, fmt.Sprintf("%c%d", 'A'+len(headers)-1, 2), dateStyle)
	file.MergeCell(sheetName, dateCell, fmt.Sprintf("%c%d", 'A'+len(headers)-1, 2))

	// Header Kolom
	for col, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+col, 5)
		file.SetCellValue(sheetName, cell, header)
		headerStyle, _ := file.NewStyle(`{"alignment":{"horizontal":"center"},"fill":{"type":"pattern","color":["#ecf0f1"],"pattern":1}}`)
		file.SetCellStyle(sheetName, cell, cell, headerStyle)
	}

	// Data
	for row, rowData := range data {
		for col, value := range rowData {
			cell := fmt.Sprintf("%c%d", 'A'+col, row+6)
			file.SetCellValue(sheetName, cell, value)
			cellStyle, _ := file.NewStyle(`{"alignment":{"horizontal":"left"}}`)
			file.SetCellStyle(sheetName, cell, cell, cellStyle)
		}
	}

	// Set lebar kolom
	setColumnWidths(sheetName, file, headers, data)

	// Simpan ke file
	err := file.SaveAs(fileName)
	if err != nil {
		return err
	}

	// Set header response dan kirim file
	c.Set("Content-Disposition", "attachment; filename="+fileName)
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return c.SendFile(fileName)
}
