package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/

func main() {
	router := gin.Default()
	router.GET("/personas", getPersonas)
	router.GET("/persona/:id", getPersonaById)
	router.POST("/persona", postPersona)

	router.GET("/salas", getSalas)
	router.POST("/sala", postSala)

	router.GET("/grupos", getGrupos)
	router.GET("/grupo/:id", getGrupoById)
	router.POST("/grupo", postGrupo)

	router.POST("/grupo/agregar_integrante/:id/:id", postGrupoIntegrante)

	router.Run("localhost:8080")
}

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido"`
	Telefono int    `json:"telefono"`
}

type Sala struct {
	ID         int     `json:"id"`
	Nombre     string  `json:"nombre" binding:"required"`
	PrecioHora float32 `json:"precioHora" binding:"required"`
}

type Grupo struct {
	ID       int       `json:"id"`
	Nombre   string    `json:"nombre" binding:"required"`
	Personas []Persona `json:"personas"`
}

var personas = []Persona{
	{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: 2615897845},
	{ID: 2, Nombre: "Jimmy", Apellido: "Hendrix", Telefono: 2616897433},
	{ID: 3, Nombre: "Steve", Apellido: "Vai", Telefono: 2615879878},
}

var salas = []Sala{
	{ID: 1, Nombre: "Prodan", PrecioHora: 300.00},
	{ID: 2, Nombre: "Cerati", PrecioHora: 500.75},
	{ID: 3, Nombre: "Marley", PrecioHora: 255.50},
}

var Grupos = []Grupo{
	{ID: 1, Nombre: "Soda Stereo", Personas: []Persona{{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: 2615897845}}},
	{ID: 2, Nombre: "Audioslave", Personas: []Persona{{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: 2615897845}}},
	{ID: 3, Nombre: "Massacre", Personas: []Persona{{ID: 2, Nombre: "Jimmy", Apellido: "Hendrix", Telefono: 2616897433}, {ID: 3, Nombre: "Steve", Apellido: "Vai", Telefono: 2615879878}}},
}

func getPersonas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, personas)
}

func postPersona(c *gin.Context) {
	var newPersona Persona
	if err := c.BindJSON(&newPersona); err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest,
		// 	gin.H{
		// 		"error": "VALIDATEERR-1",
		// 		"message": "Invalid inputs. Please check your inputs"})
		// 	return
		c.AbortWithError(http.StatusBadRequest, err)
	}

	personas = append(personas, newPersona)
	c.IndentedJSON(http.StatusCreated, newPersona)
}

func getPersonaById(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		for _, a := range personas {
			if a.ID == id2 {
				c.IndentedJSON(http.StatusOK, a)
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func getSalas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, salas)
}

func postSala(c *gin.Context) {
	var newSala Sala
	if err := c.BindJSON(&newSala); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	salas = append(salas, newSala)
	c.IndentedJSON(http.StatusCreated, newSala)
}

func getGrupos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Grupos)
}

func postGrupo(c *gin.Context) {
	var newGrupo Grupo
	if err := c.BindJSON(&newGrupo); err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest,
		// 	gin.H{
		// 		"error": "VALIDATEERR-1",
		// 		"message": "Invalid inputs. Please check your inputs"})
		// 	return
		c.AbortWithError(http.StatusBadRequest, err)
	}

	Grupos = append(Grupos, newGrupo)
	c.IndentedJSON(http.StatusCreated, newGrupo)
}

func getGrupoById(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		for _, a := range Grupos {
			if a.ID == id2 {
				c.IndentedJSON(http.StatusOK, a)
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Grupo not found"})
}

// NO FUNCIONAL
func postGrupoIntegrante(c *gin.Context) {
	var newGrupo Grupo
	if err := c.BindJSON(&newGrupo); err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest,
		// 	gin.H{
		// 		"error": "VALIDATEERR-1",
		// 		"message": "Invalid inputs. Please check your inputs"})
		// 	return
		c.AbortWithError(http.StatusBadRequest, err)
	}

	Grupos = append(Grupos, newGrupo)
	c.IndentedJSON(http.StatusCreated, newGrupo)
}
