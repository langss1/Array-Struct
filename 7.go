package main

import (
	"fmt"
)

type Pengeluaran struct {
	NamaBarang  string
	HargaBarang int
	TanggalBeli string
}

func main() {
	var n int
	fmt.Scan(&n)

	var pengeluaran [100]Pengeluaran
	for i := 0; i < n; i++ {
		fmt.Scan(&pengeluaran[i].NamaBarang, &pengeluaran[i].HargaBarang, &pengeluaran[i].TanggalBeli)
	}

	var tanggalCari string

	fmt.Scan(&tanggalCari)

	fmt.Printf("Daftar Pengeluaran pada tanggal: %s\n", tanggalCari)
	for i := 0; i < n; i++ {
		if pengeluaran[i].TanggalBeli == tanggalCari {
			fmt.Printf("%s %d\n", pengeluaran[i].NamaBarang, pengeluaran[i].HargaBarang)
		}
	}
}
