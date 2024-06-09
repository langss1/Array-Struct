package main

import (
	"fmt"
)

// Definisikan tipe data untuk Karyawan
type Karyawan struct {
	ID           int
	Nama         string
	Posisi       string
	Gaji         float64
	NilaiKinerja float64
	TahunKerja   int
	Pendidikan   string
	Usia         int
}

// Fungsi untuk menginput data karyawan
func input(Tab *[]Karyawan, n int) {
	for i := 0; i < n; i++ {
		var k Karyawan
		fmt.Scan(&k.ID, &k.Nama, &k.Posisi, &k.Gaji, &k.NilaiKinerja, &k.TahunKerja, &k.Pendidikan, &k.Usia)
		*Tab = append(*Tab, k)
	}
}

// Fungsi untuk menampilkan karyawan yang layak mendapatkan kenaikan jabatan
func tampilkan(Tab []Karyawan) {
	for _, k := range Tab {
		if k.NilaiKinerja > 8 && k.TahunKerja > 1 && k.Pendidikan == "sarjana" && k.Usia > 28 {
			// Hitung gaji setelah kenaikan jabatan (10% kenaikan)
			gajiSetelahKenaikan := int(k.Gaji * 1.10)
			fmt.Printf("%d %s %d\n", k.ID, k.Nama, gajiSetelahKenaikan)
		}
	}
}

func main() {
	var Tab []Karyawan
	var n int
	fmt.Scan(&n)
	input(&Tab, n)
	tampilkan(Tab)
}
