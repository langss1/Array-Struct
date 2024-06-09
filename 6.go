
const nMax int = 100

type Belanja struct {
	Nama  string
	Harga int
}

type keranjang [nMax]Belanja

func inputBelanja(arrKeranjang *keranjang, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		var nama string
		var harga int
		fmt.Scan(&nama)
		fmt.Scan(&harga)

		barang := Belanja{Nama: nama, Harga: harga}
		arrKeranjang[i] = barang
	}
}

func tampilData(arrKeranjang *keranjang, n *int) {
	totalHarga := 0
	for i := 0; i < *n; i++ {
		fmt.Printf("%d. %s\n", i+1, arrKeranjang[i].Nama)
		totalHarga += arrKeranjang[i].Harga
	}
	fmt.Print("Total belanja :  ", totalHarga)
}