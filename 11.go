package main

import "fmt"

type mahasiswa struct {
	nama                   string
	tugas, uts, uas, nilai int
}
type tmahasiswa [10]mahasiswa

func grade(a tmahasiswa, ab, b, c, d, e *int) {
	for i := 0; i < 10; i++ {
		if (a[i].nilai) >= 80 {
			*ab += 1
		} else if (a[i].nilai) >= 70 {
			*b += 1
		} else if (a[i].nilai) >= 59 {
			*c += 1
		} else if (a[i].nilai) >= 50 {
			*d += 1
		} else {
			*e += 1
		}
	}
}

func main() {
	var a tmahasiswa
	var ab, b, c, d, e int
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i].nama, &a[i].tugas, &a[i].uts, &a[i].uas)
		a[i].nilai = (a[i].tugas * 30 / 100) + (a[i].uts * 30 / 100) + (a[i].uas * 40 / 100)
	}
	grade(a, &ab, &b, &c, &d, &e)
	fmt.Println("Grade A : ", ab)
	fmt.Println("Grade B : ", b)
	fmt.Println("Grade C : ", c)
	fmt.Println("Grade D : ", d)
	fmt.Println("Grade E : ", e)
}
