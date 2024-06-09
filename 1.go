package main

import "fmt"

//deklarasi tipe struct
type Episode struct {
	Nomor  int
	Durasi int
}

const NMax int = 1000

type episodes [NMax]Episode

func totalDurasiEpisodeGanjil(eps episodes, n int) int {
	// Mengembalikan total durasi jika nomor episode adalah ganjil.
	var i, total int
	total = 0
	for i = 0; i <= n; i++ {
		if eps[i].Nomor%2 != 0 {
			total += eps[i].Durasi
		}
	}
	return total
}

func main() {
	/* I.S.: Terdefinisi nilai bilangan bulat n. Data string sejumlah n buah tersedia
	   pada perangkat input. */
	/* F.S.: Array episodes diisi dengan data yang diberikan. */
	//deklarasi variabel array
	var eps episodes

	var n, i int

	fmt.Scan(&n)
	i = 0
	for i <= n {
		fmt.Scan(&eps[i].Nomor, &eps[i].Durasi)
		i += 1
	}
	fmt.Println(totalDurasiEpisodeGanjil(eps, n))
}
