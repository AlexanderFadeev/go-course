package main

import (
	"net/http"
	"github.com/AlexanderFadeev/go-course/handlers"
	"os"
	"io"
	log "github.com/sirupsen/logrus"
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

	const address = ":8000"
	log.WithFields(log.Fields{
		"url": address,
	}).Info("Starting the server")

	router := handlers.Router()
	log.Fatal(http.ListenAndServe(address, router))
}
