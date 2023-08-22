package models

type AdminModel struct {
	kodePegawai string
	roles       string
}

func adminAccount(newKodePegawai string) AdminModel {
	return AdminModel{
		kodePegawai: newKodePegawai,
		roles:       "Admin",
	}
}
