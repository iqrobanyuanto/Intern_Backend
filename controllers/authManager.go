package controllers

import (
	"Intern_Backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Telp     string `json:"telp" binding:"required"`
	Alamat   string `json:"alamat" binding:"required"`
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access manager api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a manager"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login-manager [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.ManagerModel{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheckManager(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	err = db.Model(models.ManagerModel{}).Select("role", "telepon").Where("username = ?", u.Username).Take(&u).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user role"})
		return
	}

	c.Set("userRole", u.Role)

	user := map[string]string{
		"username": u.Username,
		"telepon":  u.Telepon,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a manager"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if models.IsUsernameTaken(db, input.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is already taken"})
		return
	}

	u := models.ManagerModel{}

	u.Username = input.Username
	u.Password = input.Password
	u.Telepon = input.Telp
	u.Alamat = input.Alamat

	_, err := u.SaveManager(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
		"telepon":  input.Telp,
		"alamat":   input.Alamat,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}
