package xlsx

import (
	"github.com/tealeg/xlsx"
	"log"
	. "sheet-generator/generator"
)

type XlsxGenerator struct {
	fileName string
	file     *xlsx.File
	sheet    *xlsx.Sheet
}

func NewXlsxGenerator() *XlsxGenerator {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Fatalln("XlsxGenerator create sheet error:", err)
	}
	return &XlsxGenerator{
		fileName: FileNameGen("xlsx"),
		file:     file,
		sheet:    sheet,
	}
}

func (xlsx *XlsxGenerator) Name() string {
	return xlsx.fileName
}

func (xlsx *XlsxGenerator) Init() {
	row := xlsx.sheet.AddRow()
	row.WriteSlice(&FirstLine, len(FirstLine))
}

func (xlsx *XlsxGenerator) InsertRecord(record []string) {
	row := xlsx.sheet.AddRow()
	row.WriteSlice(&record, len(record))
}

func (xlsx *XlsxGenerator) Complete() {
	if err := xlsx.file.Save(xlsx.fileName); err != nil {
		log.Fatalln("XlsxGenerator save file error:", err)
	}
}
