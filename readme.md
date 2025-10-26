# 📝 Go Fiber CV PDF Generator

Aplikasi web sederhana yang dibangun menggunakan **Go Fiber** untuk menerima konten CV (Curriculum Vitae) dalam format teks mentah dan mengkonversinya secara dinamis menjadi dokumen PDF yang rapi dan profesional.  

Proyek ini sangat berguna untuk membuat dokumen CV yang konsisten dan siap cetak tanpa memerlukan editor dokumen yang kompleks.

---

## ✨ Fitur Utama

- **Konversi Teks ke PDF**: Mengambil teks CV terstruktur dari input pengguna (format yang dihasilkan dari LLM) dan mengubahnya menjadi format PDF.  
- **Tampilan Profesional**: PDF yang dihasilkan memiliki tata letak yang bersih, termasuk penanganan judul bagian, entri terstruktur (jabatan, perusahaan, tanggal), dan poin-poin deskripsi.  
- **Penanganan Multi-baris**: Mampu menangani konten yang panjang dan poin-poin deskripsi menggunakan `MultiCell` untuk penyesuaian otomatis.  

---

## 🛠️ Teknologi yang Digunakan

- **Backend**: Go  
- **Web Framework**: Fiber (Fast, Express-inspired web framework for Go)  
- **PDF Generation**: gofpdf (PDF document generator for Go)

---

## 🚀 Cara Menjalankan Proyek (Lokal)

Proyek ini membutuhkan **Go runtime environment**.

### 1. Klon Repositori
```bash
git clone [URL_REPOSITORI_ANDA]
cd go-fiber-cv-pdf-generator
