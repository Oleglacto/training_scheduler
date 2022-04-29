package main

import (
	"context"
	"github.com/oleglacto/traning_scheduler/internal/app/training_scheduler/httpserver"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HelloWorld struct {
	Text string `json:"text"`
}

func main() {
	server := httpserver.NewServer("localhost", "4000")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(server.GetContext(), 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := server.Serve()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-server.GetContext().Done()
}
