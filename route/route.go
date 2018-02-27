package route

import "github.com/gin-gonic/gin"

func Routing(router *gin.Engine) {
	router.GET("/", defaultRoute)

	router.GET("/peoples", loadAllPersons)
	router.GET("/peoples/:id", loadSpecificPerson)
	router.POST("/people", createNewPerson)
	router.PUT("/people/:id", updatePerson)
	router.DELETE("/people/:id", deletePerson)
}
