package main

import (
	"cmp"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var (
	app         = echo.New()
	logger      zerolog.Logger
	port        = cmp.Or(getValFromEnv("SERVER_PORT"), getValfromVault("SERVER_PORT"), "9000")
	host        = cmp.Or(getValFromEnv("SERVER_HOST"), getValfromVault("SERVER_HOST"), "0.0.0.0")
	logFilename = cmp.Or(getValFromEnv("LOG_FILENAME"), getValfromVault("LOG_FILENAME"), "task_service.log")
)

func main() {
	address := fmt.Sprintf("%s:%s", host, port)

	appLogFile, _ := os.OpenFile(
		logFilename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer appLogFile.Close()

	setUp(appLogFile)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.Logger.Fatal(app.Start(address))
}
