package controllers

import (
	"fmt"
	"net/http"
	. "sala_ensayo/server/database"

	"github.com/gin-gonic/gin"
)

type Sala struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre" binding:"required"`
	Precio float32 `json:"precio" binding:"required"`
	Color  int     `json:"color" binding:"required"`
}

type SalaGrupo struct {
	ID        int    `json:"id"`
	Sala      string `json:"sala" binding:"required"`
	SalaColor int    `json:"sala_color" binding:"required"`
	Grupo     string `json:"grupo" binding:"required"`
	HoraDesde string `json:"hora_desde" binding:"required"`
	HoraHasta string `json:"hora_hasta" binding:"required"`
}

func GetSalas(c *gin.Context) {
	results := GetConsulta("SELECT * FROM sala")

	salas := []Sala{}
	for results.Next() {
		var sala Sala
		var err error
		err = results.Scan(&sala.ID, &sala.Nombre, &sala.Precio, &sala.Color)
		if err != nil {
			panic("Error al validar datos" + err.Error())
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
	sqlstm := fmt.Sprintf("INSERT INTO sala (nombre, precio, color)"+
		" VALUES ('%s','%g', '%d')",
		newSala.Nombre, newSala.Precio, newSala.Color)
	GetConsulta(sqlstm)

	c.IndentedJSON(http.StatusCreated, newSala)
}

func GetSalaGrupo(c *gin.Context) {
	fecha := c.Query("fecha")
	sqlstm := fmt.Sprintf("SELECT sg.id, s.nombre sala, s.color sala_color, g.nombre grupo, sg.hora_desde, sg.hora_hasta FROM sala_grupo sg "+
		"JOIN sala s ON sg.sala_id = s.id "+
		"JOIN grupo g ON sg.grupo_id = g.id "+
		"WHERE hora_desde BETWEEN '%s 00:00:00' AND '%s 23:59:59'", fecha, fecha)

	results := GetConsulta(sqlstm)

	grupos := []SalaGrupo{}
	for results.Next() {
		var salaGrupo SalaGrupo
		var err error
		err = results.Scan(&salaGrupo.ID, &salaGrupo.Sala, &salaGrupo.SalaColor, &salaGrupo.Grupo, &salaGrupo.HoraDesde, &salaGrupo.HoraHasta)
		if err != nil {
			panic("Error al validar datos " + err.Error())
		}
		grupos = append(grupos, salaGrupo)
	}
	c.IndentedJSON(http.StatusOK, grupos)
}
