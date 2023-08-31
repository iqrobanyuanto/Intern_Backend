package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Delete godoc
// @Summary delete barang from database.
// @Description delete barang from database by their id.
// @Tags Delete_BarangFunction
// @Param id path string true "BarangModel id as a key to delete BarangModel data"
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /product/delete/{id} [delete]
func Delete(c *gin.Context) {
	var barang models.BarangModel
	db := config.ConnectDataBase()
	id := c.Param("id")
	if db.Delete(&barang, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "id " + id + " deleted"})
}
