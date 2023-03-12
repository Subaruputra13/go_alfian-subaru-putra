package main

// penggunaan nama struct harus diawali huruf besar dan isinya harus menggunakan PascalCase
type Kendaraan struct {
	TotalRoda      int
	KecepatanMobil int
}

// penggunaan nama struct harus diawali huruf besar
type Mobil struct {
	Kendaraan
}

// membuat variable harus jelas dan mudah dipahami, penamaan funsi harus menggunakan CamelCase
func (mobil *Mobil) tambahKecepatan(kecepatanbaru int) {
	mobil.KecepatanMobil = mobil.KecepatanMobil + kecepatanbaru
}

func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func main() {
	// Penamaan Variable harus jelas dan mudah dipahami
	MobilCepat := Mobil{Kendaraan{4, 60}}
	MobilCepat.berjalan()
	MobilCepat.berjalan()
	MobilCepat.berjalan()

	MobilLamban := Mobil{Kendaraan{4, 40}}
	MobilLamban.berjalan()
}
