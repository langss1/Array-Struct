package main

import (
	"fmt"
	"strings"
)

type Kalimat struct {
	Teks string
}

func main() {
	var kalimat [1000]Kalimat
	var n int
	for i := 0; i < 1000; i++ {
		fmt.Scan(&kalimat[i].Teks)
		if strings.ToLower(kalimat[i].Teks) == "stop" {
			n = i
			break
		}
	}

	totalUppercase := 0
	totalLowercase := 0
	totalAngka := 0

	for i := 0; i < n; i++ {
		for _, char := range kalimat[i].Teks {
			if char >= 'A' && char <= 'Z' {
				totalUppercase++
			} else if char >= 'a' && char <= 'z' {
				totalLowercase++
			} else if char >= '0' && char <= '9' {
				totalAngka++
			}
		}
	}

	fmt.Print("Total Uppercase :  ", totalUppercase, "\n")
	fmt.Print("Total Lowercase :  ", totalLowercase, "\n")
	fmt.Print("Total Angka :  ", totalAngka, "\n")
}
