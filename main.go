package main

import (
	. "sala_ensayo/server/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/
// https://developer.fedoraproject.org/tools/docker/docker-installation.html
// https://blog.logrocket.com/rest-api-golang-gin-gorm/

func main() {
	router := gin.Default()
	router.GET("/personas", GetPersonas)
	router.POST("/persona", PostPersona)

	router.GET("/salas", GetSalas)
	router.POST("/sala", PostSala)
	router.GET("/sala_grupo", GetSalaGrupo)

	router.GET("/grupos", GetGrupos)
	router.POST("/grupo", PostGrupo)

	router.GET("/persona_grupo/:id", GetPersonasGrupoById)
	router.POST("/persona_grupo/:grupo_id/:persona_id", PostPersonaGrupo)

	router.Run("localhost:8080")
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
