package main

import(
	"encoding/csv"
	"fmt"
	"os"
)

func
main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(record)
	}
}
