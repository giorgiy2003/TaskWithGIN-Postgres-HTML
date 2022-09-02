package main

import (
	Handler "TaskWithGINFramework/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("html/*.html")
	router.GET("/", Handler.MainForm)
	router.GET("/GetById", Handler.GetById)
	router.GET("/Add", Handler.Add)
	router.GET("/Remove", Handler.Remove)
	router.GET("/Edit", Handler.Edit)

	router.GET("/form_handler_GetPersons", Handler.Form_handler_GetPersons)
	router.GET("/form_handler_GetById", Handler.Form_handler_GetById)
	router.POST("/form_handler_PostPerson", Handler.Form_handler_PostPerson)
	router.GET("/form_handler_DeleteById", Handler.Form_handler_DeleteById)
	router.GET("/form_handler_UpdatePersonById", Handler.Form_handler_UpdatePersonById)
	_ = router.Run("localhost:8080")
}
