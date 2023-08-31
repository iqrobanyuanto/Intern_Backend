package controllers

import (
	"Intern_Backend/config"
	"Intern_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary delete barang from database.
// @Description delete barang from database by their id.
// @Tags Delete_BarangFunction
// @Param id query int true "BarangModel id as a key to delete BarangModel data"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.BarangModel
// @Router /update-product/delete [delete]
func Delete(c *gin.Context) {
	id := c.Query("id")
	var barang models.BarangModel
	db := config.ConnectDataBase()

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	if db.Delete(&barang, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "id " + id + " deleted"})
}
