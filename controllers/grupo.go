package controllers

import (
	"database/sql"
	"net/http"
	sqldb "sala_ensayo/server/database"
	"sala_ensayo/server/helpers"

	"github.com/gin-gonic/gin"
)

type Grupo struct {
	ID       int64     `json:"id"`
	Nombre   string    `json:"nombre" binding:"required"`
	Personas []Persona `json:"personas"`
}

func GetGrupos(c *gin.Context) {
	db := sqldb.ConnectDB()
	results, err := db.Query(`SELECT id, nombre FROM grupo`)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, gin.H{"error": "No hay grupos cargados"})
			return
		}
	}

	grupos := []Grupo{}
	for results.Next() {
		var grupo Grupo
		var err error
		err = results.Scan(&grupo.ID, &grupo.Nombre)
		if err != nil {
			helpers.Error.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error al validar datos: " + err.Error()})
			return
		}
		grupos = append(grupos, grupo)
	}
	c.IndentedJSON(http.StatusOK, grupos)
}

func PostGrupo(c *gin.Context) {
	var newGrupo Grupo
	if err := c.BindJSON(&newGrupo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos: " + err.Error()})
		return
	}

	db := sqldb.ConnectDB()
	res, err := db.Exec(`INSERT INTO grupo (nombre) VALUES (?)`, newGrupo.Nombre)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Grupo: " + err.Error()})
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Grupo: " + err.Error()})
		return
	}
	newGrupo.ID = id
	c.IndentedJSON(http.StatusCreated, newGrupo)
}

func PutGrupo(c *gin.Context) {
	var newGrupo Grupo
	if err := c.BindJSON(&newGrupo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos: " + err.Error()})
		return
	}

	db := sqldb.ConnectDB()
	res, err := db.Exec(`UPDATE grupo SET nombre = ? WHERE id = ?`, newGrupo.Nombre, newGrupo.ID)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al modificar Grupo: " + err.Error()})
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al modificar Grupo: " + err.Error()})
		return
	}

	newGrupo.ID = id
	c.IndentedJSON(http.StatusOK, newGrupo)
}

func DeleteGrupo(c *gin.Context) {
	id := c.Param("id")
	db := sqldb.ConnectDB()
	_, err := db.Exec(`DELETE FROM grupo WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error al eliminar Grupo: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, "Grupo eliminado con éxito")
}

func GetPersonasGrupoById(c *gin.Context) {
	id := c.Param("id")

	db := sqldb.ConnectDB()
	grupos, err := db.Query(`SELECT id, nombre FROM grupo WHERE id = ?`, &id)

	var grupo Grupo
	for grupos.Next() {
		err = grupos.Scan(&grupo.ID, &grupo.Nombre)
		if err != nil {
			helpers.Error.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error al validar datos: " + err.Error()})
			return
		}
	}

	integrantes, err := db.Query(`SELECT p.id, p.nombre, p.apellido, p.telefono 
		FROM persona_grupo pg
		JOIN persona p on pg.persona_id = p.id 
		WHERE grupo_id = ?`, &id)

	personas := []Persona{}
	for integrantes.Next() {
		var persona Persona
		var err error
		err = integrantes.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Telefono)
		if err != nil {
			helpers.Error.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error al validar datos: " + err.Error()})
			return
		}
		personas = append(personas, persona)
	}

	grupo.Personas = personas

	c.IndentedJSON(http.StatusOK, grupo)
}

func PostPersonaGrupo(c *gin.Context) {
	idGrupo := c.Param("grupo_id")
	// _, err := strconv.Atoi(idGrupo)

	idPersona := c.Param("persona_id")
	// _, err := strconv.Atoi(idPersona)

	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// }

	db := sqldb.ConnectDB()
	_, err := db.Exec(`INSERT INTO persona_grupo (grupo_id, persona_id) VALUES (?,?)`, idGrupo, idPersona)
	if err != nil {
		helpers.Error.Println(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Integrante: " + err.Error()})
		return
	}

	// id, err := res.LastInsertId()
	// if err != nil || id == 0 {
	// 	helpers.Error.Println(err.Error())
	// 	c.JSON(http.StatusConflict, gin.H{"error": "Error al agregar Integrante: " + err.Error()})
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{"status": "Integrante agregado con éxito"})
}
