package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sohlich/survey_kiosk/domain"
	// "log"
)

func IfPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()

	//Database
	//	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost:5432/kiosk?sslmode=disable")
	err := domain.OpenDatabase("postgres://postgres:postgres@localhost:5432/kiosk?sslmode=disable")
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
	router.POST("/survey/new", CreateSurvey)
	router.POST("/question/new", CreateQuestion)
	router.POST("/answertemplate/new", CreateAnswerTemplate)
	router.POST("/answer/new", CreateAnswer)
	router.POST("/person/new", CreatePerson)
	router.GET("/survey/:id", GetSurvey)
	router.GET("/answer/:id", GetAnswer)
}
