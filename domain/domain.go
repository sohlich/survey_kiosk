package domain

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/validator.v2"
)

// wrap gorm database to custom type for testing
type Database struct {
	gorm.DB
}

var database Database

type ValidableEntity interface {
	HasValidId() bool
}

type Survey struct {
	Id           int    `gorm:"primary_key"`
	Name         string `validate:"nonzero"`
	ValidFrom    int64
	ValidTo      int64
	SurveyTypeId int `validate:"nonzero"`
	Questions    []Question
}

func (survey *Survey) HasValidId() bool {
	return survey.Id != 0
}

type Question struct {
	Id              int `gorm:"primary_key"`
	Value           string
	SurveyId        int
	Survey          *Survey
	AnswerTemplates []AnswerTemplate
}

func (question *Question) HasValidId() bool {
	return question.Id != 0
}

type AnswerTemplate struct {
	Id          int `gorm:"primary_key"`
	AnswerValue string
	QuestionId  int
	Question    *Question
}

func (answerTemplate *AnswerTemplate) HasValidId() bool {
	return answerTemplate.Id != 0
}

type Answer struct {
	Id               int `gorm:"primary_key"`
	QuestionId       int
	AnswerTemplateId int
	IsFinal          bool
	PersonId         int `validate:"nonzero"`
	AnswerTime       int64
	// Question         *Question
	// AnswerTemplate   *AnswerTemplate
	// Person           *Person
}

func (answer *Answer) HasValidId() bool {
	return answer.Id != 0
}

type Person struct {
	Id        int    `gorm:"primary_key"`
	FirstName string `validate:"nonzero"`
	LastName  string `validate:"nonzero"`
}

func (person *Person) HasValidId() bool {
	return person.Id != 0
}

//Function opens connection to database and function initDatabase
// will initialize database schema.
func OpenDatabase(connectionString string) error {
	db, err := gorm.Open("postgres", connectionString)
	database = Database{db}
	initDatabase(&database)
	return err
}

//Function to create schema and ensure indexes.
func initDatabase(database *Database) {
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
func Save(object ValidableEntity) error {
	if err := validator.Validate(object); err != nil {
		//Validation failed throw wrapped typed error
		return &ValidationError{
			InternalError: err,
		}
	}

	//Save whole object
	database.Create(object)
	if !object.HasValidId() {
		return &OperationFailError{}
	}

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
