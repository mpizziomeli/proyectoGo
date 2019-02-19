package configuration

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//LEO EL JSON
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c

}

func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
