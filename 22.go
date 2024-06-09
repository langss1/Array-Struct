type lagu struct {
	Judul  string
	Artis  string
	Durasi int
}

// Tipe data bentukan Playlist
type playlist struct {
	Lagu        [20]lagu
	TotalDurasi int
	JumlahLagu  int
}

// Fungsi untuk menambahkan lagu ke dalam playlist
func insertLagu(pl *playlist, n int) {
	for i := 0; i < n; i++ {
		var judul, artis string
		var durasi int
		fmt.Scan(&judul, &artis, &durasi)
		pl.Lagu[pl.JumlahLagu] = lagu{Judul: judul, Artis: artis, Durasi: durasi}
		pl.TotalDurasi += durasi
		pl.JumlahLagu++
	}
}

// Fungsi untuk menggabungkan dua playlist tanpa duplikasi
func gabungPlaylist(pl1, pl2 playlist, pl3 *playlist) {
	for i := 0; i < pl1.JumlahLagu; i++ {
		addToPlaylist(pl3, pl1.Lagu[i])
	}
	for i := 0; i < pl2.JumlahLagu; i++ {
		addToPlaylist(pl3, pl2.Lagu[i])
	}
}

// Fungsi untuk menambahkan lagu ke dalam playlist hasil tanpa duplikasi
func addToPlaylist(pl *playlist, lagu lagu) {
	for i := 0; i < pl.JumlahLagu; i++ {
		if pl.Lagu[i].Judul == lagu.Judul && pl.Lagu[i].Artis == lagu.Artis {
			return
		}
	}
	pl.Lagu[pl.JumlahLagu] = lagu
	pl.TotalDurasi += lagu.Durasi
	pl.JumlahLagu++
}

// Fungsi untuk menampilkan playlist
func tampilPlaylist(pl playlist) {
	for i := 0; i < pl.JumlahLagu; i++ {
		fmt.Println(pl.Lagu[i].Judul, pl.Lagu[i].Artis)
	}
	durasiMenit := (pl.TotalDurasi % 3600) / 60
	durasiDetik := pl.TotalDurasi % 60
	fmt.Printf("Durasi Playlist: %d:%d\n", durasiMenit, durasiDetik)
}