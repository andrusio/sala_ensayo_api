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
	router.POST("/sala_grupo", PostSalaGrupo)

	router.GET("/grupos", GetGrupos)
	router.POST("/grupo", PostGrupo)

	router.GET("/persona_grupo/:id", GetPersonasGrupoById)
	router.POST("/persona_grupo/:grupo_id/:persona_id", PostPersonaGrupo)

	router.Run("localhost:8080")
}
