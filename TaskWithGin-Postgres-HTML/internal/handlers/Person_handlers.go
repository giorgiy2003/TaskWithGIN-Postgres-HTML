package Handler

import (
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"

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
		c.HTML(400, "ErrorPage", gin.H{
			"Error": err,
		})
		return
	}
	c.HTML(200, "returnPage", nil)
}

func GetPersons(c *gin.Context) {
	persons, err := Logic.Read()
	if err != nil {
		c.HTML(400, "ErrorPage", gin.H{
			"Error": err,
		})
		return
	}
	if len(persons) == 0 {
		c.HTML(200, "InfoPage", gin.H{
			"Info": "Нет информации",
		})
		return
	}
	c.HTML(200, "Title", gin.H{"Title": "Список сотрудников"}) //Вывод заголовка
	for _, Person := range persons {
		c.HTML(200, "mainForm", gin.H{
			"Id":        Person.Id,
			"Email":     Person.Email,
			"Phone":     Person.Phone,
			"FirstName": Person.FirstName,
			"LastName":  Person.LastName,
		})
	}
}

func Form_handler_GetById(c *gin.Context) {
	id := c.Request.FormValue("id")
	persons, err := Logic.ReadOne(id)
	if err != nil {
		c.HTML(400, "ErrorPage", gin.H{
			"Error": err,
		})
		return
	}
	if len(persons) == 0 {
		c.HTML(200, "InfoPage", gin.H{
			"Info": "Нет информации",
		})
		return
	}
	c.HTML(200, "Title", gin.H{"Title": "Список сотрудников"}) //Вывод заголовка
	for _, Person := range persons {
		c.HTML(200, "mainForm", gin.H{
			"Id":        Person.Id,
			"Email":     Person.Email,
			"Phone":     Person.Phone,
			"FirstName": Person.FirstName,
			"LastName":  Person.LastName,
		})
	}
}

func Form_handler_DeleteById(c *gin.Context) {
	id := c.Request.FormValue("id")
	err := Logic.Delete(id)
	if err != nil {
		c.HTML(400, "ErrorPage", gin.H{
			"Error": err,
		})
		return
	}
	c.HTML(200, "returnPage", nil)
}

func Form_handler_UpdatePersonById(c *gin.Context) {
	var newPerson Model.Person
	id := c.Request.FormValue("id")
	newPerson.Email = c.Request.FormValue("email")
	newPerson.Phone = c.Request.FormValue("phone")
	newPerson.FirstName = c.Request.FormValue("firstName")
	newPerson.LastName = c.Request.FormValue("lastName")
	err := Logic.Update(newPerson, id)
	if err != nil {
		c.HTML(400, "ErrorPage", gin.H{
			"Error": err,
		})
		return
	}
	c.HTML(200, "returnPage", nil)
}

func ConnectDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := Repository.OpenTable(); err != nil {
			c.HTML(500, "InternalServerError", gin.H{
				"Error": err,
			})
			return
		}
	}
}

func Add(c *gin.Context) {
	c.HTML(200, "CreatePerson", nil)
}
func Remove(c *gin.Context) {
	c.HTML(200, "DeleteById", nil)
}
func Edit(c *gin.Context) {
	c.HTML(200, "EditPerson", nil)
}
