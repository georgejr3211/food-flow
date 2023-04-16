package gormconfig

import (
	"fmt"

	"github.com/georgejr3211/food-flow/config"
	"github.com/georgejr3211/food-flow/internal/module/product/infrastructure/database/gorm/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	config := config.NewConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	SyncDB(db)

	return db
}

func SyncDB(db *gorm.DB) {
	if config.NewConfig().DBSync {
		fmt.Println("Syncing DB...")
		db.AutoMigrate(&model.ProductModel{})
	}
}
