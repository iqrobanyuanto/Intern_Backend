package controllers

import (
	"Intern_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Add godoc
// @Summary Add Barang to database.
// @Description Insert the given barang from API to the database.
// @Tags Add_BarangFunction
// @Param Body body models.BarangModel false "the body to add barang to database"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /update-product/add [post]
func Add(c *gin.Context) {
	var barang models.BarangModel
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Check if the product with the same name (case-insensitive) exists in the database
	db := c.MustGet("db").(*gorm.DB)
	var existingProduk models.BarangModel
	if err := db.Where("nama_barang like ?", barang.NamaBarang+"%").First(&existingProduk).Error; err == nil {
		// Product with the same name exists, update the "jumlah"
		existingProduk.JumlahBarang += barang.JumlahBarang
		db.Save(&existingProduk)
		c.JSON(http.StatusOK, gin.H{"message": "Jumlah updated successfully", "data": existingProduk})
		return
	}

	db.Create(&barang)
	c.JSON(http.StatusOK, gin.H{"BarangModel": barang})
}
