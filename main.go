package main

import (
	"log"

	"code.google.com/p/gcfg"
	"github.com/gin-gonic/gin"
	"github.com/sohlich/survey_kiosk/domain"

	_ "github.com/lib/pq"
)

type Config struct {
	ConnectionString string
}

type configFile struct {
	Database Config
}

func IfPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	config := LoadConfiguration("application.conf")
	router := gin.Default()

	//Database
	err := domain.OpenDatabase(config.ConnectionString)
	IfPanic(err)
	defer domain.CloseDatabase()
	defineMiddleware(router)
	defineRouting(router)
	router.Run(":8080")
}

func defineMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}

func defineRouting(router *gin.Engine) {
	router.POST("/answer/new", CreateAnswer)
	router.POST("/question/new", CreateQuestion)
	router.POST("/survey/new", CreateSurvey)
	router.POST("/person/new", CreatePerson)
	router.GET("/survey/:id", GetSurvey)
}

func LoadConfiguration(cfgFile string) Config {
	var err error
	var cfg configFile
	if cfgFile != "" {
		err = gcfg.ReadFileInto(&cfg, cfgFile)
	} else {
		log.Panic("Cant read configuration file")
		// err = gcfg.ReadStringInto(&cfg, defaultConfig)
	}
	if err != nil {
		log.Panic(err)
	}
	return cfg.Database
}
