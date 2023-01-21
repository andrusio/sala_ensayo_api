package main

import (
	"fmt"
	"net/http"
	"strconv"

	. "sala_ensayo/server/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/
// https://developer.fedoraproject.org/tools/docker/docker-installation.html
// https://blog.logrocket.com/rest-api-golang-gin-gorm/

func main() {
	router := gin.Default()
	router.GET("/personas", getPersonas)
	router.POST("/persona", postPersona)

	router.GET("/salas", getSalas)
	router.POST("/sala", postSala)

	router.GET("/grupos", getGrupos)
	router.POST("/grupo", postGrupo)

	router.GET("/persona_grupo/:id", getPersonasGrupoById)
	router.POST("/persona_grupo/:grupo_id/:persona_id", postPersonaGrupo)

	router.Run("localhost:8080")
}

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
}

type Sala struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre" binding:"required"`
	Precio float32 `json:"precio" binding:"required"`
}

type Grupo struct {
	ID       int       `json:"id"`
	Nombre   string    `json:"nombre" binding:"required"`
	Personas []Persona `json:"personas"`
}

// var personas = []Persona{
// 	{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: "2615897845"},
// 	{ID: 2, Nombre: "Jimmy", Apellido: "Hendrix", Telefono: "2616897433"},
// 	{ID: 3, Nombre: "Steve", Apellido: "Vai", Telefono: "2615879878"},
// }

// var salas = []Sala{
// 	{ID: 1, Nombre: "Prodan", Precio: 300.00},
// 	{ID: 2, Nombre: "Cerati", Precio: 500.75},
// 	{ID: 3, Nombre: "Marley", Precio: 255.50},
// }

// var Grupos = []Grupo{
// 	{ID: 1, Nombre: "Soda Stereo", Personas: []Persona{{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: "2615897845"}}},
// 	{ID: 2, Nombre: "Audioslave", Personas: []Persona{{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: "2615897845"}}},
// 	{ID: 3, Nombre: "Massacre", Personas: []Persona{{ID: 2, Nombre: "Jimmy", Apellido: "Hendrix", Telefono: "2616897433"}, {ID: 3, Nombre: "Steve", Apellido: "Vai", Telefono: "2615879878"}}},
// }

func getPersonas(c *gin.Context) {
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

func postPersona(c *gin.Context) {
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

func getSalas(c *gin.Context) {
	results := GetConsulta("SELECT * FROM sala")

	salas := []Sala{}
	for results.Next() {
		var sala Sala
		var err error
		err = results.Scan(&sala.ID, &sala.Nombre, &sala.Precio)
		if err != nil {
			panic("Error al validar datos" + err.Error())
		}
		salas = append(salas, sala)
	}
	c.IndentedJSON(http.StatusOK, salas)
}

func postSala(c *gin.Context) {
	var newSala Sala
	if err := c.BindJSON(&newSala); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	sqlstm := fmt.Sprintf("INSERT INTO sala (nombre, precio)"+
		" VALUES ('%s','%g')",
		newSala.Nombre, newSala.Precio)
	GetConsulta(sqlstm)

	c.IndentedJSON(http.StatusCreated, newSala)
}

func getGrupos(c *gin.Context) {
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

func postGrupo(c *gin.Context) {
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

func getPersonasGrupoById(c *gin.Context) {
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

func postPersonaGrupo(c *gin.Context) {
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
