package database

import (
	"log"

	"github.com/isaque-vieira2019/challenge-bravo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	strConnect := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConnect))

	if err != nil {
		log.Panic("Error ao conectar com o banco de dados")
	}

	err = DB.AutoMigrate(&models.Currency{})

	if err != nil {
		log.Panic("Error ao Criar as tabelas")
	}
}
