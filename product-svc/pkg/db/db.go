package db

import (
	"github.com/allrested/product/product-svc/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(
		&models.Product{},
		&models.StockDecreaseLog{},
	)

	if err != nil {
		log.Fatalln(err)
	}

	return Handler{db}
}
