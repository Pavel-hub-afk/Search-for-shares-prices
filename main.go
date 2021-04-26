package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

func main() {
	var ticker string
	var date string

	fmt.Print("Enter the ticker: ")
	fmt.Scan(&ticker)
	fmt.Print("Enter the date (format 'year-month-date'): ")
	fmt.Scan(&date)

	var url string = "https://www.macrotrends.net/assets/php/stock_data_download.php?s=6086f7aae4fe1&t=" + ticker

	// URL запрос
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Создание структуры для работы с csv
	inf := csv.NewReader(resp.Body)

	for {
		record, err := inf.Read()
		if err == io.EOF { // Конец потока данных
			break
		}

		if record[0] == date {
			fmt.Printf("----------------------\nPrice on the specified date: %s$\n----------------------", record[4])
			return
		}
	}
	fmt.Print("----------------------\nData isn't correct\n----------------------")
}
