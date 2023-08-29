package models

type BarangModel struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	NamaBarang     string `json:"nama" gorm:"type:varchar(300)"`
	JumlahBarang   int    `json:"jumlah" gorm:"type:INT"`
	HargaBarang    int    `json:"harga" gorm:"type:INT"`
	KategoriBarang string `json:"kategori" gorm:"type:varchar(300)"`
}

func NewBarang(nama, kategori string, jumlah, harga int) BarangModel {
	return BarangModel{
		NamaBarang:     nama,
		JumlahBarang:   jumlah,
		HargaBarang:    harga,
		KategoriBarang: kategori,
	}
}
