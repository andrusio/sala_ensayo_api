package controllers

import (
	"database/sql"
	"net/http"
	sqldb "sala_ensayo/server/database"
	"sala_ensayo/server/helpers"

	"github.com/gin-gonic/gin"
)

type SalaGrupoAgenda struct {
	ID        int    `json:"id"`
	Sala      string `json:"sala" binding:"required"`
	SalaColor int    `json:"sala_color" binding:"required"`
	Grupo     string `json:"grupo" binding:"required"`
	HoraDesde string `json:"hora_desde" binding:"required"`
	HoraHasta string `json:"hora_hasta" binding:"required"`
}

type SalaGrupo struct {
	ID        int    `json:"id"`
	GrupoId   int    `json:"grupo_id" binding:"required"`
	SalaId    int    `json:"sala_id" binding:"required"`
	HoraDesde string `json:"hora_desde" binding:"required"`
	HoraHasta string `json:"hora_hasta" binding:"required"`
}

func GetSalaGrupo(c *gin.Context) {
	fecha := c.Query("fecha")

	db := sqldb.ConnectDB()
	res, err := db.Query(`SELECT sg.id, s.nombre sala, s.color sala_color, g.nombre grupo, sg.hora_desde, sg.hora_hasta 
		FROM sala_grupo sg 
		JOIN sala s ON sg.sala_id = s.id 
		JOIN grupo g ON sg.grupo_id = g.id
		WHERE hora_desde BETWEEN ? AND ?`, fecha+" 00:00:00", fecha+" 23:59:59")
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, gin.H{"error": "No hay turnos cargados"})
			return
		}
	}

	grupos := []SalaGrupoAgenda{}
	for res.Next() {
		var salaGrupo SalaGrupoAgenda
		var err error
		err = res.Scan(&salaGrupo.ID, &salaGrupo.Sala, &salaGrupo.SalaColor, &salaGrupo.Grupo, &salaGrupo.HoraDesde, &salaGrupo.HoraHasta)
		if err != nil {
			helpers.Error.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error al validar datos: " + err.Error()})
			return
		}
		grupos = append(grupos, salaGrupo)
	}
	c.IndentedJSON(http.StatusOK, grupos)
}

func PostSalaGrupo(c *gin.Context) {
	var newSalaGrupo SalaGrupo
	if err := c.BindJSON(&newSalaGrupo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos: " + err.Error()})
		return
	}

	db := sqldb.ConnectDB()
	_, err := db.Exec(`INSERT INTO sala_grupo (grupo_id, sala_id, hora_desde, hora_hasta) VALUES (?,?,?,?)`, newSalaGrupo.GrupoId, newSalaGrupo.SalaId, newSalaGrupo.HoraDesde, newSalaGrupo.HoraHasta)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Turno: " + err.Error()})
		return
	}

	// res, err := stmt.Exec(newSalaGrupo.GrupoId, newSalaGrupo.SalaId, newSalaGrupo.HoraDesde, newSalaGrupo.HoraHasta)
	// id, err := res.LastInsertId()
	// newSalaGrupo.ID = int(id)
	// if err != nil {
	// 	log.Fatalf("Error al agregar Turno: %s", err)
	// }

	c.JSON(http.StatusCreated, gin.H{"status": "Turno agregado con éxito"})
}
