package main

import (
	"flag"
	"github.com/icrowley/fake"
	"github.com/jaswdr/faker"
	"log"
	"math/rand"
	. "sheet-generator/generator"
	"sheet-generator/generator/csv"
	"sheet-generator/generator/xlsx"
	"strconv"
	"time"
)

func main() {

	var lines int
	var lang string
	var size string
	var format string
	flag.IntVar(&lines, "line", 10_000, "Count of lines")
	flag.StringVar(&lang, "lang", "ru", "Faker Language(ru, en)")
	flag.StringVar(&size, "size", "middle", "Size of file(low, middle, big)")
	flag.StringVar(&format, "format", "xlsx", "Format of file(xlsx, csv)")
	flag.Parse()

	unix := time.Now().Unix()
	fakeEN := faker.New()
	if err := fake.SetLang(lang); err != nil {
		log.Fatalf("Lang [%v] do not support!\n", lang)
	}

	var generator FileGenerator
	switch format {
	case "xlsx":
		generator = xlsx.NewXlsxGenerator()
	case "csv":
		generator = csv.NewCsvGenerator()
	default:
		log.Fatalf("Format [%v] is not supported!\n", format)
	}

	generator.Init()
	log.Println("Generation is started ...")
	for i := 0; i < lines; i++ {
		maxTime := int64(rand.Int31n(int32(unix)))
		var record = []string{
			fake.LastName(),
			fake.FirstName(),
			fake.Patronymic(),
			fakeEN.Phone().E164Number(),
			fakeEN.Internet().Email(),
			fakeEN.Internet().User(),
			fakeEN.Internet().User(),
			fake.JobTitle(),
			fake.Company(),
			strconv.Itoa(int(rand.Uint32())),
			time.Unix(maxTime, 0).Format("02.01.2006"),
			fake.ParagraphsN(comment(size)),
		}
		generator.InsertRecord(record)
		if i%1000 == 0 {
			log.Printf("Lines generated  %v\n", i)
		}

	}

	defer func(gen FileGenerator) {
		gen.Complete()
		fileName := gen.Name()
		size := sizeOfFile(fileName)
		log.Printf("Generation finished!\n-----\nLines:%v\nSize:%v\nLang:%v\nFile:%v \n-----\n", lines, size, lang, fileName)
	}(generator)

}
