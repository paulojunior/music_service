package database

import (
	"log"
	"time"

	"github.com/paulojunior/code-challange/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var counts int

func NewPostgresDatabase() *gorm.DB {
	return ConnectToDB()
}

func ConnectToDB() *gorm.DB {
	dsn := viper.GetString("postgresql.dsn")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready... %s", err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	connection.AutoMigrate(&entity.Music{})

	return connection, nil
}
