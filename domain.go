package main
import (
	"time"
)


type Survey struct{
	Id int `gorm:"primary_key"`
	Name string
	ValidFrom time.Time
	ValidTo time.Time
	SurveyTypeId int
}

type Question struct {
	Id int `gorm:"primary_key"`
	Value string
	SurveyId int
	Survey Survey
}

type AnswerTemplate struct {
	Id int `gorm:"primary_key"`
	question_value string
	QuestionId int
	Question Question
}

type Answer struct {
	Id int `gorm:"primary_key"`
	QuestionId int
	AnswerTemplateId int
	IsFinal bool
	PersonId int
	AnswerTime time.Time
	Question Question
	AnswerTemplate AnswerTemplate
	Person Person
}

type Person struct {
	Id int `gorm:"primary_key"`
	FirstName string
	LastName string
}
