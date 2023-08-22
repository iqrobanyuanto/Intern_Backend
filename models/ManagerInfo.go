package models

type ManagerModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Telepon  string `json:"telp"`
	Alamat   string `json:"alamat"`
	Roles    string `json:"roles" gorm:"default:manager"`
}

func NewManagerAccount(newusername, newpassword, newtelepon, newalamat string) ManagerModel {
	return ManagerModel{
		Username: newusername,
		Password: newpassword,
		Telepon:  newtelepon,
		Alamat:   newalamat,
		Roles:    "Manager_Toko",
	}
}
