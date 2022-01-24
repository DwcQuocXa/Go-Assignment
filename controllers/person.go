package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Person struct {
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	PersonalCode uuid.UUID `json:"personalCode"`
}

var PersonList []Person

func GetPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, PersonList)
}

// Create a person
func CreatePerson(c *gin.Context) {
	var newPerson Person

	if err := c.BindJSON(&newPerson); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newPerson.PersonalCode = uuid.NewV4()

	PersonList = append(PersonList, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

// Get a person
func GetPersonByCode(c *gin.Context) {
	codeParam, _ := uuid.FromString(c.Param("code"))

	for _, item := range PersonList {
		if item.PersonalCode == codeParam {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

// Update a person
func UpdatePerson(c *gin.Context) {
	codeParam, _ := uuid.FromString(c.Param("code"))

	for index, item := range PersonList {
		if item.PersonalCode == codeParam {
			PersonList = append(PersonList[:index], PersonList[index+1:]...)

			var person Person
			if err := c.BindJSON(&person); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}

			person.PersonalCode = codeParam
			PersonList = append(PersonList, person)

			c.IndentedJSON(http.StatusOK, person)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

// Delete a person
func DeletePerson(c *gin.Context) {
	codeParam, _ := uuid.FromString(c.Param("code"))

	for index, item := range PersonList {
		if item.PersonalCode == codeParam {
			PersonList = append(PersonList[:index], PersonList[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "person deleted"})
			break
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}
