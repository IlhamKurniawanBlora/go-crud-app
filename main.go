package main

import (
	"go-crud-app/database"
	"go-crud-app/models"
	"go-crud-app/routes"
	"log"
)

func main() {
	// Sambungkan ke database
	database.Connect()

	// Migrasi model Item
	err := database.DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Database migration successful!")

	// Atur route dan jalankan server
	r := routes.SetupRouter()
	r.Run(":8080") // Jalankan server di port 8080
}
