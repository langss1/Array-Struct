type nilaiMahasiswa struct {
	nama  string
	quiz  float64
	tugas float64
	ujian float64
}
type tabNilai [1000]nilaiMahasiswa

func inputArray(a *tabNilai, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat n. Data nilai mahasiswa tersedia
	  pada input device sebanyak n*/
	/*F.S. Array T berisi data yang diberikan*/
	/*Petunjuk : Lakukan perulangan sebanyak n kali untuk melakukan input terhadap
	  data nilai mahasiswa*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&a[i].nama, &a[i].quiz, &a[i].tugas, &a[i].ujian)
	}
}

func nilaiAkhir(quiz, tugas, ujian float64) float64 {
	/*Mengembalikan nilai akhir mahasiswa sesuai dengan bobot pada soal
	  Petunjuk: Nilai akhir =  20%*quiz + 10%*tugas +70%*ujian */
	return (0.2 * quiz) + (0.1 * tugas) + (0.7 * ujian)
}

func cetakNilaiAkhir(a tabNilai, n int) {
	/*I.S. Terdefinisni array T yang berisikan sebanyak n data nilai mahasiswa
	  (nama mahasiswa, nilai quiz, nilai tugas, dan nilai ujian)
	  F.S. Menampilkan ke layar daftar nama mahasiswa beserta dengan nilai akhirnya
	  Petunjuk: Untuk mendapatkan nilai akhir, manfaatkan fungsi nilaiAkhir*/
	var i int
	i = 0
	for i < n {
		fmt.Println(a[i].nama, nilaiAkhir(a[i].quiz, a[i].tugas, a[i].ujian))
		i += 1
	}
}