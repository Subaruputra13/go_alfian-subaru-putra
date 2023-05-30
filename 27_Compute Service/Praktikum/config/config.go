package config

import (
	"fmt"
	"praktikum/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "doadmin",
		DB_Password: "AVNS_6OAAT5XU-qZFKRPpW5R",
		DB_Port:     "25060",
		DB_Host:     "testing-do-user-13944740-0.b.db.ondigitalocean.com",
		DB_Name:     "testing_golang",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Blog{})

}
