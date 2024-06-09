package main

import "fmt"

type persegi struct {
	p, l, t, volume, luas int
}
type tpersegi [100]persegi

func hitungBalok(a tpersegi, i int) {
	/*  I.S data panjang, lebar, dan tinggi telah siap pada piranti masukan
	    proses: menerima masukan dari pengguna, masukan berhenti apabila
	    panjang, lebar, dan tinggi bernilai 0
	    F.S array b telah terisi sejumlah data dan menampilkan total luas permukaan semua balok, dan total volume semua balok*/
	var x, volume, luas int
	for x < i {
		a[x].volume = (a[x].p * a[x].l * a[x].t)
		a[x].luas = 2 * ((a[x].p * a[x].l) + (a[x].p * a[x].t) + (a[x].l * a[x].t))
		x += 1
	}
	x = 0
	for x < i {
		volume += a[x].volume
		luas += a[x].luas
		x += 1
	}

	fmt.Println(luas, volume)
}
func main() {
	//  lakukan pemanggilan procedure hitungBalok(...)
	var a tpersegi
	var i int
	for {
		fmt.Scan(&a[i].p, &a[i].l, &a[i].t)
		if a[i].p == 0 && a[i].l == 0 && a[i].t == 0 || i >= 100 {
			break
		}
		i++
	}
	hitungBalok(a, i)
}
