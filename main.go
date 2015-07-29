package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sohlich/survey_kiosk/domain"
)

func IfPanic(err error) {
	if err != nil {
		panic(err)
	}
}



func main() {
	router := gin.Default()

	//Database
//	db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost:5432/kiosk?sslmode=disable")
	err := domain.OpenDatabase("postgres://liferay:liferay@localhost:5432/cloudkiosk?sslmode=disable")
	IfPanic(err)
	defer domain.CloseDatabase()
	defineRouting(router)
	router.Run(":8080")
}



func defineRouting(router *gin.Engine) {
	router.POST("/survey/new", CreateSurvey)
}
