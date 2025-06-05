// Anggota
// Naufal Afnan Fadli (103032430019)
// Mohammad Rizki Dwi Sapura(103032430024)
package main

import "fmt"

// konstanta untuk nilai maximal dari array
const NMAX int = 100

// Tipe Bentukan Waktu untuk deadline
type waktu struct {
	hari, bulan, tahun string
}

// Tipe Bentukan Proyek untuk menyimpan informasi proyek: nama, klien, deadline, bayaran, dan status.
type proyek struct {
	nama     string
	klien    string
	deadline waktu
	bayaran  int
	status   string
}

// Array untuk menyimpan banyak data
type tabProyek [NMAX]proyek

func main() {
	var data tabProyek
	var nData, pilihan, pilCari, idx int
	var nama string

	tampilanAwal()

	// Menjalankan menu secara berulang dan memanggil fungsi sesuai pilihan user
	for {
		menu()
		fmt.Print("Pilihan (1/2/3/4/5/6/0): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahProyek(&data, &nData)
		case 2:
			fmt.Print("Nama proyek yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusProyek(&data, &nData, nama)
		case 3:
			fmt.Print("Nama proyek yang ingin diubah statusnya: ")
			fmt.Scan(&nama)
			ubahStatus(&data, nData, nama)
		case 4:
			tampilkanProyek(data, nData)
		case 5:
			fmt.Println("Mencari berdasarkan apa? ")
			fmt.Println("[1] Nama Proyek ")
			fmt.Println("[2] Nama Klien ")
			fmt.Print("Pilihan (1/2): ")
			fmt.Scan(&pilCari)

			switch pilCari {
			case 1:
				fmt.Print("Nama proyek yang ingin dicari: ")
				fmt.Scan(&nama)

				idx = SequentialSearch(data, nData, nama)
				if idx != -1 {
					fmt.Println("________________________________________________________________________")
					fmt.Printf("| %2s | %12s | %7s | %12s | %9s | %11s |\n", "No.", "Nama Proyek", "Klien", "Deadline", "Bayaran", "Status")
					fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
					fmt.Printf("| %2d. | %11s  | %7s | %2s/%2s/%4s | %8d | %11s |\n", idx+1, data[idx].nama, data[idx].klien, data[idx].deadline.hari, data[idx].deadline.bulan, data[idx].deadline.tahun, data[idx].bayaran, data[idx].status)
					fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
				} else {
					fmt.Printf("[ X ] Proyek %s tidak ditemukan\n", nama)
				}
			case 2:
				fmt.Print("Nama Klient yang ingin dicari: ")
				fmt.Scan(&nama)
				insertionSortByKlien(&data, nData)

				idx = binarySearch(data, nData, nama)
				if idx != -1 {
					fmt.Println("________________________________________________________________________")
					fmt.Printf("| %2s | %12s | %7s | %12s | %9s | %11s |\n", "No.", "Nama Proyek", "Klien", "Deadline", "Bayaran", "Status")
					fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
					fmt.Printf("| %2d. | %11s  | %7s | %2s/%2s/%4s | %8d | %11s |\n", idx+1, data[idx].nama, data[idx].klien, data[idx].deadline.hari, data[idx].deadline.bulan, data[idx].deadline.tahun, data[idx].bayaran, data[idx].status)
					fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
				} else {
					fmt.Println("[ X ] Klien tidak ditemukan.")
				}
			}
		case 6:
			fmt.Println("Urutkan berdasarkan:")
			fmt.Println("[1] Deadline")
			fmt.Println("[2] Bayaran")
			fmt.Print("Pilihan (1/2): ")
			fmt.Scan(&pilCari)

			switch pilCari {
			case 1:
				insertionSortByDeadline(&data, nData)
				fmt.Println("[ ✓ ] Data berhasil diurutkan berdasarkan deadline (ascending).")
			case 2:
				selectionSortByBayaran(&data, nData)
				fmt.Println("[ ✓ ] Data berhasil diurutkan berdasarkan bayaran (descending).")
			default:
				fmt.Println("[ X ] Pilihan tidak valid.")
			}
		case 0:
			fmt.Println("[ ✓ ] Keluar dari program")
			fmt.Println("===== Sampai Jumpa!!! =====")
			return
		default:
			fmt.Println("[ X ] Pilihan tidak valid")
		}
	}
}

func tampilanAwal() {
	// Menampilkan tampilan awal aplikasi
	fmt.Println(`
8888888888 8888888b.  8888888888 8888888888 888             d8888 888b    888  .d8888b.  8888888888 
888        888   Y88b 888        888        888            d88888 8888b   888 d88P  Y88b 888        
888        888    888 888        888        888           d88P888 88888b  888 888    888 888        
8888888    888   d88P 8888888    8888888    888          d88P 888 888Y88b 888 888        8888888    
888        8888888P"  888        888        888         d88P  888 888 Y88b888 888        888        
888        888 T88b   888        888        888        d88P   888 888  Y88888 888    888 888        
888        888  T88b  888        888        888       d8888888888 888   Y8888 Y88b  d88P 888        
888        888   T88b 8888888888 8888888888 88888888 d88P     888 888    Y888  "Y8888P"  8888888888
	`)
	fmt.Println("===== Selamat datang!!! =====")
	fmt.Println("FREELANCE, solusi praktis untuk manajemen proyek Anda")
}

func menu() {
	// Menampilkan daftar menu pilihan ke layar
	fmt.Println("====================")
	fmt.Println("	MENU")
	fmt.Println("====================")
	fmt.Println("[1] Tambah Proyek")
	fmt.Println("[2] Hapus Proyek")
	fmt.Println("[3] Ubah Status Proyek")
	fmt.Println("[4] Tampilkan Proyek")
	fmt.Println("[5] Cari Proyek")
	fmt.Println("[6] Urutkan Proyek")
	fmt.Println("[0] Keluar")
}

// === Menu Tambah Proyek ===
func tambahProyek(list *tabProyek, n *int) {
	// I.S: list berisi data proyek yang telah dimasukkan sebelumnya, n menunjukkan jumlah data saat ini
	// F.S: 1 proyek baru ditambahkan ke list dan n bertambah 1
	var status string

	if *n < NMAX {
		fmt.Println("Masukkan data proyek:")
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&list[*n].nama)
		fmt.Print("Klien: ")
		fmt.Scan(&list[*n].klien)
		fmt.Print("Deadline (dd mm yyyy): ")
		fmt.Scan(&list[*n].deadline.hari, &list[*n].deadline.bulan, &list[*n].deadline.tahun)
		fmt.Print("Bayaran: ")
		fmt.Scan(&list[*n].bayaran)
		for {
			fmt.Print("Status (dikerjakan/pending/selesai): ")
			fmt.Scan(&status)
			if status == "dikerjakan" || status == "pending" || status == "selesai" {
				break
			} else {
				fmt.Println("[ X ] Status yang dimasukkan salah! Mohon ulangi kembali")
			}
		}
		list[*n].status = status
		*n++
		fmt.Println("[ ✓ ] Proyek berhasil ditambahkan.")
	} else {
		fmt.Println("[ ! ] Data Penuh")
	}
}

// === Menu Ubah Status ===
func ubahStatus(list *tabProyek, n int, x string) {
	// I.S: list terisi data proyek, x adalah nama proyek yang ingin diubah statusnya
	// F.S: jika proyek ditemukan, statusnya akan diperbarui
	var status string
	var idx int

	idx = SequentialSearch(*list, n, x)

	if idx != -1 {
		for {
			fmt.Print("Status (dikerjakan/pending/selesai): ")
			fmt.Scan(&status)
			if status == "dikerjakan" || status == "pending" || status == "selesai" {
				break
			} else {
				fmt.Println("[ X ] Status yang dimasukkan salah! Mohon ulangi kembali")
			}
		}
		fmt.Println("[ ✓ ] Status telah diperbarui")
	} else {
		fmt.Println("[ X ] Nama proyek tidak ditemukan")
	}
}

// === Menu Tampilkan Proyek ===
func tampilkanProyek(list tabProyek, n int) {
	// I.S: list berisi data proyek
	// F.S: seluruh data proyek dalam list ditampilkan di layar

	fmt.Println("=== Daftar Semua Proyek ===")
	fmt.Println("________________________________________________________________________")
	fmt.Printf("| %2s | %12s | %7s | %12s | %9s | %11s |\n", "No.", "Nama Proyek", "Klien", "Deadline", "Bayaran", "Status")
	fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
	for i := 0; i < n; i++ {
		fmt.Printf("| %2d. | %11s  | %7s | %2s/%2s/%4s | %8d | %11s |\n", i+1, list[i].nama, list[i].klien, list[i].deadline.hari, list[i].deadline.bulan, list[i].deadline.tahun, list[i].bayaran, list[i].status)
		fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
	}

}

// === Menu Hapus Proyek ===
func hapusProyek(list *tabProyek, n *int, x string) {
	// I.S: list berisi data proyek, x adalah nama proyek yang ingin dihapus
	// F.S: jika proyek ditemukan, maka proyek dihapus dan n berkurang 1

	var idx int

	idx = SequentialSearch(*list, *n, x)

	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			list[i] = list[i+1]
		}
		*n--
		fmt.Printf("[ ✓ ] Proyek %s telah dihapus\n", x)
	} else {
		fmt.Println("[ X ] Nama proyek tidak ditemukan")
	}
}

// === Menu Cari Proyek ===

func SequentialSearch(list tabProyek, n int, x string) int {
	// I.S: list berisi data proyek
	// F.S: mengembalikan indeks proyek dengan nama x jika ditemukan, -1 jika tidak ditemukan
	var idx, i int

	idx = -1
	for i = 0; i < n; i++ {
		if list[i].nama == x {
			idx = i
		}
	}
	return idx
}

func binarySearch(list tabProyek, n int, x string) int {
	// I.S: list sudah diurutkan berdasarkan nama klien
	// F.S: mengembalikan indeks dari proyek dengan klien x jika ditemukan, -1 jika tidak
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if list[mid].klien == x {
			idx = mid
		} else if list[mid].klien < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func insertionSortByKlien(list *tabProyek, n int) {
	// I.S: list berisi data proyek
	// F.S: list terurut ascending berdasarkan nama klien
	var i, j int
	var temp proyek
	for i = 1; i < n; i++ {
		temp = list[i]
		j = i - 1
		for j >= 0 && list[j].klien > temp.klien {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = temp
	}
}

// Menu Urutkan proyek
func compareTanggal(a, b waktu) bool {
	// Membandingkan dua waktu (deadline). True jika a < b
	if a.tahun != b.tahun {
		return a.tahun < b.tahun
	} else if a.bulan != b.bulan {
		return a.bulan < b.bulan
	}
	return a.hari < b.hari
}

func insertionSortByDeadline(list *tabProyek, n int) {
	// I.S: list berisi data proyek
	// F.S: list terurut ascending berdasarkan tanggal deadline
	var i, j int
	var temp proyek
	for i = 1; i < n; i++ {
		temp = list[i]
		j = i - 1
		for j >= 0 && compareTanggal(temp.deadline, list[j].deadline) {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = temp
	}
}

// Prosedur untuk mengurutkan berdasarkan bayaran (descending) dengan Selection Sort
func selectionSortByBayaran(list *tabProyek, n int) {
	// I.S: list berisi data proyek
	// F.S: list terurut descending berdasarkan bayaran proyek
	var i, j, idx int
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if list[j].bayaran > list[idx].bayaran {
				idx = j
			}
		}
		if idx != i {
			list[i], list[idx] = list[idx], list[i]
		}
	}
}
