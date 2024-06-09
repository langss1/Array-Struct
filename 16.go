package main

import (
	"fmt"
)

// Definisikan tipe data untuk buku
type Buku struct {
	Judul     string
	Pengarang string
}

// Definisikan tipe bentukan tabBuku berisi array buku berkapasitas 20
type tabBuku [20]Buku

// Fungsi untuk menginput data buku
func input(Tab *tabBuku) {
	var judul, pengarang string
	for i := 0; i < 20; i++ {
		fmt.Scan(&judul)
		if judul == "#" {
			break
		}
		fmt.Scan(&pengarang)
		(*Tab)[i] = Buku{Judul: judul, Pengarang: pengarang}
	}
}

// Fungsi untuk menampilkan judul buku dan nama pengarang dengan aturan LIFO
func tampilkan(Tab tabBuku) {
	for i := len(Tab) - 1; i >= 0; i-- {
		if Tab[i].Judul != "" {
			fmt.Printf("%s %s\n", Tab[i].Judul, Tab[i].Pengarang)
		}
	}
}

func main() {
	var Tab tabBuku
	input(&Tab)
	tampilkan(Tab)
}
