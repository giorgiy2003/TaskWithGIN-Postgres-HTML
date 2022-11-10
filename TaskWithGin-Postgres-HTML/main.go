package main

import (
	Handler "myapp/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Use(Handler.ConnectDB())
	router.GET("/", Handler.GetPersons)
	router.GET("/Add", Handler.Add)
	router.GET("/Remove", Handler.Remove)
	router.GET("/Edit", Handler.Edit)
	router.GET("/Form_handler_GetById", Handler.Form_handler_GetById)
	router.POST("/Form_handler_PostPerson", Handler.Form_handler_PostPerson)
	router.GET("/Form_handler_DeleteById", Handler.Form_handler_DeleteById)
	router.GET("/Form_handler_UpdatePersonById", Handler.Form_handler_UpdatePersonById)
	router.Run("localhost:8080")
}
