package main

import "fmt"

const NMax int = 1000

type Hasil struct {
	DapatHadiah bool
	gol         int
}

type thasil [NMax]Hasil

func check(a thasil, i int) {
	var x, jumlah int
	x = 0
	for x < i {
		jumlah += a[x].gol
		x += 1
	}
	if jumlah >= 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	x = 0
	for x < i {
		fmt.Println(a[x].gol)
		x += 1
	}

}

func main() {
	var a thasil
	var i int
	i = 0
	for {
		fmt.Scan(&a[i].gol)
		if a[i].gol == -1 || i >= NMax {
			break
		}
		i += 1
	}
	check(a, i)
}
