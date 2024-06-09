package main

import "fmt"

type kubus struct {
	sisi int
}
type arrKubus [100]kubus

func hitungKubus(a *arrKubus, n int) {
	/*  I.S data sejumlah bilangan yang menyatakan sisi telah siap pada piranti masukan
	    F.S array k telah terisi sejumlah data dan menampilkan total luas permukaan semua kubus, serta total volume semua kubus */
	var i, totV, totP int
	i = 0
	totV = 0
	totP = 0
	for i < n {
		totV += (a[i].sisi * a[i].sisi * a[i].sisi)
		totP += (6 * (a[i].sisi * a[i].sisi))
		i++
	}
	fmt.Println(totP, totV)
}
func main() {
	var a arrKubus
	//  lakukan pemanggilan prosedur hitungKubus(...)
	var i int
	for {
		fmt.Scan(&a[i].sisi)
		if a[i].sisi == 0 {
			break
		}
		i += 1
	}

	hitungKubus(&a, i)
}
