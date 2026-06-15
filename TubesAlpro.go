package main
import "fmt"

const NMAX = 100

type riwayatservice struct {
	tanggal        string
	jeniskerusakan string
	detailservice  string
}

type pemilik struct {
	nama   string
	notelp string
}

type kendaraan struct {
	platNomor     string
	tahunProduksi int
	dataPemilik   pemilik
	riwayat       [NMAX]riwayatservice
	jumlahriwayat int
}

type tabkendaraan [NMAX]kendaraan

func tambahkendaraan(datakendaraan *tabkendaraan,jumlahkendaraan *int) {
	var n,i int
	
	if *jumlahkendaraan >= NMAX {
		fmt.Println("Jumlah kendaraan sudah mencapai batas maksimum")
		return
	}
	fmt.Print("Masukkan jumlah kendaraan yang ingin ditambahkan: ")
	fmt.Scan(&n)
	if *jumlahkendaraan + n > NMAX {
		fmt.Print("Jumlah kendaraan yang ingin ditambah mencapai batas maksimum,anda hanya bisa menambahkan %d kendaraan\n",NMAX - *jumlahkendaraan)
	}
	for i = *jumlahkendaraan;i < *jumlahkendaraan + n;i++{
		fmt.Printf("Masukkan data kendaraan ke %d",i+1)
		fmt.Print("Masukkan plat nomor: ")
		fmt.Scan(&datakendaraan[i].platNomor)
		fmt.Print("Masukkan tahun produksi :")
		fmt.Scan(&datakendaraan[i].tahunProduksi)
		fmt.Print("Masukkan nama pemilik: ")
		fmt.Scan(&datakendaraan[i].dataPemilik.nama)
		fmt.Print("Masukkan nomor telpon pemilik: ")
		fmt.Scan(&datakendaraan[i].dataPemilik.notelp)
		datakendaraan[i].jumlahriwayat = 0
	}
	*jumlahkendaraan = *jumlahkendaraan + n
}

func hapuskendaraan(datakendaraan *tabkendaraan,jumlahkendaraan *int){
	var platNomor string
	var i,j int
	
	fmt.Print("Masukkan platNomor yang ingin dihapus: ")
	fmt.Scan(&platNomor)
	for i = 0;i < *jumlahkendaraan;i++ {
		if datakendaraan[i].platNomor == platNomor {
			for j = 0;j < *jumlahkendaraan;j++ {
			datakendaraan[j] = datakendaraan[j+1]
			}
			*jumlahkendaraan--
			fmt.Println("Kendaraan berhasil dihapus")
			return
		}
	}
	fmt.Println("kendaraan tidak ditemukan")
}

func ubahdatakendaraan(datakendaraan *tabkendaraan,jumlahkendaraan *int) {
	var platNomor string
	var i int
	fmt.Print("Masukkan platnomor kendaaraan yang ingin diubah: ")
	fmt.Scan(&platNomor)
	for i = 0;i < *jumlahkendaraan;i++{
		if datakendaraan[i].platNomor == platNomor {
			fmt.Print("Tahun produksi baru: ")
			fmt.Scan(&datakendaraan[i].tahunProduksi)
			fmt.Println("Data berhasil diubah")
			return 
		}
	}
	fmt.Print("Data tidak ditemukan")
}

func ubahdatapemilik(datakendaraan *tabkendaraan,jumlahkendaraan *int) {
	var platNomor string
	var i int
	
	fmt.Print("Masukkan platNomor pemilik yang ingin diubah: ")
	fmt.Scan(&platNomor)
	for i = 0;i < *jumlahkendaraan;i++ {
		if datakendaraan[i].platNomor == platNomor {
			fmt.Print("Masukkan nama pemilik baru: ")
			fmt.Scan(&datakendaraan[i].dataPemilik.nama)
			fmt.Println("Data berhasil diubah")
		}
	}
	fmt.Print("Data tidak ditemukan")
}

func tambahriwayatservice(datakendaraan *tabkendaraan,jumlahkendaraan *int) {
	var platNomor string
	var i int
	fmt.Print("Masukkan platNomor yang ingin ditambah riwayat service: ")
	fmt.Scan(&platNomor)
	for i = 0;i < *jumlahkendaraan;i++{
		if datakendaraan[i].platNomor == platNomor {
			if datakendaraan[i].jumlahriwayat < NMAX {
			fmt.Print("tanggal service: ")
			fmt.Scan(&datakendaraan[i].riwayat[datakendaraan[i].jumlahriwayat].tanggal)
			fmt.Println("jenis kerusakan: ")
			fmt.Scan(&datakendaraan[i].riwayat[datakendaraan[i].jumlahriwayat].jeniskerusakan)
			fmt.Println("Detai service: ")
			fmt.Scan(&datakendaraan[i].riwayat[datakendaraan[i].jumlahriwayat].detailservice)
			datakendaraan[i].jumlahriwayat++
			fmt.Println("Riwayat service berhasil ditambahkan")
			}else {
			fmt.Print("Jumlah riwayat service sudah penuh")
			}
			return
		}
		fmt.Print("Kendaraan tidak ditemukan")	
	}
}

func tampilkanriwayatservice(datakendaraan tabkendaraan,jumlahkendaraan int) {
	var platNomor string
	var i,j int
	fmt.Print("Masukkan platNomor kendaraan yang ingin di tampilkan riwayat service: ")
	fmt.Scan(&platNomor)
	for i = 0;i < jumlahkendaraan;i++ {
		if datakendaraan[i].platNomor == platNomor {
			fmt.Println("Riwayat service: ")
		for j = 0;j < datakendaraan[i].jumlahriwayat;j++ {
				fmt.Printf("Tanggal: %s\n",datakendaraan[i].riwayat[j].tanggal)
				fmt.Printf("Riwayat kerusakan %s\n",datakendaraan[i].riwayat[j].jeniskerusakan)
				fmt.Printf("Riwayat detail kerusakan %s\n",datakendaraan[i].riwayat[j].detailservice)
			}
			return
		}
	}
	fmt.Println("Kendaraan tidak ditemukan")
}

func tampilkansemuariwayatservice(datakendaraan tabkendaraan,jumlahkendaraan int){
	var i,j int
	fmt.Print("Semua riwayat service")
	for i = 0;i < jumlahkendaraan; i++ {
		fmt.Printf("Kendaraan dengan plat nomor: %s\n",datakendaraan[i].platNomor)
		if datakendaraan[i].jumlahriwayat == 0 {
			fmt.Println("Tidak ada riwayat service")
		} 
		for j = 0; j < datakendaraan[j].jumlahriwayat; j++ {
			fmt.Printf("Tanggal: %s\n",datakendaraan[i].riwayat[j].tanggal)
			fmt.Printf("Jenis kerusakan: %s\n",datakendaraan[i].riwayat[j].jeniskerusakan)
			fmt.Printf("Detail kerusakan: %s\n",datakendaraan[i].riwayat[j].detailservice)
		}
		
	}
}

func tampilkansemuakendaraan(datakendaraan tabkendaraan,jumlahkendaraan int){
	var i int
	
	for i = 0;i < jumlahkendaraan; i++{
		fmt.Printf("%d. Kendaraan dengan plat nomor: %s\n",i+1,datakendaraan[i].platNomor)
		fmt.Printf("Tahun produksi: %d\n",datakendaraan[i].tahunProduksi)
		fmt.Printf("Nomor telpon pemilik: %s\n",datakendaraan[i].dataPemilik.notelp)
	}
}

func binarysearch(datakendaraan tabkendaraan,jumlahkendaraan int,platNomor string) int {
	var kiri,kanan,mid int
	kiri = 0
	kanan = jumlahkendaraan - 1
	for kiri <= kanan {
		mid = (kiri + kanan) / 2
		if datakendaraan[mid].platNomor == platNomor {
			return mid
		}else if datakendaraan[mid].platNomor < platNomor {
			kiri = mid + 1
		}else if datakendaraan[mid].platNomor > platNomor {
			kanan = mid - 1
		}
	}
	return -1
}

func sequentialsearch(datakendaraan tabkendaraan,jumlahkendaraan int,platNomor string) int {
	for i := 0;i < jumlahkendaraan;i++ {
		if datakendaraan[i].platNomor == platNomor {
			return i
		} 
	}
	return -1
}

func carikendaraan(datakendaraan tabkendaraan,jumlahkendaraan int) {
	var platNomor string
	var pilihan,index int
	
	if jumlahkendaraan < 1 {
		fmt.Println("Tidak ada kendaraan yang dapat dicari")
		return 
	}
	fmt.Print("Pilih metode pencarian 1 untuk sequentialsearch dan 2 untuk binarysearch: ")
	fmt.Scan(&pilihan)
	fmt.Print("Masukkan plat nomor kendaaraan yang ingin dicari: ")
	fmt.Scan(&platNomor)
	if pilihan == 1 {
		index = sequentialsearch(datakendaraan, jumlahkendaraan, platNomor)
			if index != -1 {
			fmt.Printf("Kendaraan Plat Nomor: %s\n", datakendaraan[index].platNomor)
			fmt.Printf("  Tahun Produksi: %d\n", datakendaraan[index].tahunProduksi)
			fmt.Printf("  Nama Pemilik: %s\n", datakendaraan[index].dataPemilik.nama)
			fmt.Printf("  Nomor Telepon: %s\n", datakendaraan[index].dataPemilik.notelp)
			}else{
				fmt.Println("Kendaraan tidak ditemukan")
			}
    
	} else if pilihan == 2 {
		index = binarysearch(datakendaraan, jumlahkendaraan, platNomor)
			if index != -1 {
			fmt.Printf("Kendaraan Plat Nomor: %s\n", datakendaraan[index].platNomor)
			fmt.Printf("  Tahun Produksi: %d\n", datakendaraan[index].tahunProduksi)
			fmt.Printf("  Nama Pemilik: %s\n", datakendaraan[index].dataPemilik.nama)
			fmt.Printf("  Nomor Telepon: %s\n", datakendaraan[index].dataPemilik.notelp)
			}else{
				fmt.Println("Kendaraan tidak ditemukan")
			}
	}else{
		fmt.Println("Pilihan metode pencarian tidak valid")
	}
}

func selectionsort(datakendaraan *tabkendaraan, jumlahkendaraan *int) {
	for i := 0; i < *jumlahkendaraan-1; i++ {
		minIndex := i
		for j := i + 1; j < *jumlahkendaraan; j++ {
			if datakendaraan[j].tahunProduksi < datakendaraan[minIndex].tahunProduksi {
				minIndex = j
			}
		}
		datakendaraan[i], datakendaraan[minIndex] = datakendaraan[minIndex], datakendaraan[i]
	}
	fmt.Println("[Selection Sort] Kendaraan diurutkan berdasarkan Tahun Produksi :")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("   %-12s %-5s\n", "Plat", "Tahun")
	fmt.Println("-------------------------------------------------------------")
	for i := 0; i < *jumlahkendaraan; i++ {
		fmt.Printf("%d. %-12s %d\n", i+1, datakendaraan[i].platNomor, datakendaraan[i].tahunProduksi)
	}
}

func insertionsort(datakendaraan *tabkendaraan, jumlahkendaraan *int) {
	for i := 1; i < *jumlahkendaraan; i++ {
		key := datakendaraan[i]
		j := i - 1
		for j >= 0 && datakendaraan[j].tahunProduksi < key.tahunProduksi {
			datakendaraan[j+1] = datakendaraan[j]
			j--
		}
		datakendaraan[j+1] = key
	}

	fmt.Println("[Insertion Sort] Kendaraan diurutkan berdasarkan Tahun Produksi :")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("   %-12s %-5s\n", "Plat", "Tahun")
	fmt.Println("-------------------------------------------------------------")
	for i := 0; i < *jumlahkendaraan; i++ {
		fmt.Printf("%d. %-12s %d\n", i+1, datakendaraan[i].platNomor, datakendaraan[i].tahunProduksi)
	}
	fmt.Println("-------------------------------------------------------------")
}

func urutkankendaraan(datakendaraan *tabkendaraan, jumlahkendaraan *int) {
	if *jumlahkendaraan < 1 {
		fmt.Println("Tidak ada kendaraan yang tersedia untuk diurutkan.")
		return
	}
	var pilihan int
	fmt.Print("Pilih metode pengurutan (1 untuk Selection Sort, 2 untuk Insertion Sort): ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		selectionsort(datakendaraan, jumlahkendaraan)
		fmt.Println("Data kendaraan berhasil diurutkan menggunakan Selection Sort.")
	} else if pilihan == 2 {
		insertionsort(datakendaraan, jumlahkendaraan)
		fmt.Println("Data kendaraan berhasil diurutkan menggunakan Insertion Sort.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func menu() {
	cetakGaris()
	fmt.Println("  Aplikasi Manajemen dan Riwayat Servis Kendaraan")
	cetakGarisTipis()
	fmt.Println(" 1. Tambah Kendaraan")
	fmt.Println(" 2. Hapus Kendaraan")
	fmt.Println(" 3. Ubah Data Kendaraan")
	fmt.Println(" 4. Tambah Riwayat Service")
	fmt.Println(" 5. Tampilkan Riwayat Service")
	fmt.Println(" 6. Tampilkan Semua Riwayat Service")
	fmt.Println(" 7. Tampilkan Semua Data Kendaraan")
	fmt.Println(" 8. Cari Kendaraan")
	fmt.Println(" 9. Urutkan Kendaraan")
	fmt.Println(" 10. Statistik")
	fmt.Println(" 0. Keluar")
	cetakGarisTipis()
}

func prosesmenu(datakendaraan *tabkendaraan, jumlahkendaraan *int) {
	var pilihan int
	for {
		menu()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahkendaraan(datakendaraan, jumlahkendaraan)
		case 2:
			hapuskendaraan(datakendaraan, jumlahkendaraan)
		case 3:
			ubahdatakendaraan(datakendaraan, jumlahkendaraan)
		case 4:
			tambahriwayatservice(datakendaraan, jumlahkendaraan)
		case 5:
			tampilkanriwayatservice(*datakendaraan, *jumlahkendaraan)
		case 6:
			tampilkansemuariwayatservice(*datakendaraan, *jumlahkendaraan)
		case 7:
			tampilkansemuakendaraan(*datakendaraan, *jumlahkendaraan)
		case 8:
			carikendaraan(*datakendaraan, *jumlahkendaraan)
		case 9:
			urutkankendaraan(datakendaraan, jumlahkendaraan)
		case 10:
			statistik(datakendaraan, jumlahkendaraan)
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func statistik(datakendaraan *tabkendaraan, jumlahkendaraan *int) {
	// Fungsi untuk menampilkan statistik kendaraan yang diservis per bulan dan kategori kerusakan paling sering
	if *jumlahkendaraan == 0 {
		fmt.Println("Tidak ada data kendaraan.")
		return
	}

	fmt.Println("\nJumlah Kendaraan Diservis per Bulan:")
	fmt.Println("-------------------------------------------------")
	namaBulan := [13]string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	var jumlahPerBulan [13]int

	for i := 0; i < *jumlahkendaraan; i++ {
		k := datakendaraan[i]
		for j := 0; j < k.jumlahriwayat; j++ {
			tanggal := k.riwayat[j].tanggal
			if len(tanggal) >= 7 {
				bulan := int(tanggal[5]-'0')*10 + int(tanggal[6]-'0')
				if bulan >= 1 && bulan <= 12 {
					jumlahPerBulan[bulan]++
				}
			}
		}
	}

	ada := false
	for b := 1; b <= 12; b++ {
		if jumlahPerBulan[b] > 0 {
			bar := ""
			for k := 0; k < jumlahPerBulan[b]; k++ {
				bar = bar + "*"
			}
			fmt.Printf(" %-10s : %s (%d)\n", namaBulan[b], bar, jumlahPerBulan[b])
			ada = true
		}
	}
	if !ada {
		fmt.Println("  (Belum ada data servis)")
	}

	fmt.Println("\nKategori Kerusakan Paling Sering:")
	fmt.Println("-------------------------------------------------")
	var jenisUnik [NMAX]string
	var hitungJenis [NMAX]int
	jumlahJenis := 0

	for i := 0; i < *jumlahkendaraan; i++ {
		k := datakendaraan[i]
		for j := 0; j < k.jumlahriwayat; j++ {
			jenis := k.riwayat[j].jeniskerusakan
			if jenis == "" {
				// tidak ada jenis kerusakan pada riwayat ini
			} else {
				ketemu := false
				for x := 0; x < jumlahJenis && !ketemu; x++ {
					if jenisUnik[x] == jenis {
						hitungJenis[x]++
						ketemu = true
					}
				}

				if !ketemu && jumlahJenis < NMAX {
					jenisUnik[jumlahJenis] = jenis
					hitungJenis[jumlahJenis] = 1
					jumlahJenis++
				}
			}
		}
	}

	if jumlahJenis == 0 {
		fmt.Println("  (Belum ada data kerusakan)")
	} else {
		for i := 0; i < jumlahJenis-1; i++ {
			idxMax := i
			for j := i + 1; j < jumlahJenis; j++ {
				if hitungJenis[j] > hitungJenis[idxMax] {
					idxMax = j
				}
			}
			jenisUnik[i], jenisUnik[idxMax] = jenisUnik[idxMax], jenisUnik[i]
			hitungJenis[i], hitungJenis[idxMax] = hitungJenis[idxMax], hitungJenis[i]
		}

		for i := 0; i < jumlahJenis; i++ {
			bar := ""
			for k := 0; k < hitungJenis[i]; k++ {
				bar = bar + "*"
			}
			fmt.Printf(" %-15s : %s (%d)\n", jenisUnik[i], bar, hitungJenis[i])
		}
	}
}

func cetakGaris() {
	fmt.Println("==================================================")
}

func cetakGarisTipis() {
	fmt.Println("--------------------------------------------------")
}


func main() {
	var datakendaraan tabkendaraan
	var jumlahkendaraan int
	fmt.Println("Selamat datang di AutoCare")
	fmt.Println("Silakan pilih menu untuk mengelola data kendaraan dan riwayat service.")
	prosesmenu(&datakendaraan, &jumlahkendaraan)
}