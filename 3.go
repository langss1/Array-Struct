package main

import "fmt"

type Parkir struct {
	tipe    string
	besaran int64
}

const NMax int = 1000

type parkir [NMax]Parkir

func main() {
	var a parkir
	var i, motor, mobil, n int64
	i = 1
	fmt.Scan(&n)
	for i <= n {
		fmt.Scan(&a[i].tipe, &a[i].besaran)
		i += 1
	}
	i = 1
	for i <= n {
		if a[i].tipe == "m" {
			motor += a[i].besaran
		} else if a[i].tipe == "b" {
			mobil += a[i].besaran
		}
		i += 1
	}
	fmt.Println(mobil, motor)
}
