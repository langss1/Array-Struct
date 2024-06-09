package main

import (
	"fmt"
)

// Deklarasi tipe bentukan Mahasiswa
type Mahasiswa struct {
	Nama string
	NIM  int
}

// Deklarasi tipe bentukan tabMhs
type tabMhs struct {
	Data [40]Mahasiswa
	Neff int
}

func main() {
	var students tabMhs
	var inputNama string
	var inputNIM int

	// Membaca input dari pengguna
	for students.Neff < 40 {
		fmt.Scan(&inputNama)
		if inputNama == "#" {
			break
		}
		fmt.Scan(&inputNIM)

		students.Data[students.Neff] = Mahasiswa{Nama: inputNama, NIM: inputNIM}
		students.Neff++
	}

	// Membaca input untuk menentukan ganjil atau genap
	var filter string
	fmt.Scan(&filter)

	// Mencetak nama mahasiswa berdasarkan NIM ganjil atau genap
	for i := 0; i < students.Neff; i++ {
		if filter == "ganjil" && students.Data[i].NIM%2 != 0 {
			fmt.Print(students.Data[i].Nama, " ")
		} else if filter == "genap" && students.Data[i].NIM%2 == 0 {
			fmt.Print(students.Data[i].Nama, " ")
		}
	}
	fmt.Println()
}
