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

func GetSalaGrupo(c *gin.Context) {
	fecha := c.Query("fecha")
	print(fecha)

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`SELECT sg.id, s.nombre sala, s.color sala_color, g.nombre grupo, sg.hora_desde, sg.hora_hasta 
		FROM sala_grupo sg 
		JOIN sala s ON sg.sala_id = s.id 
		JOIN grupo g ON sg.grupo_id = g.id
		WHERE hora_desde BETWEEN ? AND ?`)
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Query(fecha+" 00:00:00", fecha+" 23:59:59")
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(204)
			return
		}
		log.Fatal(err)
	}

	grupos := []SalaGrupoAgenda{}
	for result.Next() {
		var salaGrupo SalaGrupoAgenda
		var err error
		err = result.Scan(&salaGrupo.ID, &salaGrupo.Sala, &salaGrupo.SalaColor, &salaGrupo.Grupo, &salaGrupo.HoraDesde, &salaGrupo.HoraHasta)
		if err != nil {
			panic("Error al validar datos " + err.Error())
		}
		grupos = append(grupos, salaGrupo)
	}
	c.IndentedJSON(http.StatusOK, grupos)
}

func PostSalaGrupo(c *gin.Context) {
	var newSalaGrupo SalaGrupo

	if err := c.BindJSON(&newSalaGrupo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`INSERT INTO sala_grupo (grupo_id, sala_id, hora_desde, hora_hasta) VALUES (?,?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(newSalaGrupo.GrupoId, newSalaGrupo.SalaId, newSalaGrupo.HoraDesde, newSalaGrupo.HoraHasta)
	id, err := res.LastInsertId()
	newSalaGrupo.ID = int(id)
	if err != nil {
		log.Fatalf("Error al agregar Turno: %s", err)
	}

	c.String(http.StatusCreated, "Turno agregado con exito")
}
