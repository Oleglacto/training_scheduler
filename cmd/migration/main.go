package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/oleglacto/traning_scheduler/internal/configs"
	"github.com/pressly/goose/v3"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/oleglacto/traning_scheduler/internal/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := configs.DataBaseConfig{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	flags.Parse(os.Args[1:])
	args := flags.Args()

	secret, command := args[0], args[1]

	if secret != os.Getenv("SECRET_FOR_MIGRATIONS") {
		fmt.Println("Can't migrate")
	}

	db, err := goose.OpenDBWithDriver("pgx", cfg.GetConnectionUrl())
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	dir := "./internal/migrations"

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
