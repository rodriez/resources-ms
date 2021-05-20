package main

import (
	"net/http"
	"os"
	"resources-ms/app"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file", err)
	}

	logrus.Info("Environment variables loaded")
}

func main() {
	lvl, _ := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	logrus.SetLevel(lvl)
	logrus.Info("Listening port :8080")

	if err := http.ListenAndServe(":8080", app.Router); err != nil {
		logrus.Fatal(err)
	}
}
