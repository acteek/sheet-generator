package csv

import (
	"encoding/csv"
	"log"
	"os"
	. "sheet-generator/generator"
)

// CsvGenerator - File Generator for CSV
type CsvGenerator struct {
	fileName string
	file     *os.File
	w        *csv.Writer
}

func NewCsvGenerator() *CsvGenerator {
	fileName := FileNameGen("csv")
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("CsvGenerator failed to open file:", err)
	}
	return &CsvGenerator{
		fileName: fileName,
		file:     file,
		w:        csv.NewWriter(file),
	}
}

func (csv *CsvGenerator) Name() string {
	return csv.fileName
}

func (csv *CsvGenerator) Init() {
	if err := csv.w.Write(FirstLine); err != nil {
		log.Fatalln("CsvGenerator error writing record to file", err)
	}
}

func (csv *CsvGenerator) InsertRecord(record []string) {
	if err := csv.w.Write(record); err != nil {
		log.Fatalln("CsvGenerator error writing record to file", err)
	}
}

func (csv *CsvGenerator) Complete() {
	csv.w.Flush()
	err := csv.file.Close()
	if err != nil {
		log.Println("CsvGenerator file close error:", err)
	}
}
