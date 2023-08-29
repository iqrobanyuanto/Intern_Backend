package controllers

import (
	"Intern_Backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KodeInput struct {
	Kode string `json:"kode" binding:"required"`
}

// LoginAdmin godoc
// @Summary Login as as admin.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body KodeInput true "the body to login a admin"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login-admin [post]
func LoginAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input KodeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.AdminModel{}
	u.KodePegawai = input.Kode

	token, err := u.LoginCheckAdmin(db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is incorrect."})
		return
	}

	admin := map[string]string{
		"kode": u.KodePegawai,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": admin, "token": token})

}
