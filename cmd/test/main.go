package main

import (
	"context"
	"fmt"

	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/oleglacto/traning_scheduler/internal/configs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := configs.Database{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)

	conn, err := pgx.Connect(context.Background(), cfg.GetConnectionUrl())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int64
	var name string
	err = conn.QueryRow(context.Background(), "select id, name from test where id=$1", 1).Scan(&id, &name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, name)
}
