package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add godoc
// @Summary Add Barang to database.
// @Description Insert the given barang from API to the database.
// @Tags Add_BarangFunction
// @Param Body body BarangModel true "the body to add barang to database"
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /product/add [post]
func Add(c *gin.Context) {
	var barang models.BarangModel
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	db := config.ConnectDataBase()
	db.Create(&barang)
	c.JSON(http.StatusOK, gin.H{"BarangModel": barang})
}
