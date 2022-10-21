package Handler

import (
	"fmt"
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Form_handler_PostPerson(c *gin.Context) {
	var newPerson Model.Person
	newPerson.Email = c.Request.FormValue("email")
	newPerson.Phone = c.Request.FormValue("phone")
	newPerson.FirstName = c.Request.FormValue("firstName")
	newPerson.LastName = c.Request.FormValue("lastName")
	err := Logic.Create(newPerson)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func Form_handler_GetPersons(c *gin.Context) {
	persons, err := Logic.Read()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, persons)
}

func Form_handler_GetById(c *gin.Context) {
	id := c.Request.FormValue("id")
	persons, err := Logic.ReadOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, persons)
}

func Form_handler_DeleteById(c *gin.Context) {
	id := c.Request.FormValue("id")
	err := Logic.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, "Запись удалена")
}

func Form_handler_UpdatePersonById(c *gin.Context) {
	var newPerson Model.Person
	newPerson.Id = c.Request.FormValue("id")
	newPerson.Email = c.Request.FormValue("email")
	newPerson.Phone = c.Request.FormValue("phone")
	newPerson.FirstName = c.Request.FormValue("firstName")
	newPerson.LastName = c.Request.FormValue("lastName")
	err := Logic.Update(newPerson, newPerson.Id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	c.IndentedJSON(http.StatusOK, "Запись обновлена")
}

func MainForm(c *gin.Context) {
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
