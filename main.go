package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func Hello() {
	println("Hello World")
}

func main() {
	// Call the function
	Hello()

	// Call the function
	Hello()

	a := 10
	if a > 5 {
		println("a is greater than 5")
	} else {
		println("a is less than or equal to 5")
	}

	// ローカルにあるcsvファイルを読み込む
	ReadCSV()
}

func ReadCSV() {
	// ファイルを開く
	file, err := os.Open(`./data.csv`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// CSVファイルを読み込む
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
