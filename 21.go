package main

import "fmt"

const MAX_SONGS = 20

// Definisi tipe data bentukan untuk Lagu
type Lagu struct {
	judul  string
	artis  string
	durasi int
}

// Definisi tipe data bentukan untuk Playlist
type Playlist struct {
	lagu        [MAX_SONGS]Lagu
	totalDurasi int
	jumlahLagu  int
}

// Fungsi untuk memasukkan lagu ke dalam playlist
func insertLagu(pl *Playlist, n int) {
	for i := 0; i < n; i++ {
		var judul, artis string
		var durasi int
		fmt.Scan(&judul, &artis, &durasi)
		pl.lagu[pl.jumlahLagu] = Lagu{judul, artis, durasi}
		pl.totalDurasi += durasi
		pl.jumlahLagu++
	}
}

// Fungsi untuk menggabungkan dua playlist dan menghapus duplikat
func playlistFav(pl1, pl2 Playlist, pl3 *Playlist) {
	for i := 0; i < pl1.jumlahLagu; i++ {
		pl3.lagu[pl3.jumlahLagu] = pl1.lagu[i]
		pl3.totalDurasi += pl1.lagu[i].durasi
		pl3.jumlahLagu++
	}

	for i := 0; i < pl2.jumlahLagu; i++ {
		duplicate := false
		for j := 0; j < pl1.jumlahLagu; j++ {
			if pl2.lagu[i].judul == pl1.lagu[j].judul && pl2.lagu[i].artis == pl1.lagu[j].artis {
				duplicate = true
				break
			}
		}
		if !duplicate {
			pl3.lagu[pl3.jumlahLagu] = pl2.lagu[i]
			pl3.totalDurasi += pl2.lagu[i].durasi
			pl3.jumlahLagu++
		}
	}
}

// Fungsi untuk menampilkan playlist
func tampilPlaylist(pl Playlist) {
	for i := 0; i < pl.jumlahLagu; i++ {
		fmt.Printf("%s %s\n", pl.lagu[i].judul, pl.lagu[i].artis)
	}
	fmt.Printf("Durasi Playlist: %d:%d\n", pl.totalDurasi/60, pl.totalDurasi%60)
}

func main() {
	var pl1, pl2, pl3 Playlist
	var nZaki, nMifta int

	fmt.Scanln(&nZaki)
	insertLagu(&pl1, nZaki)

	fmt.Scanln(&nMifta)
	insertLagu(&pl2, nMifta)

	playlistFav(pl1, pl2, &pl3)
	tampilPlaylist(pl3)
}