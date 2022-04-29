package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CsvCity struct {
	name            string
	region          string
	federalDistrict string
	latitude        float64
	longitude       float64
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	records := readCsvFile(path + "/utf8_cities.csv")
	cities := rawToStruct(records)
	fmt.Println(cities)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = '\t'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func rawToStruct(records [][]string) []CsvCity {
	cities := make([]CsvCity, 0, len(records))

	for _, row := range records {
		latitude, err := strconv.ParseFloat(row[3], 64)
		longitude, err := strconv.ParseFloat(row[4], 64)

		if err != nil {
			continue
		}

		city := CsvCity{row[0], row[1], row[2], latitude, longitude}

		cities = append(cities, city)
	}

	return cities
}
