package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// чтение аргументов
	var datafile *string
	// объявляем датафайл как строку с указателем
	datafile = flag.String("datafile", "./indications.json", `Path to datafile. Default: "./indications.json"`)
				// инициализация: назначаем флаг c типом строка датафайл на json , указываем путь к json
	flag.Parse() // пробегаем по файлу

	log.Printf("Datafile: %s\n", *datafile)

	// открытие файла
	f, err := os.Open(*datafile) // открытие файла
	if err != nil { // выявление ошибки
		log.Printf("ERROR os.Open, %s\n", err)
		return
	}

	// преобразование в срез
	data, err := parseFile(f)
	if err != nil {
		log.Printf("ERROR parseFile, %s\n", err)
		return
	}

	// сортировка данных
	err = sortByDate(data)
	if err != nil {
		log.Printf("ERROR sortByDate, %s\n", err)
		return
	}

	// форматированный вывод
	print(data) 
}

// Indication структура показания
type Indication struct { // создание структуры 
	Indicator string    `json:"indicator"` // теги на которые мы ориентируемся в json
	Value     int       `json:"value"`
	Date      time.Time `json:"date"`
}

// преобразование файла в срез показаний приборов
func parseFile(file *os.File) (data []Indication, err error) {
				// считывание файла, выходные данные типа срез и ошибка
	var (
		b []byte // число от 0 до 255 занимающее 1 байт
	)

	// чтение содержимого файла
	b, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// преобразование из json в срез структур
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return
}

// функция сортировки данных по дате. Заменяет ссылку входного среза на отсортированный срез
func sortByDate(data []Indication) (err error) {
	var (
		sortedData []Indication
	)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
		// timestamp := item.Date.Unix() // числовое представление даты (секунды с 1 января 1970 года). поле для сортировки
			if data[j].Date.Unix() < data[j+1].Date.Unix() {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}	
		}
	}

	data = sortedData
	

	return
}

// функция форматированного вывода среза показаний
func print(data []Indication) {
	var (
		res string
	)

	for _, i := range data {
		res = fmt.Sprintf("%s", res)
		res += fmt.Sprintf("Indicator: %v ", i.Indicator)
		res += fmt.Sprintf("value: %v ", i.Value) // поле "значение" показания прибора
		res += fmt.Sprintf("date: %v", i.Date)
		res += fmt.Sprintf("\n")
	}

	log.Printf(res)
}
