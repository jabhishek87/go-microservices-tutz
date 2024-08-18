package main

import (
	"cmp"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

var (
	app         = echo.New()
	logger      zerolog.Logger
	port        = cmp.Or(getValFromEnv("SERVER_PORT"), getValfromVault("SERVER_PORT"), "9000")
	host        = cmp.Or(getValFromEnv("SERVER_HOST"), getValfromVault("SERVER_HOST"), "0.0.0.0")
	logFilename = cmp.Or(getValFromEnv("LOG_FILENAME"), getValfromVault("LOG_FILENAME"), "task_service.log")
	dbFilename  = cmp.Or(getValFromEnv("DB_FILENAME"), getValfromVault("DB_FILENAME"), "test.db")
	db          *gorm.DB
)

// @title Swagger Task API
// @version 0.1
// @description Everything about taks API.
// @schemes http https
// @host localhost:9000
// @BasePath /

// @tag.name Root
// @tag.description Root Description
// @tag.docs.url https://example.com
// @tag.docs.description Root documentation

// @tag.name Tasks
// @tag.description Task Description
// @tag.docs.url https://example.com
// @tag.docs.description Task Documentation

//
//go:generate swag init
func main() {
	address := fmt.Sprintf("%s:%s", host, port)

	appLogFile, _ := os.OpenFile(
		logFilename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer appLogFile.Close()

	setUp(appLogFile)
	db = setUpDB()

	app.Add("GET", "/", getRoot)

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	v1Group := app.Group("/api/v1")
	v1Group.Add("GET", "/tasks", getTasks)
	v1Group.Add("POST", "/tasks", postTask)
	v1Group.Add("GET", "/tasks/:uuid", getTask)
	v1Group.Add("PUT", "/tasks/:uuid", putTask)
	v1Group.Add("DELETE", "/tasks/:uuid", delTask)

	// https://echo.labstack.com/docs/cookbook/crud

	app.Logger.Fatal(app.Start(address))
}
