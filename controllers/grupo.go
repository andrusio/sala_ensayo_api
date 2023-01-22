package controllers

import (
	"fmt"
	"net/http"
	. "sala_ensayo/server/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Grupo struct {
	ID       int       `json:"id"`
	Nombre   string    `json:"nombre" binding:"required"`
	Personas []Persona `json:"personas"`
}

func GetGrupos(c *gin.Context) {
	results := GetConsulta("SELECT * FROM grupo")

	grupos := []Grupo{}
	for results.Next() {
		var grupo Grupo
		var err error
		err = results.Scan(&grupo.ID, &grupo.Nombre)
		if err != nil {
			panic("Error al validar datos" + err.Error())
		}
		grupos = append(grupos, grupo)
	}
	c.IndentedJSON(http.StatusOK, grupos)
}

func PostGrupo(c *gin.Context) {
	var newGrupo Grupo
	if err := c.BindJSON(&newGrupo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	sqlstm := fmt.Sprintf("INSERT INTO grupo (nombre)"+
		" VALUES ('%s')",
		newGrupo.Nombre)
	GetConsulta(sqlstm)

	c.IndentedJSON(http.StatusCreated, newGrupo)
}

func GetPersonasGrupoById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	sqlstm := fmt.Sprintf("SELECT * FROM grupo"+
		" WHERE id = '%d'",
		idInt)
	result := GetConsulta(sqlstm)

	var grupo Grupo
	for result.Next() {
		err = result.Scan(&grupo.ID, &grupo.Nombre)
		if err != nil {
			panic("Error al validar datos" + err.Error())
		}
	}

	sqlstm2 := fmt.Sprintf("SELECT p.id, p.nombre, p.apellido, p.telefono "+
		"FROM persona_grupo pg "+
		"JOIN persona p on pg.persona_id = p.id "+
		"WHERE grupo_id = '%d'", idInt)
	results := GetConsulta(sqlstm2)

	personas := []Persona{}
	for results.Next() {
		var persona Persona
		var err error
		err = results.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Telefono)
		if err != nil {
			panic("Error al validar datos" + err.Error())
		}
		personas = append(personas, persona)
	}

	grupo.Personas = personas

	c.IndentedJSON(http.StatusOK, grupo)
}

func PostPersonaGrupo(c *gin.Context) {
	idGrupo := c.Param("grupo_id")
	grupo_id, err := strconv.Atoi(idGrupo)

	idPersona := c.Param("persona_id")
	persona_id, err := strconv.Atoi(idPersona)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	sqlstm := fmt.Sprintf("INSERT INTO persona_grupo (grupo_id, persona_id)"+
		" VALUES ('%d','%d')",
		grupo_id, persona_id)

	GetConsulta(sqlstm)
	c.String(http.StatusCreated, "Integrante agregado con exito")
}
