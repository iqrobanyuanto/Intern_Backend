package models

type BarangModel struct {
	id             int
	namaBarang     string
	jumlahBarang   int
	hargaBarang    int
	kategoriBarang string
}

func NewBarang(id_ int, nama, kategori string, jumlah, harga int) BarangModel {
	return BarangModel{
		id:             id_,
		namaBarang:     nama,
		jumlahBarang:   jumlah,
		hargaBarang:    harga,
		kategoriBarang: kategori,
	}
}
