package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sohlich/survey_kiosk/domain"
)

type Response struct {
	Reason string
}

func CreateAnswerTemplate(ctx *gin.Context) {
	answerTemplate := domain.AnswerTemplate{}
	create(&answerTemplate, ctx)
}

func CreateAnswer(ctx *gin.Context) {
	answer := domain.Answer{}
	create(answer, ctx)
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

func create(entity interface{}, ctx *gin.Context) {
	if ctx.BindJSON(entity) == nil {
		domain.Save(entity)
		ctx.JSON(200, entity)
	} else {
		ctx.JSON(405, Response{"Malformed object"})
	}
}

func GetSurvey(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	survey := domain.Survey{Id: id}
	_ = domain.Find(&survey, id)
	ctx.JSON(200, survey)
}
