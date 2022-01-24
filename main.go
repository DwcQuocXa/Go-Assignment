package main

import (
	"example.com/zenniz-go-asignemnt/controllers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func main() {
	//Init router
	r := gin.Default()

	//First data
	controllers.PersonList = append(controllers.PersonList, controllers.Person{
		FirstName:    "John",
		LastName:     "Doe",
		PersonalCode: uuid.NewV4()})

	//Routes
	r.GET("/api/persons", controllers.GetPersons)
	r.POST("/api/persons", controllers.CreatePerson)
	r.GET("/api/persons/:code", controllers.GetPersonByCode)
	r.PUT("/api/persons/:code", controllers.UpdatePerson)
	r.DELETE("/api/persons/:code", controllers.DeletePerson)

	//Server and port
	r.Run(":8000")
}
