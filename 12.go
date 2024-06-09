package main

import (
	"fmt"
)

type Orang struct {
	Nama string
	Usia int
	KTP  string
}

type TabOrang [100]Orang

func main() {
	var T TabOrang
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].Nama, &T[i].Usia, &T[i].KTP)
	}

	fmt.Println("Orang-orang yang berhak memilih pemilu:")
	for i := 0; i < n; i++ {
		if T[i].Usia >= 17 && T[i].KTP == "ada" {
			fmt.Println(T[i].Nama)
		}
	}
}
