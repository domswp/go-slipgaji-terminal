# 🧾 Slip Gaji Sederhana pakai golang dan terminal

Project iseng-iseng tapi berfaedah — kamu bisa input nama pegawai, posisi, gaji pokok, dan bonus, terus dapet slip gaji yang tampil rapi di terminal. Cocok banget buat yang lagi belajar `struct`, `input`, dan formatting angka di Golang 🤓

## ✨ Apa Aja yang Dipelajari?
- Cara pakai `fmt.Scanln` buat ambil input user
- Gunain `struct` buat representasi data pegawai
- Format angka jadi `Rp3.500.000` (biar nggak `3.5e+06` 😵)
- Cetak hasilnya dengan gaya ✨ aesthetic ✨

## 💡 Contoh Output
``` Bash
🧾 Slip Gaji Sederhana
Nama : Freya
Posisi : Backend Developer
Gaji Pokok : Rp3.500.000
Bonus : Rp500.000

Total Gaji : Rp4.000.000
```

## 🛠️ Yang Kamu Butuhin
- Golang v1.20+ 
- Package: `golang.org/x/text/message`

Install dulu dependensinya kalau belum:
```bash
go get golang.org/x/text/message
```

## ▶️ Cara Jalanin
```Bash
go run slip-gaji.go
```

## 📁 Struktur
```Bash
.
├── go.mod
└── slip-gaji.go
```

##🔖 Lisensi

MIT License. Boleh banget dipake ulang, dimodif, atau dijadiin bahan belajar. Kalau memang bermanfaat jangan lupa kasih kredit kecil-kecilan ya, terima kasih
