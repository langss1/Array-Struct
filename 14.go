package main

import (
	"fmt"
	"sort"
)

// Tipe data bentukan untuk menyimpan data peserta
type Peserta struct {
	Nama  string
	Asal  string
	Poin1 float64
	Poin2 float64
	Poin  float64
}

// Fungsi untuk menghitung rata-rata poin peserta
func hitungRataRata(p *Peserta) {
	p.Poin = (p.Poin1 + p.Poin2) / 2
}

// Fungsi untuk mengisi data peserta
func isiPeserta(peserta *[]Peserta) {
	for {
		var nama, asal string
		var poin1, poin2 float64
		fmt.Scan(&nama, &asal, &poin1, &poin2)
		if nama == "0" && asal == "0" && poin1 == 0 && poin2 == 0 {
			break
		}
		p := Peserta{Nama: nama, Asal: asal, Poin1: poin1, Poin2: poin2}
		hitungRataRata(&p)
		*peserta = append(*peserta, p)
	}
}

// Fungsi untuk mencetak juara 1, 2, dan 3 berdasarkan poin rata-rata
func cetakJuara(peserta []Peserta) {
	sort.Slice(peserta, func(i, j int) bool {
		return peserta[i].Poin > peserta[j].Poin
	})

	fmt.Printf("Juara 1 : %s\n", peserta[0].Nama)
	fmt.Printf("Juara 2 : %s\n", peserta[1].Nama)
	fmt.Printf("Juara 3 : %s\n", peserta[2].Nama)
}

func main() {
	var peserta []Peserta
	isiPeserta(&peserta)
	cetakJuara(peserta)
}
