package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sohlich/survey_kiosk/domain"
)

func IfPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var database gorm.DB

func main() {
	router := gin.Default()

	//Database
	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost:5432/kiosk?sslmode=disable")
	IfPanic(err)
	database = db
	defer database.Close()
	initDatabase(&database)
	defineRouting(router)
	router.Run(":8080")
}

func initDatabase(database *gorm.DB) {
	database.CreateTable(&domain.Survey{})
	database.CreateTable(&domain.Question{})
	database.CreateTable(&domain.Person{})
	database.CreateTable(&domain.AnswerTemplate{})
	database.CreateTable(&domain.Answer{})

	//Define FKs
	database.Model(&domain.Question{}).AddForeignKey("survey_id", "surveies", "CASCADE", "RESTRICT")
	database.Model(&domain.AnswerTemplate{}).AddForeignKey("question_id", "questions", "CASCADE", "RESTRICT")
	database.Model(&domain.Answer{}).AddForeignKey("question_id", "questions", "CASCADE", "RESTRICT")
	database.Model(&domain.Answer{}).AddForeignKey("answer_template_id", "answer_templates", "CASCADE", "RESTRICT")
	database.Model(&domain.Answer{}).AddForeignKey("person_id", "persons", "CASCADE", "RESTRICT")
}

func defineRouting(router *gin.Engine) {
	router.POST("/survey/new", CreateSurvey)
}
