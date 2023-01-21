package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	DBName     string `yaml:"dbname"`
	DBUser     string `yaml:"dbuser"`
	DBPassword string `yaml:"dbpassword"`
	DBHost     string `yaml:"dbhost"`
}

func (conf *DatabaseConfig) getDBConf() *DatabaseConfig {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return conf
}

func GetConsulta(sqlstm string) sql.Rows {
	var conf DatabaseConfig
	conf.getDBConf()

	db, err := sql.Open("mysql", conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+")/"+conf.DBName)
	if err != nil {
		fmt.Println("Error al conectar con la base de datos", err.Error())
	}

	defer db.Close()
	results, err := db.Query(sqlstm)
	if err != nil {
		fmt.Println("Error al realizar consulta", err.Error())
	}

	return *results
}
