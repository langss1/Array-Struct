type Kerucut struct {
	R, S, T, Volume, LuasPermukaan float64
}

// Tipe data bentukan untuk menyimpan array dari kerucut
type tabK struct {
	Kerucuts [10]Kerucut
}

// Fungsi untuk menghitung volume dan luas permukaan kerucut
func volumeDanLuasP(tab *tabK, n int) {
	for i := 0; i < n; i++ {
		r := tab.Kerucuts[i].R
		s := tab.Kerucuts[i].S
		t := tab.Kerucuts[i].T

		// Menghitung volume
		tab.Kerucuts[i].Volume = (1.0 / 3.0) * 3.14 * r * r * t

		// Menghitung luas permukaan
		tab.Kerucuts[i].LuasPermukaan = 3.14 * r * (r + s)
	}
}

// Fungsi untuk mengisi data kerucut
func input(tab *tabK, n int) {
	for i := 0; i < n; i++ {
		var r, s, t float64
		fmt.Scan(&r, &s, &t)
		tab.Kerucuts[i] = Kerucut{R: r, S: s, T: t}
	}
}

// Fungsi untuk menampilkan hasil perhitungan
func tampilkan(tab tabK, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Volume Kerucut %d = %.2f\n", i+1, tab.Kerucuts[i].Volume)
		fmt.Printf("Luas Permukaan Kerucut %d = %.2f\n\n", i+1, tab.Kerucuts[i].LuasPermukaan)
	}
}
