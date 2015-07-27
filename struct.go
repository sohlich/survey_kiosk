package main
import (
	"time"
)


type Survey struct{
	Id int `gorm:"primary_key"`
	Question string
	ValidFrom time.Time
	ValidTo time.Time
	SurveyTypeId int
	SurveyType SurveyType
}


type SurveyType struct {
	Id int `gorm:"primary_key"`
	Name string
}

type Response struct {
	Id int `gorm:"primary_key"`
	ResponseTime time.Time
	Answer string
	SurveyId int
	Survey Survey
}
