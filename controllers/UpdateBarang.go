package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateBarang(c *gin.Context) {
	id := c.Param("id")
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
