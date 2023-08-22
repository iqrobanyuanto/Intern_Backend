package models

type BarangModel struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	NamaBarang     string `json:"nama"`
	JumlahBarang   int    `json:"jumlah"`
	HargaBarang    int    `json:"harga"`
	KategoriBarang string `json:"kategori"`
}

func NewBarang(nama, kategori string, jumlah, harga int) BarangModel {
	return BarangModel{
		NamaBarang:     nama,
		JumlahBarang:   jumlah,
		HargaBarang:    harga,
		KategoriBarang: kategori,
	}
}
