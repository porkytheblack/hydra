package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

func NewConnection ( config *Config ) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s`,
		config.Host, config.Port, config.User, config.Password, config.DBName, 
	)

	fmt.Printf("\n\nDSN IS:: %s\n\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("\n\nAn error occured::%s\n\n", err)
		return nil, err
	}

	log.Printf("\n\nConnected to DB Successfully!!\n\n");

	return db, nil
}