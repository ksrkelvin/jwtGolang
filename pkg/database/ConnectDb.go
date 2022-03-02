package database

import (
	"diino/pkg/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB Conexao com o DB
var DB *gorm.DB

//ConnectDb connect to db
func ConnectDb() {
	ClientesDb := os.Getenv("DIINO")

	DbConnection, err := gorm.Open(mysql.Open(ClientesDb), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database.")
	}

	DB = DbConnection

	DbConnection.AutoMigrate(&models.User{})

}
