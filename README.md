# Go Automated Parking App CLI

Project ini adalah implementasi sederhana dari parking lot automation menggunakan bahasa Go (Golang).
Aplikasi berjalan melalui CLI, membaca perintah dari file input, lalu mengeksekusi operasi parkir sesuai instruksi.

## Fitur
- Membuat parking lot dengan jumlah slot tertentu
- Memarkir kendaraan
- Mengosongkan slot (leave) + menghitung biaya parkir
- Menampilkan status slot yang terisi
- Validasi command & error handling
- Program otomatis berhenti jika terjadi error fatal (misal: park sebelum parking lot dibuat)

## Cara Menjalankan
1. Pastikan sudah menginstall Go 1.20+
2. Clone repository: 
```bash
git clone https://github.com/irfanyahyaa/vocagame-parkee.git
cd vocagame-parkee
```
3. Jalankan aplikasi dengan file input:
```bash
go run . input.txt
```

## Command yang tersedia
- Membuat parking lot sebanyak `n` : create_parking_lot {capacity} 
- Parkir kendaraan : park {car_number} 
- Parkir selesai : leave {car_number} {hours} 
- Status parking lot : status
> note: program ini ekspek command yang sesuai, dan number tidak ada yang minus