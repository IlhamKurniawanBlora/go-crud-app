package controllers

import (
	"go-crud-app/database"
	"go-crud-app/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var newItem models.Item

	// Bind form data (kecuali file)
	newItem.Name = c.PostForm("name")
	newItem.Description = c.PostForm("description")

	// Konversi harga price dari string ke float64
	price := c.PostForm("price")
	priceValue, err := strconv.ParseFloat(price, 64) // Gunakan strconv.ParseFloat
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
		return
	}
	newItem.Price = priceValue

	// Handle file upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	// Simpan file ke folder uploads/
	filename := filepath.Base(file.Filename)
	filepath := "./uploads/" + filename
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Simpan path gambar ke database
	newItem.Image = "/uploads/" + filename

	// Simpan data item ke database
	if err := database.DB.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created successfully", "item": newItem})
}
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Item{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
