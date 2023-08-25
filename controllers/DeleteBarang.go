package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Delete(c *gin.Context) {
	var barang models.BarangModel
	db := config.ConnectDataBase()
	id := map[string]int{"id": 0}
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if db.Delete(&barang, id["id"]).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "id " + strconv.Itoa(id["id"]) + " deleted"})
}
