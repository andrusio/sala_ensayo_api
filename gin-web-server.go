package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func main() {
	router := gin.Default()
	router.GET("/personas", getPersonas)
	router.GET("/persona/:id", getPersonaById)
	router.POST("/persona", postPersona)

	router.GET("/salas", getSalas)
	router.POST("/sala", postSala)

	router.Run("localhost:8080")
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
