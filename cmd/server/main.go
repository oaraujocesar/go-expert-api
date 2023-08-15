package main

import (
	"fmt"
	"net/http"

	"github.com/oaraujocesar/go-expert-api/configs"
	"github.com/oaraujocesar/go-expert-api/internal/entity"
	"github.com/oaraujocesar/go-expert-api/internal/infra/database"
	"github.com/oaraujocesar/go-expert-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig("./cmd/server")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	fmt.Println("Server running on http://localhost:8000...")
	http.ListenAndServe(":8000", nil)
}
