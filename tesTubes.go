package main

import "fmt"

const NMAX = 100000

type proyek struct {
	nama, teknologi, kategori string
	kesulitan, tanggal int
}

type daftarProyek struct {
	data [NMAX]proyek
	tersimpan int
}

func tambahProyek(dp *daftarProyek, p proyek) {
	if dp.tersimpan < NMAX {
		dp.data[dp.tersimpan] = p
		dp.tersimpan++
		fmt.Println("Proyek berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas proyek sudah penuh!")
	}
}

func tampilkanSemuaProyek(dp daftarProyek) {
	var i int
	if dp.tersimpan == 0 {
		fmt.Println("Belum ada data proyek.")
	} else {
		for i = 0; i < dp.tersimpan; i++ {
			fmt.Printf("|=========== Proyek ke-%d ==========|\n", i+1)
			fmt.Printf("|%-10s : %-20s |\n","Nama", dp.data[i].nama)
			fmt.Printf("|%-10s : %-20s |\n","Teknologi", dp.data[i].teknologi)
			fmt.Printf("|%-10s : %-20s |\n","Kategori", dp.data[i].kategori)
			fmt.Printf("|%-10s : %-20d |\n","Kesulitan", dp.data[i].kesulitan)
			fmt.Printf("|%-10s : %-20d |\n","Tanggal", dp.data[i].tanggal)
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

func urutKesulitan(dp *daftarProyek, n int) {
	var i, idx, pass int
	var temp proyek

	n = dp.tersimpan
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if dp.data[i].kesulitan < dp.data[idx].kesulitan {
				idx = i
			}
			i++
		}
		temp = dp.data[pass-1]
		dp.data[pass-1] = dp.data[idx]
		dp.data[idx] = temp
		pass++
	}
}


func urutTanggal(dp *daftarProyek, n int) {
	var i, pass int
	var temp proyek

	n = dp.tersimpan
	pass = 1
	for pass <= n - 1 {
		i = pass
		temp = dp.data[pass]
		for i > 0 && temp.tanggal < dp.data[i-1].tanggal {
			dp.data[i] = dp.data[i-1]
			i--
		}
		dp.data[i] = temp
		pass++
	}
}

func jumProyekKatagori(dp daftarProyek) {
	
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
		fmt.Print("Tanggal Pembuatan (1-31): ")
		fmt.Scan(&p.tanggal)
		dp.data[index] = p
		fmt.Println("Proyek berhasil diubah")
	}
}

func hapusProyek(dp *daftarProyek, nama string) {
	
}

func main() {
	var daftar daftarProyek
	var pilihan, index, nData int
	var p proyek
	var cariNama string
	var selesai bool

	selesai = false
	for !selesai {
		fmt.Println()
		fmt.Println("|=========||| SELAMAT DATANG |||=========|")
		fmt.Println("|                                        |")
		fmt.Println("|-----Menu Aplikasi Manajemen Proyek-----|")
		fmt.Println("|1. Tambah Proyek                        |")
		fmt.Println("|2. Tampilkan Semua Proyek               |")
		fmt.Println("|3. Cari Proyek berdasarkan Nama         |")
		fmt.Println("|4. Urutkan berdasarkan Kesulitan        |")
		fmt.Println("|5. Urutkan berdasarkan Tanggal          |")
		fmt.Println("|6. Tampilkan Statistik Kategori         |")
		fmt.Println("|7. Ubah Proyek                          |")
		fmt.Println("|8. Hapus Proyek                         |")
		fmt.Println("|9. Keluar                               |")
		fmt.Println("|========================================|")
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
			if p.kesulitan < 1 || p.kesulitan > 10 {
				fmt.Println("Input tidak valid, kesulitan harus 1-10!")
				selesai = true
			} else {
				fmt.Print("Tanggal Pembuatan (1-31): ")
				fmt.Scan(&p.tanggal)
				if p.tanggal < 1 || p.tanggal > 31 {
					fmt.Println("Input tidak valid, tanggal harus 1-31")
					selesai = true
				} else {
					tambahProyek(&daftar, p)
				}
			}
		} else if pilihan == 2 {
			tampilkanSemua(daftar)
		} else if pilihan == 3 {
			fmt.Print("Masukkan Nama Proyek: ")
			fmt.Scan(&cariNama)
			index = cariProyek(daftar, cariNama)
			if index != -1 {
				fmt.Println("Proyek ditemukan:")
				fmt.Println("Nama :", daftar.data[index].nama)
				fmt.Println("Teknologi :", daftar.data[index].teknologi)
				fmt.Println("Kategori :", daftar.data[index].kategori)
				fmt.Println("Kesulitan :", daftar.data[index].kesulitan)
				fmt.Println("Tanggal :", daftar.data[index].tanggal)
			} else {
				fmt.Println("Proyek tidak ditemukan")
			}
		} else if pilihan == 4 {
			urutkanKesulitan(&daftar, nData)
			fmt.Println("Proyek berhasil diurutkan berdasarkan kesulitan")
		} else if pilihan == 5 {
			urutkanTanggal(&daftar, nData)
			fmt.Println("Proyek berhasil diurutkan berdasarkan tanggal")
		} else if pilihan == 6 {
			jumProyekKatagori(daftar)
		} else if pilihan == 7 {
			fmt.Print("Masukkan nama proyek yang ingin diubah: ")
			fmt.Scan(&cariNama)
			ubahProyek(&daftar, cariNama)
		} else if pilihan == 8 {
			fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
			fmt.Scan(&cariNama)
			hapusProyek(&daftar, cariNama)
		} else {
			if pilihan == 9 {
				fmt.Println("Terima kasih telah menggunakan aplikasi")
			} else {
				if pilihan != 1 || pilihan != 2 || pilihan != 3 || pilihan != 4 || pilihan != 5 || pilihan != 6 || pilihan != 7 || pilihan != 8 || pilihan != 9 {
					fmt.Println("Pilihan tidak valid")
				}
			}
			selesai = true
		}
	}
}
