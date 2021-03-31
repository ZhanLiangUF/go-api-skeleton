package main

import (
	"log"
	"os"

	router "github.com/ZhanLiangUF/go-api-skeleton/api/router"
	"github.com/ZhanLiangUF/go-api-skeleton/api/router/test"
	"github.com/ZhanLiangUF/go-api-skeleton/pkg/jsonmessage"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Set up logger")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: jsonmessage.RFC3339NanoFixed,
		FullTimestamp:   true,
	})
	var filename string = "logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		// Cannot open log file. Logging to stderr
		logrus.Error(err)
	} else {
		log.SetOutput(f)
	}

	routers := []router.Router{
		test.NewRouter(),
	}

}
