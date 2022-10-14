package generate

import "time"

const dateTimeFormat = "02.01.2006T15.04.05"
const fileNameTemp = "Таблица_для_импорта_кандидатов_в_ХантФлоу_"

var FirstLine = []string{
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

func FileNameGen(ext string) string {
	dateTime := time.Now().Format(dateTimeFormat)
	name := fileNameTemp + dateTime + "." + ext
	return name
}
