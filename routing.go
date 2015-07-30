package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sohlich/survey_kiosk/domain"
)

type Response struct {
	Reason string
}

func CreateSurvey(ctx *gin.Context) {
	survey := domain.Survey{}
	create(&survey, ctx)
}

func CreateQuestion(ctx *gin.Context) {
	question := domain.Question{}
	create(&question, ctx)
}

func CreateAnswerTemplate(ctx *gin.Context) {
	answerTemplate := domain.AnswerTemplate{}
	create(&answerTemplate, ctx)
}

func create(entity interface{}, ctx *gin.Context) {
	if ctx.BindJSON(entity) == nil {
		domain.Save(entity)
		ctx.JSON(200, entity)
	} else {
		ctx.JSON(405, Response{"Malformed object"})
	}
}
