package models

type ManagerModel struct {
	username string
	password string
	telepon  string
	alamat   string
	roles    string
}

func NewManagerAccount(newusername, newpassword, newtelepon, newalamat string) ManagerModel {
	return ManagerModel{
		username: newusername,
		password: newpassword,
		telepon:  newtelepon,
		alamat:   newalamat,
		roles:    "Manager_Toko",
	}
}
