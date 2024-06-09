package main

import (
	"fmt"
)

// Struct untuk menyimpan data siswa
type Siswa struct {
	Nama   string
	Berat  float64
	Tinggi float64
	bmi    float64
}

type tsiswa [100]Siswa

func cetak(a tsiswa, n int) {
	var b int
	for i := 0; i < n; i++ {
		if a[i].bmi < 18.5 {
			b = 1
			fmt.Println("Siswa yang berat badannya kurang : ")
			for z := 0; z < n; z++ {
				if a[z].bmi <= 18.5 {
					fmt.Println(b, a[z].Nama)
					b += 1
				}
			}
			break
		}
	}

	for i := 0; i < n; i++ {
		if a[i].bmi >= 18.5 && a[i].bmi < 25.0 {
			b = 1
			fmt.Println("Siswa yang berat badannya normal : ")
			for z := 0; z < n; z++ {
				if a[z].bmi >= 18.5 && a[z].bmi < 25.0 {
					fmt.Println(b, a[z].Nama)
					b++
				}
			}
			break
		}
	}

	for i := 0; i < n; i++ {
		if a[i].bmi >= 25.0 {
			b = 1
			fmt.Println("Siswa yang berat badannya lebih : ")
			for z := 0; z < n; z++ {
				if a[z].bmi >= 25.0 {
					fmt.Println(b, a[z].Nama)
					b += 1
				}
			}
			break
		}
	}
}

func main() {
	var a tsiswa
	var i int
	i = 0
	for {
		fmt.Scan(&a[i].Nama, &a[i].Berat, &a[i].Tinggi)
		if a[i].Nama == "0" && a[i].Berat == 0 && a[i].Tinggi == 0 || i >= 100 {
			break
		}
		a[i].bmi = a[i].Berat / ((a[i].Tinggi / 100) * (a[i].Tinggi / 100))
		i += 1
	}
	cetak(a, i)
}
