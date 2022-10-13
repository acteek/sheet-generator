package main

import (
	"flag"
	"github.com/jaswdr/faker"
	"github.com/tealeg/xlsx"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var firstLine = []string{
	"Фамилия",
	"Имя",
	"Отчество",
	"Телефон",
	"Электронная почта",
	"Telegram",
	"Скайп",
	"Должность",
	"Компания",
	"Зарплатные ожидания",
	"Дата рождения (ДД.ММ.ГГГГ)",
	"Комментарии",
}

func main() {

	var count int
	flag.IntVar(&count, "c", 1000, "Count of lines, default 1000")
	flag.Parse()

	unix := time.Now().Unix()
	dateTime := time.Now().Format("02.01.2006T15.04.05")
	fileName := "Таблица_для_импорта_кандидатов_в_ХантФлоу_" + dateTime + ".xlsx"
	fake := faker.New()

	//if err != nil {
	//	log.Fatalln("Failed to open file", err)
	//}

	//w := csv.NewWriter(f)

	//if err := w.Write(firstLine); err != nil {
	//	log.Fatalln("Failed to write first line", err)
	//}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	rowMain := sheet.AddRow()
	rowMain.WriteSlice(&firstLine, len(firstLine))
	if err != nil {
		log.Fatalln("Create Sheet error:", err)
	}

	log.Println("Generation is started ...")
	for i := 0; i < count; i++ {
		maxTime := int64(rand.Int31n(int32(unix)))
		row := sheet.AddRow()

		var record = []string{
			fake.Person().LastName(),
			fake.Person().FirstName(),
			fake.Pet().Name(),
			fake.Phone().E164Number(),
			fake.Internet().Email(),
			fake.Internet().User(),
			fake.Internet().User(),
			fake.Company().JobTitle(),
			fake.Company().Name(),
			strconv.Itoa(int(fake.UInt32())),
			time.Unix(maxTime, 0).Format("02.01.2006"),
			fake.Lorem().Sentence(10),
		}

		row.WriteSlice(&record, len(record))
		//if err := w.Write(record); err != nil {
		//	log.Fatalln("Failed to write line #"+string(rune(i)), err)
		//}
		//log.Println(record)
		//log.Println(time.Now().Format("02.01.2006"))

	}

	defer func(f *xlsx.File, name string) {
		//st, _ := f.Stat()
		//size := strconv.Itoa(int(st.Size() / 1024))

		if err := f.Save(fileName); err != nil {
			log.Fatalln("Save file error:", err)
		}

		log.Printf("Generation finished, lines:%v, file:%v \n", count, fileName)
		//file.Save(fileName + ".xlsx")
		//err := f.Close()
		//if err != nil {
		//	log.Println("File Close with Error:", err)
		//}
	}(file, fileName)

}
