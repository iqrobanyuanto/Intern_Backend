package models

import (
	"html"
	"strings"

	"Intern_Backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ManagerModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Telepon  string `json:"telp"`
	Alamat   string `json:"alamat"`
	Role     string `json:"role" gorm:"default:manager"`
}

func NewManagerAccount(newusername, newpassword, newtelepon, newalamat string) ManagerModel {
	return ManagerModel{
		Username: newusername,
		Password: newpassword,
		Telepon:  newtelepon,
		Alamat:   newalamat,
		Role:     "Manager_Toko",
	}
}

func VerifyPasswordManager(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheckManager(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := ManagerModel{}

	err = db.Model(ManagerModel{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPasswordManager(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateTokenManager(u.ID, u.Role, u.Username, u.Telepon)
	if err != nil {
		return "", err
	}

	return token, nil

}

func IsUsernameTaken(db *gorm.DB, username string) bool {
	var count int64
	db.Model(&ManagerModel{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (u *ManagerModel) SaveManager(db *gorm.DB) (*ManagerModel, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &ManagerModel{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &ManagerModel{}, err
	}
	return u, nil
}
