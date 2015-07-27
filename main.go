package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func IfPanic(err error){
	if(err != nil){
		panic(err)
	}
}

var database gorm.DB;


func main(){
	router := gin.Default()

	//Database
	db, err := gorm.Open("postgres", "postgres://liferay:liferay@localhost:5432/cloudkiosk?sslmode=disable")
	IfPanic(err)
	database = db
	defer database.Close()
	initDatabase(&database)
	defineRouting(router)
	router.Run(":8080")
}

func initDatabase(database *gorm.DB){
	database.CreateTable(&Survey{})
	database.CreateTable(&SurveyType{})
	database.CreateTable(&Response{})

	//Define FKs
	database.Model(&Survey{}).AddForeignKey("survey_type_id", "survey_types", "CASCADE", "RESTRICT")
}

func defineRouting(router *gin.Engine){

}



func SaveSurvey(survey *Survey){
	database.Save(survey)
}




//func SaveSurvey(context *gin.Context){
//	context.BindJSON()
//}