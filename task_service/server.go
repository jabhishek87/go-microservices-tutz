package main

import (
	"cmp"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	app         = echo.New()
	logger      zerolog.Logger
	port        = cmp.Or(getValFromEnv("SERVER_PORT"), getValfromVault("SERVER_PORT"), "9000")
	host        = cmp.Or(getValFromEnv("SERVER_HOST"), getValfromVault("SERVER_HOST"), "0.0.0.0")
	logFilename = cmp.Or(getValFromEnv("LOG_FILENAME"), getValfromVault("LOG_FILENAME"), "task_service.log")
)

// @title Swagger Task API
// @version 1.0
// @description This Task Service Swagger.
// @host localhost:9000
// @BasePath /api/v1
func main() {
	address := fmt.Sprintf("%s:%s", host, port)

	appLogFile, _ := os.OpenFile(
		logFilename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer appLogFile.Close()

	setUp(appLogFile)

	app.Add("GET", "/", getRoot)

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	v1Group := app.Group("/api/v1")
	v1Group.Add("GET", "/tasks", getTasks)
	v1Group.Add("GET", "/tasks/:id", getTask)
	v1Group.Add("POST", "/tasks", postTask)
	v1Group.Add("PUT", "/tasks/:id", putTask)
	v1Group.Add("DELETE", "/tasks/:id", delTask)

	// https://echo.labstack.com/docs/cookbook/crud

	app.Logger.Fatal(app.Start(address))
}
