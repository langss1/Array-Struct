package main

import (
	"fmt"
)

type Film struct {
	Judul  string
	Rating float64
}

type TabFilm [10]Film

func main() {
	var tab TabFilm
	var n int
	for i := 0; i < 10; i++ {
		fmt.Scan(&tab[i].Judul, &tab[i].Rating)
		if tab[i].Judul == "#" {
			n = i
			break
		}
	}

	for i := 0; i < n; i++ {
		if tab[i].Rating > 4.00 && tab[i].Rating <= 5.00 {
			fmt.Print(tab[i].Judul, " ")
		}
	}
	fmt.Println()
}
