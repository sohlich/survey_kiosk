package domain

import (
	"testing"
	"time"
)

//Tests if validation works.
func SurveyValidationFailTest(t *testing.T) {
	survey := &Survey{
		ValidFrom: time.Now(),
		ValidTo:   time.Now(),
	}

	err := SaveSurvey(survey)

	_, ok := err.(*ValidationError)

	if !ok {
		t.Error("Negative validation of object failed")
	}
}
