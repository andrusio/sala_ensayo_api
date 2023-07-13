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

	// Log General App
	// gin.DisableConsoleColor()
	// f, err := os.Create("log/gin.log")
	// if err != nil {
	// 	fmt.Println("Open Log File Failed", err)
	// }
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/personas", GetPersonas)
	router.POST("/persona", PostPersona)
	router.PUT("/persona", PutPersona)
	router.DELETE("/persona/:id", DeletePersona)

	router.GET("/salas", GetSalas)
	router.POST("/sala", PostSala)
	router.PUT("/sala", PutSala)
	router.DELETE("/sala/:id", DeleteSala)

	router.GET("/grupos", GetGrupos)
	router.POST("/grupo", PostGrupo)
	router.PUT("/grupo", PutGrupo)
	router.DELETE("/grupo/:id", DeleteGrupo)

	router.GET("/sala_grupo", GetSalaGrupo)
	router.POST("/sala_grupo", PostSalaGrupo)

	router.GET("/persona_grupo/:id", GetPersonasGrupoById)
	router.POST("/persona_grupo/:grupo_id/:persona_id", PostPersonaGrupo)

	router.Run("localhost:8080")
}
