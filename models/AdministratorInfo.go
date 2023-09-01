package models

import (
	"Intern_Backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminModel struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	KodePegawai string `json:"kode"`
	Role        string `json:"role" gorm:"default:admin"`
}

func adminAccount(newKodePegawai string) AdminModel {
	return AdminModel{
		KodePegawai: newKodePegawai,
		Role:        "Admin",
	}
}

func (u *AdminModel) LoginCheckAdmin(db *gorm.DB) (string, error) {
	var err error

	adminFromDB := AdminModel{}
	err = db.Model(&AdminModel{}).Where("kode_pegawai = ?", u.KodePegawai).Take(&adminFromDB).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminFromDB.KodePegawai), []byte(u.KodePegawai))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(adminFromDB.ID, adminFromDB.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AdminModel) SaveAdmin(db *gorm.DB) (*AdminModel, error) {
	// Turn kode into hash
	hashedKode, err := bcrypt.GenerateFromPassword([]byte(u.KodePegawai), bcrypt.DefaultCost)
	if err != nil {
		return &AdminModel{}, err
	}
	u.KodePegawai = string(hashedKode)

	err = db.Create(&u).Error
	if err != nil {
		return &AdminModel{}, err
	}
	return u, nil
}
