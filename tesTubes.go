package main

import "fmt"

const NMAX int = 100000

type proyek struct {
	nama, teknologi, kategori string
	kesulitan, tanggal int
}

type tabInt [NMAX]proyek

func tambahProyek(A *tabInt, n *int, p proyek) {
	if *n < NMAX {	
		A[*n] = p
		*n = *n + 1
		fmt.Println("Proyek berhasil ditambahkan")
	}
}

func tampilkanProyek(A tabInt, n int) {
	var i int
	if n == 0 {
		fmt.Println("Belum ada proyek")
	} else {
		for i = 0; i < n; i++ {
			fmt.Printf("|=========== Proyek ke-%d ==========|\n", i+1)
			fmt.Printf("|%-10s : %-20s |\n", "Nama", A[i].nama)
			fmt.Printf("|%-10s : %-20s |\n", "Teknologi", A[i].teknologi)
			fmt.Printf("|%-10s : %-20s |\n", "Kategori", A[i].kategori)
			fmt.Printf("|%-10s : %-20d |\n", "Kesulitan", A[i].kesulitan)
			fmt.Printf("|%-10s : %-20d |\n", "Tanggal", A[i].tanggal)
		}
	}
}

func sequentialSearchNama(A *tabInt, n int, nama string) int {
	var i int
	for i = 0; i < n; i++ {
		if (A[i].nama) == nama {
			return i 
		}
	}
	return -1
}

func urutkanKategori(A *tabInt, n int) {
	var i, idx, pass int
	var temp proyek
	
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if (A[i].kategori) < (A[idx].kategori) {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func binarySearchKategori(A tabInt, n int, kategori string) {
	var left, right, middle int
	var i, idx int
	left = 0
	right = n - 1
	idx = -1
	
	for left <= right && idx == -1 {
		middle = (left + right) / 2
		if kategori < A[middle].kategori {
			right = middle - 1
		} else if kategori > A[middle].kategori {
			left = middle + 1
		} else {
			idx = middle
		}
	}
	if idx == -1 {
		fmt.Println("Proyek dengan kategori tersebut tidak ditemukan")
	} else {
		for idx > 0 && A[idx-1].kategori == kategori {
			idx--
		}
		i = idx
		for i < n && A[i].kategori == kategori {
			fmt.Printf("|=========== Proyek ke-%d ==========|\n", i+1)
			fmt.Printf("|%-10s : %-20s |\n", "Nama", A[i].nama)
			fmt.Printf("|%-10s : %-20s |\n", "Teknologi", A[i].teknologi)
			fmt.Printf("|%-10s : %-20s |\n", "Kategori", A[i].kategori)
			fmt.Printf("|%-10s : %-20d |\n", "Kesulitan", A[i].kesulitan)
			fmt.Printf("|%-10s : %-20d |\n", "Tanggal", A[i].tanggal)
			i++
		}
	}
}

func urutkanKesulitan(A *tabInt, n int) {
	var i, idx, pass int
	var temp proyek

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].kesulitan < A[idx].kesulitan {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func urutkanTanggal(A *tabInt, n int) {
	var i, pass int
	var temp proyek

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.tanggal < A[i-1].tanggal {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func jumKategoriProyek(A *tabInt, n int) {
	var ml, dv, dc, nlp, eda int
	var i int
	for i = 0; i < n; i++ {
		switch (A[i].kategori) {
		case "Machine-Learning":
			ml++
		case "Data-Visualization":
			dv++
		case "Data-Cleaning":
			dc++
		case "NLP":
			nlp++
		case "EDA":
			eda++
		}
	}
	fmt.Println("Jumlah statistik proyek per kategori: ")
	fmt.Println("Machine Learning:", ml)
	fmt.Println("Data Visualization:", dv)
	fmt.Println("Data Cleaning:", dc)
	fmt.Println("NLP:", nlp)
	fmt.Println("EDA:", eda)
}

func ubahProyek(A *tabInt, n int, nama string) {
	var index int
	var p proyek
	index = sequentialSearchNama(A, n, nama)
	if index == -1 {
		fmt.Println("Proyek tidak ditemukan")
	} else {
		fmt.Println("Masukkan data baru untuk proyek:")
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&p.nama)
		p.teknologi = pilihTeknologi()
		p.kategori = pilihKategori()
		fmt.Print("Tingkat Kesulitan (1-10): ")
		fmt.Scan(&p.kesulitan)
		fmt.Print("Tanggal Pembuatan (1-31): ")
		fmt.Scan(&p.tanggal)
		A[index] = p
		fmt.Println("Proyek berhasil diubah")
	}
}

func hapusProyek(A *tabInt, n *int, nama string) {
	var index int
	var i int
	index = sequentialSearchNama(A, *n, nama)
	if index == -1 {
		fmt.Println("Proyek tidak ditemukan")
	} else {
		for i = index; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("Proyek berhasil dihapus")
	}
}


func daftarKeahlian(A tabInt, n int) {
	var keahlian [NMAX]string
	var jumlahKeahlian, i, j int
	var status bool

	jumlahKeahlian = 0
	for i = 0; i < n; i++ {
		status = false
		for j = 0; j < jumlahKeahlian; j++ {
			if A[i].teknologi == keahlian[j] {
				status = true
			}
		}
		if !status {
			keahlian[jumlahKeahlian] = A[i].teknologi
			jumlahKeahlian++
		}
	}
	for i = 0; i < jumlahKeahlian; i++ {
		fmt.Println("-", keahlian[i])
	}
	if jumlahKeahlian == 0{
		fmt.Println("Belum ada proyek yang ditambahkan")
	}
}

func pilihTeknologi() string {
	var pilih int
	var status bool
	var teknologi string
	
	status = false
	for !status {
		fmt.Println("Pilih Teknologi: ")
		fmt.Println("1. Python")
		fmt.Println("2. Pandas")
		fmt.Println("3. Matplotlib")
		fmt.Println("4. Scikit-learn")
		fmt.Println("5. NLTK")
		fmt.Print("Masukkan pilihan (1-5): ")
		fmt.Scan(&pilih)
	
		if pilih == 1 {
			teknologi = "Python"
			status = true
		} else if pilih == 2 {
			teknologi = "Pandas"
			status = true
		} else if pilih == 3 {
			teknologi = "Matplotlib"
			status = true
		} else if pilih == 4 {
			teknologi = "Scikit-learn"
			status = true
		} else if pilih == 5 {
			teknologi = "NLTK"
			status = true
		} else {
			fmt.Println("Pilihan tidak valid, mohon masukkan pilihan yang ada!")
		}
	}
	return teknologi
}

func pilihKategori() string {
	var pilih int
	var status bool
	var kategori string

	status = false
	for !status {
		fmt.Println("Pilih Kategori: ")
		fmt.Println("1. Machine Learning")
		fmt.Println("2. Data Visualization")
		fmt.Println("3. Data Cleaning")
		fmt.Println("4. NLP")
		fmt.Println("5. EDA")
		fmt.Print("Masukkan pilihan (1-5): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			kategori = "Machine-Learning"
			status = true
		} else if pilih == 2 {
			kategori = "Data-Visualization"
			status = true
		} else if pilih == 3 {
			kategori = "Data-Cleaning"
			status = true
		} else if pilih == 4 {
			kategori = "NLP"
			status = true
		} else if pilih == 5 {
			kategori = "EDA"
			status = true
		} else {
			fmt.Println("Pilihan tidak valid, mohon masukkan pilihan yang ada!")
		}
	}
	return kategori
}

func main() {
	var data tabInt
	var nData int
	var pilih, index int
	var p proyek
	var cariNama string
	var status bool

	status = false
	for !status {
		fmt.Println()
		fmt.Println("|=========||| SELAMAT DATANG |||=========|")
		fmt.Println("|                                        |")
		fmt.Println("|-----Menu Aplikasi Manajemen Proyek-----|")
		fmt.Println("|1. Tambah Proyek                        |")
		fmt.Println("|2. Tampilkan Semua Proyek               |")
		fmt.Println("|3. Cari Proyek berdasarkan Nama         |")
		fmt.Println("|4. Cari Proyek berdasarkan Kategori     |")
		fmt.Println("|5. Urutkan berdasarkan Kesulitan        |")
		fmt.Println("|6. Urutkan berdasarkan Tanggal          |")
		fmt.Println("|7. Tampilkan Statistik Kategori         |")
		fmt.Println("|8. Ubah Proyek                          |")
		fmt.Println("|9. Hapus Proyek                         |")
		fmt.Println("|10. Daftar Keahlian                     |")
		fmt.Println("|11. Keluar                              |")
		fmt.Println("|========================================|")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			fmt.Println("Mohon untuk menambahkan proyek disarankan memberi nama yang berbeda agar tidak terjadi duplikasi:)")
			fmt.Println()
			fmt.Print("Nama Proyek: ")
			fmt.Scan(&p.nama)
			p.teknologi = pilihTeknologi()
			fmt.Println()
			p.kategori = pilihKategori()
			fmt.Print("Tingkat Kesulitan (1-10): ")
			fmt.Scan(&p.kesulitan)
			if p.kesulitan < 1 || p.kesulitan > 10 {
				fmt.Println("Input tidak valid, Kesulitan harus 1-10!")
			} else {
				fmt.Print("Tanggal Pembuatan (1-31): ")
				fmt.Scan(&p.tanggal)
				if p.tanggal < 1 || p.tanggal > 31 {
					fmt.Println("Input tidak valid, Tanggal harus 1-31!")
				} else {
					tambahProyek(&data, &nData, p)
				}
			}
		} else if pilih == 2 {
			tampilkanProyek(data, nData)
		} else if pilih == 3 {
			fmt.Print("Masukkan nama proyek yang ingin dicari: ")
			fmt.Scan(&cariNama)
			index = sequentialSearchNama(&data, nData, cariNama)
			if index != -1 {
				fmt.Println("Proyek ditemukan:")
				fmt.Println("Nama :", data[index].nama)
				fmt.Println("Teknologi :", data[index].teknologi)
				fmt.Println("Kategori :", data[index].kategori)
				fmt.Println("Kesulitan :", data[index].kesulitan)
				fmt.Println("Tangga1l :", data[index].tanggal)
			} else {
				fmt.Println("Proyek dengan nama tersebut tidak ditemukan")
			}
		} else if pilih == 4 {
			urutkanKategori(&data, nData)
			fmt.Print("Masukkan kategori proyek yang ingin dicari: ")
			fmt.Scan(&cariNama)
			binarySearchKategori(data, nData, cariNama)
		} else if pilih == 5 {
			urutkanKesulitan(&data, nData)
			fmt.Println("Proyek berhasil diurutkan berdasarkan tingkat kesulitan")
		} else if pilih == 6 {
			urutkanTanggal(&data, nData)
			fmt.Println("Proyek berhasil diurutkan berdasarkan tanggal")
		} else if pilih == 7 {
			jumKategoriProyek(&data, nData)
		} else if pilih == 8 {
			fmt.Print("Masukkan nama proyek yang ingin diubah: ")
			fmt.Scan(&cariNama)
			ubahProyek(&data, nData, cariNama)
		} else if pilih == 9 {
			fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
			fmt.Scan(&cariNama)
			hapusProyek(&data, &nData, cariNama)
		} else if pilih == 10 {
			fmt.Println("Daftar keahlian yang digunakan dalam proyek: ")
			daftarKeahlian(data, nData)
		} else {
			if pilih == 11 {
				fmt.Println("Terima kasih sudah menggunakan aplikasi")
			} else {
				if pilih != 1 || pilih != 2 || pilih != 3 || pilih != 4 || pilih != 5 || pilih != 6 || pilih != 7 || pilih != 8 || pilih != 9 || pilih != 10 || pilih != 11 {
					fmt.Println("Pilihan anda tidak valid, mohon memilih 1-11")
				}
			}
			status = true
		}
	}
}
