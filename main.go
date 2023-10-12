package main

import (
	"context"
	"flag"
	"github.com/MrRytis/go-weather/internal"
	"github.com/MrRytis/go-weather/internal/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Go Weather API
// @description This is a sample weather API server written in Go.
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1

// @contactName Rytis Jančeris
// @contactEmail rytis.janceris@gmail.com

// @schemes http
// @schemes https

// @securityDefinitions.basic ApiKeyAuth
func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := internal.NewApp()
	r := router.NewRouter(app)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Server started at http://%s", srv.Addr)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other service
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
