package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchBarang(c *gin.Context) {
	nama := c.Query("nama")
	db := config.ConnectDataBase()
	var baranglist []models.BarangModel
	if err := db.Where("nama_barang LIKE ?", nama+"%").Find(&baranglist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"barangModel": baranglist})
}

func FilterBarang(c *gin.Context) {
	filters := c.QueryArray("kategori")
	db := config.ConnectDataBase()
	var kategorilist []models.BarangModel
	query := db.Model(&kategorilist)
	if kategori := filters; len(kategori) > 0 {
		query = query.Where("kategori_barang IN ?", kategori)
	}
	query.Find(&kategorilist)
	c.JSON(http.StatusOK, gin.H{"barangModel": kategorilist})
}
