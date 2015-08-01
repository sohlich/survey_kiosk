package domain

import (
	"github.com/jinzhu/gorm"
	"log"
	"testing"
)

var testSurvey = Survey{
	Id:           1,
	Name:         "MySurvey",
	ValidFrom:    12345,
	ValidTo:      52345,
	SurveyTypeId: 1,
}

//Mocking function
func (instance *Database) Save(obj interface{}) *gorm.DB {
	if survey, ok := obj.(*Survey); ok {
		survey.Id = 1
	}
	return nil
}
func (instance *Database) Find(obj interface{}, conditions ...interface{}) *gorm.DB {
	if survey, ok := obj.(*Survey); ok {
		if survey.Id == testSurvey.Id {
			survey.Id = testSurvey.Id
			survey.Name = testSurvey.Name
			survey.SurveyTypeId = testSurvey.SurveyTypeId
		}
	}
	return nil
}

//Tests if validation works.
func TestSurveyValidationFail(t *testing.T) {
	survey := &Survey{}
	err := Save(survey)
	_, ok := err.(*ValidationError)
	if !ok {
		t.Error("Validation of object failed")
	}
}

func TestSurveyValidationOK(t *testing.T) {
	survey := &Survey{
		Name:         "My survey",
		SurveyTypeId: 1,
	}
	err := Save(survey)
	if err != nil {
		log.Println(err)
		t.Error("Save operation failed")
	}
}
