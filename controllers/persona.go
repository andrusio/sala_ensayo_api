package controllers

import (
	"fmt"
	"net/http"
	. "sala_ensayo/server/database"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
}

func GetPersonas(c *gin.Context) {
	results := GetConsulta("SELECT * FROM persona")
	personas := []Persona{}
	for results.Next() {
		var persona Persona
		var err error
		err = results.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Telefono)
		if err != nil {
			panic(err.Error())
		}
		personas = append(personas, persona)
	}

	c.IndentedJSON(http.StatusOK, personas)
}

func PostPersona(c *gin.Context) {
	var newPersona Persona
	if err := c.BindJSON(&newPersona); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	sqlstm := fmt.Sprintf("INSERT INTO persona (nombre, apellido, telefono)"+
		" VALUES ('%s','%s','%s')",
		newPersona.Nombre, newPersona.Apellido, newPersona.Telefono)
	GetConsulta(sqlstm)

	c.IndentedJSON(http.StatusCreated, newPersona)
}
