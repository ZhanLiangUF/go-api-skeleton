package main

import (
	"net"
	"os"

	"github.com/ZhanLiangUF/go-api-skeleton/api"
	"github.com/ZhanLiangUF/go-api-skeleton/api/middleware"
	router "github.com/ZhanLiangUF/go-api-skeleton/api/router"
	"github.com/ZhanLiangUF/go-api-skeleton/api/router/test"
	"github.com/ZhanLiangUF/go-api-skeleton/pkg/jsonmessage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initLogging()

	newConfig()
	// TODO: Unmarshal config
	config := &api.Config{}
	api := api.New(config)

	// accept listener
	l, _ := net.Listen("tcp", ":0080")
	api.Accept(":0080", l)

	// use middleware

	// if config.CorsHeader != "" {
	c := middleware.NewCORSMiddleware(config.CorsHeader)
	api.UseMiddleware(c)
	// }

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

// initLogging sets up logger configurations
func initLogging() {
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
		logrus.SetOutput(f)
		// set logging mode depending on different environment
		// logrus.SetLevel(logrus.TraceLevel)
	}
}

func newConfig() {
	// Inject config name and path to be use from jenkins?
	viper.SetConfigName("development")   // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../../configs") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Unable to read in configuration file: %v", err)
	}
}
