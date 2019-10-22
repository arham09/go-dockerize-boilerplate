package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	ghandlers "github.com/gorilla/handlers"

	"app/pkg/middlewares"
	"app/pkg/routers"
)

func main() {
	router := routers.New()

	loggedRouter := middlewares.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, router))

	server := &http.Server{
		Addr:         ":2019",
		Handler:      loggedRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Connected to port 2019 !!!")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
