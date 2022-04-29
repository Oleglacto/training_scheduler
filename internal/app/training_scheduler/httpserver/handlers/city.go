package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetAllCities(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		log.Panicln("Request canceled")
	default:
		time.Sleep(5 * time.Second)
		w.Write([]byte(fmt.Sprintf("All right")))
	}

	return nil
}

func AddCity(w http.ResponseWriter, r *http.Request) error {

	return nil
}
