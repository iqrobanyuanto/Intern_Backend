package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIdBarang godoc
// @Summary Get barang from database by their id.
// @Description get every barang from database using barang id.
// @Tags GetById_BarangFunction
// @Param id path string true "BarangModel id as a key to get the BarangModel data"
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /product/get/{id} [get]
func GetByIdBarang(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDataBase()
	var barang models.BarangModel
	if err := db.Where("id=?", id).Find(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"barangModel": barang})
}

// SearchBarang godoc
// @Summary Get barang from database by their name.
// @Description get every barang from database that related to their input parameter(nama barang).
// @Tags Search_BarangFunction
// @Param nama path string true "BarangModel nama as a key to get the BarangModel data"
// @Produce json
// @Success 200 {object} []models.BarangModel
// @Router /product/get/search?nama={nama} [get]
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

// FilterBarang godoc
// @Summary Get barang from database by their category.
// @Description get every barang from database that related to their input parameter(category).
// @Tags Search_BarangFunction
// @Param kategori path string true "BarangModel kategori as a key to get the BarangModel data"
// @Produce json
// @Success 200 {object} []models.BarangModel
// @Router /product/get/filter?kategori={kategori} [get]
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
