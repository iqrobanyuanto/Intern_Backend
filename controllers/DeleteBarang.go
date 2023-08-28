package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Delete godoc
// @Summary delete barang from database.
// @Description delete barang from database by their id.
// @Tags Delete_BarangFunction
// @Param id path string true "BarangModel id as a key to delete BarangModel data"
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /product/delete [delete]
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
