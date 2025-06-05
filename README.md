
<h1 align="center">📊 FREELANCE - Aplikasi Manajemen Proyek</h1>
<p align="center">
  <strong>Aplikasi CLI sederhana untuk mengelola proyek freelance</strong><br>
  Dibuat dengan bahasa pemrograman Go sebagai Tugas Besar Pemrograman Dasar
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Status-Selesai-brightgreen?style=flat-square"/>
  <img src="https://img.shields.io/badge/Bahasa-Go-blue?style=flat-square&logo=go"/>
  <img src="https://img.shields.io/badge/Lisensi-MIT-lightgrey?style=flat-square"/>
</p>

---

## 🧩 Deskripsi

**FREELANCE** adalah aplikasi berbasis terminal (Command Line Interface) yang memungkinkan pengguna untuk:

- Menambahkan, menghapus, dan memperbarui proyek
- Menampilkan semua proyek yang tersimpan
- Mencari proyek berdasarkan nama atau klien
- Mengurutkan proyek berdasarkan deadline atau bayaran

Semua data disimpan dalam array statis dan dikelola menggunakan tipe data bentukan (`struct`) dalam bahasa Go.

---

## ⚙️ Fitur

| Fitur                        | Deskripsi                                                                 |
|-----------------------------|--------------------------------------------------------------------------|
| ➕ Tambah Proyek             | Input nama, klien, deadline, bayaran, dan status proyek                  |
| 🗑️ Hapus Proyek             | Menghapus proyek berdasarkan nama                                        |
| 🔄 Ubah Status              | Mengubah status menjadi `dikerjakan`, `pending`, atau `selesai`          |
| 🔍 Cari Proyek              | Pencarian berdasarkan nama (sequential) atau klien (binary search)       |
| 📋 Tampilkan Semua          | Menampilkan semua proyek dalam format tabel                              |
| 📅 Urutkan Deadline         | Mengurutkan berdasarkan deadline secara ascending (insertion sort)       |
| 💰 Urutkan Bayaran          | Mengurutkan berdasarkan bayaran secara descending (selection sort)       |

---

## 🏗️ Struktur Data

```go
type waktu struct {
	hari, bulan, tahun string
}

type proyek struct {
	nama     string
	klien    string
	deadline waktu
	bayaran  int
	status   string
}

type tabProyek [100]proyek
```

---

## 🚀 Cara Menjalankan

1. Pastikan Go sudah terinstal di sistem kamu:  
   👉 [https://golang.org/dl/](https://golang.org/dl/)

2. Clone repo ini atau salin file `TugasBesar_Kelompok17.go` ke dalam folder kerja kamu

3. Jalankan program dengan perintah berikut:

```bash
go run TugasBesar_Kelompok17.go
```

4. Ikuti menu interaktif yang muncul pada layar

---

## 🧠 Contoh Tampilan Awal

```text
8888888888 8888888b.  8888888888 8888888888 888             d8888 888b    888  .d8888b.  8888888888 
888        888   Y88b 888        888        888            d88888 8888b   888 d88P  Y88b 888        
888        888    888 888        888        888           d88P888 88888b  888 888    888 888        
8888888    888   d88P 8888888    8888888    888          d88P 888 888Y88b 888 888        8888888    
888        8888888P"  888        888        888         d88P  888 888 Y88b888 888        888        
888        888 T88b   888        888        888        d88P   888 888  Y88888 888    888 888        
888        888  T88b  888        888        888       d8888888888 888   Y8888 Y88b  d88P 888        
888        888   T88b 8888888888 8888888888 88888888 d88P     888 888    Y888  "Y8888P"  8888888888 
```

---

## 👥 Anggota Kelompok

| Nama Lengkap              | NIM           |
|---------------------------|---------------|
| Naufal Afnan Fadli        | 103032430019  |
| Mohammad Rizki Dwi Sapura | 103032430024  |

---

## 📄 Lisensi

Kode sumber dalam repositori ini dirilis di bawah [MIT License](LICENSE).  
Silakan digunakan, dimodifikasi, dan dibagikan dengan tetap mencantumkan lisensi ini.

---

> ✨ Terima kasih telah menggunakan aplikasi FREELANCE! Semoga bermanfaat untuk membantu manajemen proyekmu 🚀
