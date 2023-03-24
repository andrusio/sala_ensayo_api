package controllers

import (
	"database/sql"
	"log"
	"net/http"
	sqldb "sala_ensayo/server/database"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
}

func GetPersonas(c *gin.Context) {
	db := sqldb.ConnectDB()
	results, err := db.Query("SELECT id, nombre, apellido, telefono FROM persona")
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(204)
			return
		}
	}

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

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`INSERT INTO persona (nombre, apellido, telefono) VALUES (?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(newPersona.Nombre, newPersona.Apellido, newPersona.Telefono)
	id, err := res.LastInsertId()
	newPersona.ID = int(id)
	if err != nil {
		log.Fatalf("Error al agregar persona: %s", err)
	}

	c.IndentedJSON(http.StatusCreated, newPersona)
}
