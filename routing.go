package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sohlich/survey_kiosk/domain"
)

func CreateSurvey(ctx *gin.Context) {
	log.Println("receiving survey")
	survey := domain.Survey{}
	if ctx.BindJSON(&survey) == nil {
		domain.Save(&survey)
		ctx.String(200, "OK")
	} else {
		ctx.String(405, "Bad request")
	}

}
