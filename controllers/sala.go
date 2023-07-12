package controllers

import (
	"database/sql"
	"log"
	"net/http"
	sqldb "sala_ensayo/server/database"

	"github.com/gin-gonic/gin"
)

type Sala struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre" binding:"required"`
	Precio float32 `json:"precio" binding:"required"`
	Color  int     `json:"color" binding:"required"`
}

func GetSalas(c *gin.Context) {
	db := sqldb.ConnectDB()
	results, err := db.Query("SELECT id, nombre, precio, color FROM sala")
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(204)
			return
		}
	}

	salas := []Sala{}
	for results.Next() {
		var sala Sala
		var err error
		err = results.Scan(&sala.ID, &sala.Nombre, &sala.Precio, &sala.Color)
		if err != nil {
			c.String(http.StatusBadRequest, "Error al validar datos: %s", err.Error())
			return
		}
		salas = append(salas, sala)
	}
	c.IndentedJSON(http.StatusOK, salas)
}

func PostSala(c *gin.Context) {
	var newSala Sala
	if err := c.BindJSON(&newSala); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`INSERT INTO sala (nombre, precio, color) VALUES (?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(newSala.Nombre, newSala.Precio, newSala.Color)
	if err != nil {
		log.Fatalf("Error al agregar sala: %s", err)
	}
	id, err := res.LastInsertId()
	newSala.ID = int(id)

	c.IndentedJSON(http.StatusCreated, newSala)
}

func PutSala(c *gin.Context) {
	var newSala Sala
	if err := c.BindJSON(&newSala); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`UPDATE sala SET nombre = ?, precio = ?, color = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(newSala.Nombre, newSala.Precio, newSala.Color, newSala.ID)
	if err != nil {
		c.String(http.StatusBadRequest, "Error al actualizar sala: %s", err.Error())
		return
	}
	id, err := res.LastInsertId()
	newSala.ID = int(id)

	c.IndentedJSON(http.StatusOK, newSala)
}

func DeleteSala(c *gin.Context) {
	id := c.Param("id")
	db := sqldb.ConnectDB()
	_, err := db.Exec(`DELETE FROM sala WHERE id = ?`, id)
	if err != nil {
		c.String(http.StatusBadRequest, "Error al eliminar sala: %s", err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "Sala eliminada con éxito")
}
