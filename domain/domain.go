package domain

import (
	// "time"
	// "fmt"
	// "log"

	"github.com/jinzhu/gorm"
	"gopkg.in/validator.v2"
)

var database gorm.DB

type Survey struct {
	Id           int    `gorm:"primary_key"`
	Name         string `validate:"nonzero"`
	ValidFrom    int64
	ValidTo      int64
	SurveyTypeId int `validate:"nonzero"`
	Questions    []Question
}

type Question struct {
	Id              int `gorm:"primary_key"`
	Value           string
	SurveyId        int
	Survey          Survey
	AnswerTemplates []AnswerTemplate
}

type AnswerTemplate struct {
	Id             int    `gorm:"primary_key"`
	Question_value string //TODO prejmenovat ??
	QuestionId     int
	Question       Question
}

type Answer struct {
	Id               int `gorm:"primary_key"`
	QuestionId       int
	AnswerTemplateId int
	IsFinal          bool
	PersonId         int
	AnswerTime       int64
	Question         Question
	AnswerTemplate   AnswerTemplate
	Person           Person
}

type Person struct {
	Id        int `gorm:"primary_key"`
	FirstName string
	LastName  string
}

//Function opens connection to database and function initDatabase
// will initialize database schema.
func OpenDatabase(connectionString string) error {
	var err error
	database, err = gorm.Open("postgres", connectionString)
	initDatabase(&database)
	return err
}

//Function to create schema and ensure indexes.
func initDatabase(database *gorm.DB) {
	database.CreateTable(&Survey{})
	database.CreateTable(&Question{})
	database.CreateTable(&Person{})
	database.CreateTable(&AnswerTemplate{})
	database.CreateTable(&Answer{})

	//Define FKs
	database.Model(&Question{}).AddForeignKey("survey_id", "surveies", "CASCADE", "RESTRICT")
	database.Model(&AnswerTemplate{}).AddForeignKey("question_id", "questions", "CASCADE", "RESTRICT")
	database.Model(&Answer{}).AddForeignKey("question_id", "questions", "CASCADE", "RESTRICT")
	database.Model(&Answer{}).AddForeignKey("answer_template_id", "answer_templates", "CASCADE", "RESTRICT")
	database.Model(&Answer{}).AddForeignKey("person_id", "persons", "CASCADE", "RESTRICT")
}

//Function to close database connection at the end of application.
//Usualy called as DEFER at main method.
func CloseDatabase() error {
	return database.Close()
}

//Saves survey to database. Befoore it saves the object, the object
//is validated. If validatiom fails, the ValidationError is thrown.
func Save(object interface{}) error {
	if err := validator.Validate(object); err != nil {
		//Validation failed throw wrapped typed error
		return &ValidationError{
			InternalError: err,
		}
	}
	//Save whole object
	database.Save(object)
	return nil
}

func Find(object interface{}, id int) error {
	//Switch provides preloading based on type
	switch object.(type) {
	case *Survey:
		database.Preload("Questions").Preload("Questions.AnswerTemplates").Find(object, id)
	default:
		database.Find(object, id)
	}
	return nil
}
