package models

import (
	"Intern_Backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminModel struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	KodePegawai string `json:"kode"`
	Roles       string `json:"roles" gorm:"default:admin"`
}

func adminAccount(newKodePegawai string) AdminModel {
	return AdminModel{
		KodePegawai: newKodePegawai,
		Roles:       "Admin",
	}
}

func VerifyKodeAdmin(kode, hashedkode string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedkode), []byte(kode))
}

func LoginCheckAdmin(kode string, db *gorm.DB) (string, error) {

	var err error

	u := AdminModel{}

	err = VerifyKodeAdmin(kode, u.KodePegawai)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID, u.Roles)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *AdminModel) SaveAdmin(db *gorm.DB) (*AdminModel, error) {
	//turn password into hash
	hashedKode, errKode := bcrypt.GenerateFromPassword([]byte(u.KodePegawai), bcrypt.DefaultCost)
	if errKode != nil {
		return &AdminModel{}, errKode
	}
	u.KodePegawai = string(hashedKode)

	var err error = db.Create(&u).Error
	if err != nil {
		return &AdminModel{}, err
	}
	return u, nil
}
