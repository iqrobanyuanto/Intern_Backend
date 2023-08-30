package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateBarang godoc
// @Summary Update Barang to database.
// @Description Update the given barang from API to the database by their id.
// @Tags Update_BarangFunction
// @Param id query int true "BarangModel id as a path to update related BarangModel data"
// @Param Body body models.BarangModel true "the body to update barang to database"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /update-product [put]
func UpdateBarang(c *gin.Context) {
	id := c.Query("id")
	var barang models.BarangModel
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db := config.ConnectDataBase()
	if db.Model(&barang).Where("id = ?", id).Updates(&barang).RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"messages": "Update barang gagal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": "Update berhasil"})
}
