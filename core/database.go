package core

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var db *gorm.DB

func newDB() (*gorm.DB, error) {
	// os.Setenv("db_host", "localhost")
	// os.Setenv("db_user", "anandkay")
	// os.Setenv("db_password", "mypostgres")
	// os.Setenv("db_name", "linkedin")
	// os.Setenv("db_port", "5432")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("db_host"), os.Getenv("db_user"), os.Getenv("db_password"), os.Getenv("db_name"), os.Getenv("db_port"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
