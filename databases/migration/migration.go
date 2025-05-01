package main

import (
	"fmt"
	"sushi-api/config"
	"sushi-api/databases"
	"sushi-api/entities"

	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.Connect())

	tx := db.Connect().Begin()

	customerMigration(tx)
	categoryMigration(tx)
	menuMigration(tx)
	orderMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

}

func customerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Customer{})
}

func categoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Category{})
}

func menuMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Menu{})
}

func orderMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Order{})
}
