package sqldb

import (
	"database/sql"
	"io/ioutil"
	"log"

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

func ConnectDB() *sql.DB {

	var conf DatabaseConfig
	conf.getDBConf()

	db, err := sql.Open("mysql", conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+")/"+conf.DBName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
