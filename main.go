package main

import (
	"net/http"
	"github.com/AlexanderFadeev/go-course/handlers"
	"os"
	"io"
	log "github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
	"context"
)

const defaultLogFile = "log.log"

func main() {
	logFile, err := os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(logWriter)
		defer logFile.Close()
	}

	log.SetFormatter(new(log.JSONFormatter))

	killSignalChan := getKillSignalChan()

	const address = ":8000"
	log.WithFields(log.Fields{
		"url": address,
	}).Info("Starting the server")

	server := startServer(address)

	waitKillSignalChan(killSignalChan)
	server.Shutdown(context.Background())
}

func startServer(address string) *http.Server {
	router := handlers.Router()
	server := http.Server{
		Addr:    address,
		Handler: router,
	}
	go func() {
		log.Fatal(http.ListenAndServe(address, router))
	}()
	return &server
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitKillSignalChan(killSignalChan <-chan os.Signal) {
	signal := <-killSignalChan
	switch signal {
	case os.Interrupt:
		log.Info("got SIGINT")
	case syscall.SIGTERM:
		log.Info("got SIGTERM")
	}
}
