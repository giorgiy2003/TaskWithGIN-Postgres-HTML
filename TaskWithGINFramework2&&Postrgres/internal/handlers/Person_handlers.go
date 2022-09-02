package Handler

import (
	Model "TaskWithGINFramework/internal/model"
	Logic "TaskWithGINFramework/internal/logic"
	Repository "TaskWithGINFramework/internal/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Form_handler_PostPerson(c *gin.Context) {
	var newPerson Model.Person
	newPerson.Email = c.Request.FormValue("email")
	newPerson.Phone = c.Request.FormValue("phone")
	newPerson.FirstName = c.Request.FormValue("firstName")
	newPerson.LastName = c.Request.FormValue("lastName")

	Logic.Create(newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func Form_handler_GetPersons(c *gin.Context) {
	persons, _ := Logic.Read()

	c.IndentedJSON(http.StatusOK, persons)
}

func Form_handler_GetById(c *gin.Context) {

	id := c.Request.FormValue("id")
	persons, _ := Logic.ReadOne(id)

	c.IndentedJSON(http.StatusOK, persons)
}

func Form_handler_DeleteById(c *gin.Context) {

	id := c.Request.FormValue("id")
	Logic.Delete(id)

	c.IndentedJSON(http.StatusOK, Model.Person{})
}

func Form_handler_UpdatePersonById(c *gin.Context) {
	var newPerson Model.Person
	newPerson.Id = c.Request.FormValue("id")
	newPerson.Email = c.Request.FormValue("email")
	newPerson.Phone = c.Request.FormValue("phone")
	newPerson.FirstName = c.Request.FormValue("firstName")
	newPerson.LastName = c.Request.FormValue("lastName")
	Logic.Update(newPerson,newPerson.Id)

	c.IndentedJSON(http.StatusOK, newPerson)
}

func MainForm(c *gin.Context) {
	err := Repository.OpenTable()
	if err != nil {
		panic(err)
	}
	c.HTML(200, "mainForm.html", nil)
}

func Add(c *gin.Context) {
	c.HTML(200, "CreatePerson.html", nil)
}
func Remove(c *gin.Context) {
	c.HTML(200, "DeleteById.html", nil)
}
func Edit(c *gin.Context) {
	c.HTML(200, "EditPerson.html", nil)
}
func GetById(c *gin.Context) {
	c.HTML(200, "GetByID.html", nil)
}
