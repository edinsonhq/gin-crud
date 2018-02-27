package route

import (
	"strconv"

	"../controller"
	"github.com/gin-gonic/gin"
)

func defaultRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func loadAllPersons(c *gin.Context) {
	persons := controller.NewPerson()
	c.JSON(200, gin.H{
		"persons": persons.LoadAllPersons(),
	})
}

func loadSpecificPerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	person := controller.NewPerson()
	c.JSON(200, gin.H{
		"person": person.LoadSpecificPerson(id),
	})
}

func createNewPerson(c *gin.Context) {
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	person := controller.NewPerson()
	person.CreateNewPerson(firstname, lastname)
	c.JSON(200, "OK")
}

func updatePerson(c *gin.Context) {
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	person := controller.NewPerson()
	c.JSON(200, gin.H{
		"person": person.UpdatePerson(id, firstname, lastname),
	})
}

func deletePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	person := controller.NewPerson()
	person.DeletePerson(id)
	c.JSON(200, "OK")
}
