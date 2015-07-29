package domain

import (
	"time"

	"gopkg.in/validator.v2"
	"github.com/jinzhu/gorm"
)


var database gorm.DB

type Survey struct {
	Id           int `gorm:"primary_key"`
	Name         string
	ValidFrom    time.Time
	ValidTo      time.Time
	SurveyTypeId int
}

type Question struct {
	Id       int `gorm:"primary_key"`
	Value    string
	SurveyId int
	Survey   Survey
}

type AnswerTemplate struct {
	Id             int `gorm:"primary_key"`
	Question_value string
	QuestionId     int
	Question       Question
}

type Answer struct {
	Id               int `gorm:"primary_key"`
	QuestionId       int
	AnswerTemplateId int
	IsFinal          bool
	PersonId         int
	AnswerTime       time.Time
	Question         Question
	AnswerTemplate   AnswerTemplate
	Person           Person
}

type Person struct {
	Id        int `gorm:"primary_key"`
	FirstName string
	LastName  string
}


func OpenDatabase(connectionString string) error {
	var err error;
	database,err = gorm.Open("postgres",connectionString);
	initDatabase(&database)
	return err;
}

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

func CloseDatabase() error {
	return database.Close()
}


func Save(survey *Survey) error {
	if err := validator.Validate(survey);err != nil{
		return err
	}
	database.Save(survey)
	return nil
}

func Delete() {
//	database.Dele
}





