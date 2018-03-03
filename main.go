package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexanderFadeev/go-course/handlers"
	"github.com/AlexanderFadeev/go-course/uploader"
	log "github.com/sirupsen/logrus"
)

const defaultLogFile = "log.log"
const defaultStaticDir = "/home/alexander/Programming/Go/src/github.com/AlexanderFadeev/go-course/static"

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

	server := startServer(address, defaultStaticDir)

	waitKillSignalChan(killSignalChan)
	server.Shutdown(context.Background())
}

func startServer(address, staticDir string) *http.Server {
	fileUploader := uploader.New(staticDir)
	router := handlers.NewRouter(staticDir, fileUploader)
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
	switch <-killSignalChan {
	case os.Interrupt:
		log.Info("got SIGINT")
	case syscall.SIGTERM:
		log.Info("got SIGTERM")
	}
}
