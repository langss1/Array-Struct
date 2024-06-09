package main

import (
	"fmt"
)

// Deklarasi tipe bentukan Mahasiswa
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
}

// Deklarasi tipe bentukan tabMhs
type tabMhs struct {
	Data [40]Mahasiswa
	Neff int
}

func main() {
	var students tabMhs
	var inputNama, inputNIM, inputJurusan string

	// Membaca input dari pengguna
	for students.Neff < 40 {
		fmt.Scan(&inputNama)
		if inputNama == "#" {
			break
		}
		fmt.Scan(&inputNIM)
		fmt.Scan(&inputJurusan)

		students.Data[students.Neff] = Mahasiswa{Nama: inputNama, NIM: inputNIM, Jurusan: inputJurusan}
		students.Neff++
	}

	// Membaca input untuk menentukan ganjil atau genap
	var filter string
	fmt.Scan(&filter)

	// Mencetak nama mahasiswa berdasarkan indeks array ganjil atau genap
	for i := 0; i < students.Neff; i++ {
		if filter == "genap" && (i+1)%2 != 0 {
			fmt.Println(students.Data[i].Nama, students.Data[i].NIM, students.Data[i].Jurusan)
		} else if filter == "ganjil" && (i+1)%2 == 0 {
			fmt.Println(students.Data[i].Nama, students.Data[i].NIM, students.Data[i].Jurusan)
		}
	}
}
