package config

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgres() (*gorm.DB, error) {
	host := Env.PostgresHost
	user := Env.PostgresUser
	password := Env.PostgresPassword
	dbName := Env.PostgresDbName
	port := Env.PostgresPort

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}
	log.Println("SUCCESS CONNECT TO DATABASE")

	// Migrate Table
	err = migratePostgresTable(db)
	if err != nil {
		log.Println("ERROR MIGRATE TABLE")
	}

	log.Println("SUCCESS MIGRATE TABLE")

	return db, nil
}

func migratePostgresTable(db *gorm.DB) error {
	models := []interface{}{
		&schema.User{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
