package main

import (
	"encoding/csv"
	"flag"
	"github.com/jaswdr/faker"
	"log"
	"math/rand"
	"os"
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
	fileName := "Таблица_для_импорта_кандидатов_в_ХантФлоу_" + dateTime
	fake := faker.New()

	f, err := os.Create(fileName + ".csv")
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(f)

	if err := w.Write(firstLine); err != nil {
		log.Fatalln("Failed to write first line", err)
	}

	log.Println("Generation is started ...")
	for i := 0; i < count; i++ {
		maxTime := int64(rand.Int31n(int32(unix)))
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

		if err := w.Write(record); err != nil {
			log.Fatalln("Failed to write line #"+string(rune(i)), err)
		}
		//log.Println(record)
		//log.Println(time.Now().Format("02.01.2006"))

	}

	defer func(f *os.File, w *csv.Writer) {
		st, _ := f.Stat()
		size := strconv.Itoa(int(st.Size() / 1024))
		log.Printf("Generation finished, lines:%v size:%vKB, file:%v \n", count, size, f.Name())
		w.Flush()
		err := f.Close()
		if err != nil {
			log.Println("File Close with Error:", err)
		}
	}(f, w)

}
