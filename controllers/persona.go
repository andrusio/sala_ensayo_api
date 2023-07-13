package controllers

import (
	"database/sql"
	"net/http"
	sqldb "sala_ensayo/server/database"
	"sala_ensayo/server/helpers"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
}

func GetPersonas(c *gin.Context) {
	db := sqldb.ConnectDB()
	results, err := db.Query("SELECT id, nombre, apellido, telefono FROM persona")
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, gin.H{"error": "No hay personas cargadas"})
			return
		}
	}

	personas := []Persona{}
	for results.Next() {
		var persona Persona
		var err error
		err = results.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Telefono)
		if err != nil {
			helpers.Error.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error al validar datos: " + err.Error()})
			return
		}
		personas = append(personas, persona)
	}

	c.IndentedJSON(http.StatusOK, personas)
}

func PostPersona(c *gin.Context) {
	var newPersona Persona
	if err := c.BindJSON(&newPersona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos: " + err.Error()})
		return
	}

	db := sqldb.ConnectDB()
	res, err := db.Exec(`INSERT INTO persona (nombre, apellido, telefono) VALUES (?,?,?)`, newPersona.Nombre, newPersona.Apellido, newPersona.Telefono)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Persona: " + err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Persona: " + err.Error()})
		return
	}
	newPersona.ID = id
	c.IndentedJSON(http.StatusCreated, newPersona)
}

func PutPersona(c *gin.Context) {
	var newPersona Persona
	if err := c.BindJSON(&newPersona); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	db := sqldb.ConnectDB()
	res, err := db.Exec(`UPDATE persona SET nombre = ?, apellido = ?, telefono = ? WHERE id = ?`, newPersona.Nombre, newPersona.Apellido, newPersona.Telefono, newPersona.ID)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al modificar Persona: " + err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al modificar Persona: " + err.Error()})
		return
	}

	newPersona.ID = id
	c.IndentedJSON(http.StatusOK, newPersona)
}

func DeletePersona(c *gin.Context) {
	id := c.Param("id")
	db := sqldb.ConnectDB()
	_, err := db.Exec(`DELETE FROM persona WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error al eliminar Persona: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, "Persona eliminada con éxito")
}
