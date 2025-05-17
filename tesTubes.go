package main

import "fmt"

const NMAX = 5

type proyek struct {
	nama      string
	teknologi string
	kategori  string
	kesulitan int
	tanggal   int
}

type daftarProyek struct {
	data      [NMAX]proyek
	tersimpan int
}

func tambahProyek(dp *daftarProyek, p proyek) {
	if dp.tersimpan < NMAX {
		dp.data[dp.tersimpan] = p
		dp.tersimpan++
	} else {
		fmt.Println("Kapasitas proyek sudah penuh!")
	}
	fmt.Println("Proyek berhasil ditambahkan!")
}

func tampilkanSemua(dp daftarProyek) {
	var i int
	if dp.tersimpan == 0 {
		fmt.Println("Belum ada data proyek.")
	} else {
		for i = 0; i < dp.tersimpan; i++ {
			fmt.Println("========== Proyek ke-",i+1, "==========")
			fmt.Println("Nama      :", dp.data[i].nama)
			fmt.Println("Teknologi :", dp.data[i].teknologi)
			fmt.Println("Kategori  :", dp.data[i].kategori)
			fmt.Println("Kesulitan :", dp.data[i].kesulitan)
			fmt.Println("Tanggal   :", dp.data[i].tanggal)
		}
	}
}

func cariProyek(dp daftarProyek, nama string) int {
	var i int
	for i = 0; i < dp.tersimpan; i++ {
		if dp.data[i].nama == nama {
			return i
		}
	}
	return -1
}

func urutkanKesulitan(dp *daftarProyek, n int) {
	var i, j, minIdx int
	var temp proyek

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if dp.data[j].kesulitan < dp.data[minIdx].kesulitan {
				minIdx = j
			}
		}
		temp = dp.data[i]
		dp.data[i] = dp.data[minIdx]
		dp.data[minIdx] = temp
	}
}

func urutkanTanggal(dp *daftarProyek, n int) {
	var i, j int
	var temp proyek

	for i = 1; i < n; i++ {
		temp = dp.data[i]
		j = i - 1
		for j >= 0 && dp.data[j].tanggal > temp.tanggal {
			dp.data[j+1] = dp.data[j]
			j = j - 1
		}
		dp.data[j+1] = temp
	}
}

func tampilkanStatistikKategori(dp daftarProyek) {
	
}

func ubahProyek(dp *daftarProyek, nama string) {
	var index int
	var p proyek
	index = cariProyek(*dp, nama)
	if index == -1 {
		fmt.Println("Proyek tidak ditemukan.")
	} else {
		fmt.Println("Masukkan data baru untuk proyek:")
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&p.nama)
		fmt.Print("Teknologi Digunakan: ")
		fmt.Scan(&p.teknologi)
		fmt.Print("Kategori Proyek: ")
		fmt.Scan(&p.kategori)
		fmt.Print("Tingkat Kesulitan (1-10): ")
		fmt.Scan(&p.kesulitan)
		fmt.Print("Tanggal Pembuatan (dd/mm/yyyy): ")
		fmt.Scan(&p.tanggal)
		dp.data[index] = p
		fmt.Println("Proyek berhasil diubah.")
	}
}

func hapusProyek(dp *daftarProyek, nama string) {
	

func main() {
	var daftar daftarProyek
	var pilihan, index int
	var p proyek
	var namaCari string
	var selesai bool

	selesai = false
	for !selesai {
		fmt.Println("\n=== Menu Aplikasi Manajemen Proyek ===")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tampilkan Semua Proyek")
		fmt.Println("3. Cari Proyek berdasarkan Nama")
		fmt.Println("4. Urutkan berdasarkan Kesulitan")
		fmt.Println("5. Urutkan berdasarkan Tanggal")
		fmt.Println("6. Tampilkan Statistik Kategori")
		fmt.Println("7. Ubah Proyek")
		fmt.Println("8. Hapus Proyek")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Nama Proyek: ")
			fmt.Scan(&p.nama)
			fmt.Print("Teknologi Digunakan: ")
			fmt.Scan(&p.teknologi)
			fmt.Print("Kategori Proyek: ")
			fmt.Scan(&p.kategori)
			fmt.Print("Tingkat Kesulitan (1-10): ")
			fmt.Scan(&p.kesulitan)
			fmt.Print("Tanggal Pembuatan (dd/mm/yyyy): ")
			fmt.Scan(&p.tanggal)
			tambahProyek(&daftar, p)

		} else if pilihan == 2 {
			tampilkanSemua(daftar)

		} else if pilihan == 3 {
			fmt.Print("Masukkan Nama Proyek: ")
			fmt.Scan(&namaCari)
			index = cariProyek(daftar, namaCari)
			if index != -1 {
				fmt.Println("Proyek ditemukan:")
				fmt.Println("Nama      :", daftar.data[index].nama)
				fmt.Println("Teknologi :", daftar.data[index].teknologi)
				fmt.Println("Kategori  :", daftar.data[index].kategori)
				fmt.Println("Kesulitan :", daftar.data[index].kesulitan)
				fmt.Println("Tanggal   :", daftar.data[index].tanggal)
			} else {
				fmt.Println("Proyek tidak ditemukan.")
			}

		} else if pilihan == 4 {
			urutkanKesulitan(&daftar, daftar.tersimpan)
			fmt.Println("Proyek berhasil diurutkan berdasarkan kesulitan.")

		} else if pilihan == 5 {
			urutkanTanggal(&daftar, daftar.tersimpan)
			fmt.Println("Proyek berhasil diurutkan berdasarkan tanggal.")

		} else if pilihan == 6 {
			tampilkanStatistikKategori(daftar)

		} else if pilihan == 7 {
			fmt.Print("Masukkan nama proyek yang ingin diubah: ")
			fmt.Scan(&namaCari)
			ubahProyek(&daftar, namaCari)

		} else if pilihan == 8 {
			fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
			fmt.Scan(&namaCari)
			hapusProyek(&daftar, namaCari)

		} else {
			if pilihan == 9 {
				fmt.Println("Terima kasih telah menggunakan aplikasi.")
			} else {
				if pilihan != 1 || pilihan != 2 || pilihan != 3 || pilihan != 4 || pilihan != 5 || pilihan != 6 || pilihan != 7 || pilihan != 8 || pilihan != 9 {
					fmt.Println("Pilihan tidak valid")
				}
			}
			selesai = true
		}
	}
}
