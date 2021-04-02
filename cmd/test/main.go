package main

import (
	"log"
	"net"
	"os"

	"github.com/ZhanLiangUF/go-api-skeleton/api"
	"github.com/ZhanLiangUF/go-api-skeleton/api/middleware"
	router "github.com/ZhanLiangUF/go-api-skeleton/api/router"
	"github.com/ZhanLiangUF/go-api-skeleton/api/router/test"
	"github.com/ZhanLiangUF/go-api-skeleton/pkg/jsonmessage"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set up logging
	logrus.Info("Set up logger")
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: jsonmessage.RFC3339NanoFixed,
		FullTimestamp:   true,
	})
	var filename string = "logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logrus.Error(err)
	} else {
		log.SetOutput(f)
	}

	config := &api.Config{}
	api := api.New(config)

	// accept listener
	l, _ := net.Listen("tcp", ":3000")
	api.Accept(":3000", l)

	// use middleware
	if config.CorsHeader != "" {
		c := middleware.NewCORSMiddleware(config.CorsHeader)
		api.UseMiddleware(c)
	}
	// if tls?

	// init routers
	routers := []router.Router{
		test.NewRouter(),
	}
	// unpacks the slice with ellipses after the slice
	api.InitRouters(routers...)

	if err := api.ServeAPI(); err != nil {
		logrus.Errorf("ServeAPI error: %v", err)
	}
}
