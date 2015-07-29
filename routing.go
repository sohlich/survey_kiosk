package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sohlich/survey_kiosk/domain"
)

func CreateSurvey(ctx *gin.Context) {
	survey := domain.Survey{}
	if ctx.BindJSON(&survey) == nill {
		log.Println(survey)
		ctx.String(200, "OK")
	} else {
		ctx.String(405, "Bad request")
	}

}
