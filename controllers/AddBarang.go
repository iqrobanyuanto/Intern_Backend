package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
