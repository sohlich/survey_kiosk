package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sohlich/survey_kiosk/domain"
)

type Response struct {
	Reason string
	Error  error
}

func CreateAnswerTemplate(ctx *gin.Context) {
	answerTemplate := domain.AnswerTemplate{}
	create(&answerTemplate, ctx)
}

func CreateAnswer(ctx *gin.Context) {
	answer := domain.Answer{}
	// answer.Question = nil
	create(&answer, ctx)
}

func CreateSurvey(ctx *gin.Context) {
	survey := domain.Survey{}
	create(&survey, ctx)
}

func CreateQuestion(ctx *gin.Context) {
	question := domain.Question{}
	create(&question, ctx)
}

func CreatePerson(ctx *gin.Context) {
	person := domain.Person{}
	create(&person, ctx)
}

func create(entity domain.ValidableEntity, ctx *gin.Context) {
	err := ctx.BindJSON(entity)
	if IsError(err, ctx) {
		return
	}
	err = domain.Save(entity)
	if IsError(err, ctx) {
		return
	}
	ctx.JSON(200, entity)
}

func IsError(err error, ctx *gin.Context) bool {

	if err != nil {
		log.Println(err)
		ctx.JSON(405, Response{err.Error(), err})
		return true
	}
	return false
}

func GetSurvey(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	survey := domain.Survey{Id: id}
	_ = domain.Find(&survey, id)
	ctx.JSON(200, survey)
}

func GetAnswer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	answer := domain.Answer{Id: id}
	_ = domain.Find(&answer, id)
	ctx.JSON(200, answer)
}
