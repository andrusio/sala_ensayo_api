package controllers

import (
	"database/sql"
	"log"
	"net/http"
	sqldb "sala_ensayo/server/database"

	"github.com/gin-gonic/gin"
)

type Grupo struct {
	ID       int       `json:"id"`
	Nombre   string    `json:"nombre" binding:"required"`
	Personas []Persona `json:"personas"`
}

func GetGrupos(c *gin.Context) {
	db := sqldb.ConnectDB()
	results, err := db.Query(`SELECT id, nombre FROM grupo`)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(204)
			return
		}
	}

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

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`INSERT INTO grupo (nombre) VALUES (?)`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(newGrupo.Nombre)
	id, err := res.LastInsertId()
	newGrupo.ID = int(id)
	if err != nil {
		log.Fatalf("Error al agregar grupo: %s", err)
	}
	c.IndentedJSON(http.StatusCreated, newGrupo)
}

func PutGrupo(c *gin.Context) {
	var newGrupo Grupo
	if err := c.BindJSON(&newGrupo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`UPDATE grupo SET nombre = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(newGrupo.Nombre, newGrupo.ID)
	if err != nil {
		c.String(http.StatusBadRequest, "Error al actualizar grupo: %s", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, newGrupo)
}

func DeleteGrupo(c *gin.Context) {
	id := c.Param("id")
	db := sqldb.ConnectDB()
	_, err := db.Exec(`DELETE FROM grupo WHERE id = ?`, id)
	if err != nil {
		c.String(http.StatusBadRequest, "Error al eliminar grupo: %s", err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "Grupo eliminado con éxito")
}

func GetPersonasGrupoById(c *gin.Context) {
	id := c.Param("id")

	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`SELECT id, nombre FROM grupo WHERE id = ?`)

	result_grupo, err := stmt.Query(&id)

	var grupo Grupo
	for result_grupo.Next() {
		err = result_grupo.Scan(&grupo.ID, &grupo.Nombre)
		if err != nil {
			panic("Error al validar datos" + err.Error())
		}
	}

	stmt2, err := db.Prepare(`SELECT p.id, p.nombre, p.apellido, p.telefono 
		FROM persona_grupo pg
		JOIN persona p on pg.persona_id = p.id 
		WHERE grupo_id = ?`)

	result_personas, err := stmt2.Query(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(204)
			return
		}
	}

	personas := []Persona{}
	for result_personas.Next() {
		var persona Persona
		var err error
		err = result_personas.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Telefono)
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
	// grupo_id, err := strconv.Atoi(idGrupo)

	idPersona := c.Param("persona_id")
	// persona_id, err := strconv.Atoi(idPersona)

	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// }
	db := sqldb.ConnectDB()
	stmt, err := db.Prepare(`INSERT INTO persona_grupo (grupo_id, persona_id) VALUES (?,?)`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(idGrupo, idPersona)
	_ = res
	if err != nil {
		log.Fatalf("Error al agregar persona: %s", err)
	}

	c.String(http.StatusCreated, "Integrante agregado con exito")
}
