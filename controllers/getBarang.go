package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetByIdBarang godoc
// @Summary Get barang from database by their id.
// @Description get every barang from database using barang id.
// @Tags GetById_BarangFunction
// @Param id path string true "BarangModel id as a key to get the BarangModel data"
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /get-product [get]
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
// @Param nama query string false "BarangModel nama as a key to get the BarangModel data"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.BarangModel
// @Router /get-product/search [get]
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
// @Param kategori query string false "BarangModel kategori as a key to get the BarangModel data"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.BarangModel
// @Router /get-product/filter [get]
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
