package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Telefono int    `json:"telefono"`
}

type sala struct {
	ID         int     `json:"id"`
	Nombre     string  `json:"nombre"`
	PrecioHora float32 `json:"precioHora"`
}

// personas slice to seed record album data.
var personas = []persona{
	{ID: 1, Nombre: "Tom", Apellido: "Morello", Telefono: 2615897845},
	{ID: 2, Nombre: "Jimmy", Apellido: "Hendrix", Telefono: 2616897433},
	{ID: 3, Nombre: "Steve", Apellido: "Vai", Telefono: 2615879878},
}

var salas = []sala{
	{ID: 1, Nombre: "Prodan", PrecioHora: 300.00},
	{ID: 2, Nombre: "Cerati", PrecioHora: 500.75},
	{ID: 3, Nombre: "Marley", PrecioHora: 255.50},
}

func main() {
	router := gin.Default()
	router.GET("/personas", getPersonas)
	router.GET("/personas/:id", getPersonasById)
	router.POST("/personas", postPersonas)

	router.GET("/salas", getSalas)
	router.POST("/salas", postSalas)

	router.Run("localhost:8080")
}

// getPersonas responds with the list of all personas as JSON.
func getPersonas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, personas)
}

// postPersonas adds an album from JSON received in the request body.
func postPersonas(c *gin.Context) {
	var newPersona persona

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPersona); err != nil {
		return
	}

	// Add the new album to the slice.
	personas = append(personas, newPersona)
	c.IndentedJSON(http.StatusCreated, newPersona)
}

// getPersonasById locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getPersonasById(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.Atoi(id)
	// Loop through the list of personas, looking for
	// an album whose ID value matches the parameter.
	if err != nil {
		for _, a := range personas {
			if a.ID == id2 {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "persona not found"})
}

func getSalas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, salas)
}

func postSalas(c *gin.Context) {
	var newSala sala

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newSala); err != nil {
		println(err)
		return
	}

	// Add the new album to the slice.
	// newSala.PrecioHora = strconv.ParseFloat(newSala.PrecioHora, 32)
	// strconv.ParseFloat(newSala.PrecioHora, 64)
	salas = append(salas, newSala)
	c.IndentedJSON(http.StatusCreated, newSala)
}
