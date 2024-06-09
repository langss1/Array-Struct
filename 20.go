type mahasiswa struct {
	nama  string
	nilai float64
}
type siswa [100]mahasiswa
type remidial [100]mahasiswa

func hitungRata(a *siswa, r *remidial) {
	/*  I.S data sejumlah nama dan niai siswa telah siap pada piranti masukan
	    proses: simpan data mahasiswa yang remedial pada array remidial dan hitung banyaknya mahasiswa yang remedial
	    F.S array siswa dan remidial telah terisi sejumlah data dan menampilkan nilai rata-rata, banyaknya siswa yang remedial, serta nama-nama mahasiswa yang remedial*/
	var i, x int
	var rata float64
	for {
		fmt.Scan(&a[i].nama, &a[i].nilai)
		if a[i].nama == "0" && a[i].nilai == 0 || i >= 100 {
			break
		}
		rata += a[i].nilai
		if a[i].nilai < 50 {
			r[x].nama = a[i].nama
			x += 1
		}
		i += 1
	}

	rata = rata / float64(i)
	fmt.Print("Rata-rata : ")
	fmt.Printf("%.2f\n", rata)
	fmt.Print("Remidi : ", x, " Siswa")
	fmt.Println()
	for z := 0; z < x; z++ {
		fmt.Print(z+1, " . ", r[z].nama)
		fmt.Println()
	}
}