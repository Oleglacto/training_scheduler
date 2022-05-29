package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/oleglacto/traning_scheduler/internal/configs"
	"github.com/oleglacto/traning_scheduler/internal/pkg/models"
	"log"
	"os"
	"strconv"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	records := readCsvFile(path + "/cmd/cities_parser/utf8_cities.csv")
	cities := rawToModel(records)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := configs.Database{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	conn, err := pgx.Connect(context.Background(), cfg.Dsn())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = seed(conn, cities)

	if err != nil {
		log.Fatal(err)
	}
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

func seed(conn *pgx.Conn, cities []models.City) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Prepare(
		context.Background(),
		"insert_cities",
		"INSERT INTO cities "+
			"(uuid, name, region, federal_district, latitude, longitude)"+
			" VALUES "+
			"($1,$2,$3,$4,$5,$6)",
	)

	if err != nil {
		return err
	}

	for _, v := range cities {
		_, err = tx.Exec(
			context.Background(),
			"insert_cities",
			v.ID,
			v.Name,
			v.Region,
			v.FederalDistrict,
			v.Location.Latitude,
			v.Location.Longitude,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}

func rawToModel(records [][]string) []models.City {
	cities := make([]models.City, 0, len(records))

	for _, row := range records {
		latitude, err := strconv.ParseFloat(row[3], 64)
		longitude, err := strconv.ParseFloat(row[4], 64)

		if err != nil {
			continue
		}

		city := models.City{
			ID:              uuid.NewString(),
			Name:            row[0],
			Region:          row[1],
			FederalDistrict: row[2],
			Location: models.Location{
				Latitude: latitude, Longitude: longitude,
			},
		}

		cities = append(cities, city)
	}

	return cities
}
