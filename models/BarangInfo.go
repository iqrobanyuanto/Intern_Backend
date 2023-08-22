package models

type BarangModel struct {
	namaBarang     string
	jumlahBarang   int
	hargaBarang    int
	kategoriBarang string
}

func NewBarang(nama, kategori string, jumlah, harga int) BarangModel {
	return BarangModel{
		namaBarang:     nama,
		jumlahBarang:   jumlah,
		hargaBarang:    harga,
		kategoriBarang: kategori,
	}
}
