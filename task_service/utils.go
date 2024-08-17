package main

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// fallback values in Golang
// variadic functions

func getValFromEnv(key string) string {
	return os.Getenv(key)
}

func getValfromVault(key string) string {
	// not implemented
	// implement vault here
	//fmt.Println(key)
	return ""
}

func setUp(appLogFile *os.File) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	multi := zerolog.MultiLevelWriter(os.Stdout, appLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	log.Info().Msg("Log Setup")
	app.Logger.SetOutput(multi)

	app.Use(middleware.Logger())
	// app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogURI:    true,
	// 	LogStatus: true,
	// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	// 		log.logger.Info().
	// 			Str("URI", v.URI).
	// 			Int("status", v.Status).
	// 			Msg("request")

	// 		return nil
	// 	},
	// }))
}

func setUpDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Task{})
	// Create
	// db.Create(&Task{
	// 	Title:  "New Task",
	// 	UUID:   uuid.New(),
	// 	Status: "todo",
	// 	UserID: 0,
	// })

	return db
}
