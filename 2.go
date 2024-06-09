package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah karakter adalah vokal
func isVowel(c rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if c == v {
			return true
		}
	}
	return false
}

func main() {
	var consonants, vowels []rune
	var countConsonants, countVowels int

	fmt.Println("Masukkan karakter satu per satu (akhiri dengan '.'): ")

	for {
		var input rune
		fmt.Scanf("%c\n", &input)

		if input == '.' {
			break
		}

		// Mengecek apakah karakter adalah vokal atau konsonan
		if isVowel(input) {
			vowels = append(vowels, input)
			countVowels++
		} else if (input >= 'a' && input <= 'z') || (input >= 'A' && input <= 'Z') {
			consonants = append(consonants, input)
			countConsonants++
		}
	}

	// Mencetak hasil analisis
	fmt.Printf("Deret konsonan: %s\n", string(consonants))
	fmt.Printf("Jumlah konsonan: %d\n", countConsonants)
	fmt.Printf("Deret vokal: %s\n", string(vowels))
	fmt.Printf("Jumlah vokal: %d\n", countVowels)
}
