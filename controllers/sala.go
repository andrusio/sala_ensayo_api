package controllers

import (
	"database/sql"
	"net/http"
	sqldb "sala_ensayo/server/database"
	"sala_ensayo/server/helpers"

	"github.com/gin-gonic/gin"
)

type Sala struct {
	ID     int64   `json:"id"`
	Nombre string  `json:"nombre" binding:"required"`
	Precio float32 `json:"precio" binding:"required"`
	Color  int     `json:"color" binding:"required"`
}

func GetSalas(c *gin.Context) {

	db := sqldb.ConnectDB()
	res, err := db.Query("SELECT id, nombre, precio, color FROM sala")
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, gin.H{"error": "No hay salas cargadas"})
			return
		}
	}

	salas := []Sala{}
	for res.Next() {
		var sala Sala
		var err error
		err = res.Scan(&sala.ID, &sala.Nombre, &sala.Precio, &sala.Color)
		if err != nil {
			helpers.Error.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error al validar datos: " + err.Error()})
			return
		}
		salas = append(salas, sala)
	}
	c.IndentedJSON(http.StatusOK, salas)
}

func PostSala(c *gin.Context) {
	var newSala Sala
	if err := c.BindJSON(&newSala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos: " + err.Error()})
		return
	}

	db := sqldb.ConnectDB()
	res, err := db.Exec("INSERT INTO sala (nombre, precio, color) VALUES (?,?,?)", newSala.Nombre, newSala.Precio, newSala.Color)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Sala: " + err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Sala: " + err.Error()})
		return
	}
	newSala.ID = id
	c.IndentedJSON(http.StatusCreated, newSala)
}

func PutSala(c *gin.Context) {
	var newSala Sala
	if err := c.BindJSON(&newSala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos: " + err.Error()})
		return
	}

	db := sqldb.ConnectDB()
	res, err := db.Exec("UPDATE sala SET nombre = ?, precio = ?, color = ? WHERE id = ?", newSala.Nombre, newSala.Precio, newSala.Color, newSala.ID)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al modificar sala: " + err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al modificar Sala: " + err.Error()})
		return
	}

	newSala.ID = id
	c.IndentedJSON(http.StatusOK, newSala)
}

func DeleteSala(c *gin.Context) {
	id := c.Param("id")
	db := sqldb.ConnectDB()
	_, err := db.Exec(`DELETE FROM sala WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error al eliminar sala: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Sala eliminada con éxito"})
}
