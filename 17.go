package main

import (
	"fmt"
)

// Tipe data bentukan untuk menyimpan resep masakan
type Resep struct {
	Nama  string
	Bahan [5]string
}

// Tipe data bentukan untuk menyimpan daftar resep masakan
type DaftarResep struct {
	Data [10]Resep
	N    int
}

// Fungsi untuk mengisi array dengan data resep masakan
func isiResep(resep *DaftarResep) {
	for {
		var nama string
		fmt.Scan(&nama)
		if nama == "#" || resep.N >= 10 {
			break
		}
		resep.Data[resep.N].Nama = nama
		for i := 0; i < 5; i++ {
			fmt.Scan(&resep.Data[resep.N].Bahan[i])
		}
		resep.N++
	}
}

// Fungsi untuk mencetak daftar resep masakan
func cetakResep(resep DaftarResep) {
	for i := 0; i < resep.N; i++ {
		fmt.Printf("Resep %d : %s\n", i+1, resep.Data[i].Nama)
		fmt.Printf("Bahan : %s %s %s %s %s\n", resep.Data[i].Bahan[0], resep.Data[i].Bahan[1], resep.Data[i].Bahan[2], resep.Data[i].Bahan[3], resep.Data[i].Bahan[4])
	}
}

func main() {
	var resep DaftarResep
	isiResep(&resep)
	cetakResep(resep)
}
