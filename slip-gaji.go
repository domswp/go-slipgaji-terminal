package main

import (
	"fmt"
	"golang.org/x/text/message"
)

type Pegawai struct {
	Nama      string
	Posisi    string
	GajiPokok float64
	Bonus     float64
}

func main() {
	var nama, posisi string
	var gajiPokok, bonus float64

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan posisi: ")
	fmt.Scanln(&posisi)

	fmt.Print("Masukkan gaji pokok: Rp")
	fmt.Scanln(&gajiPokok)
	if gajiPokok < 0 {
		fmt.Println("❌ Gaji tidak boleh negatif.")
		return
	}

	fmt.Print("Masukkan bonus: Rp")
	fmt.Scanln(&bonus)
	if bonus < 0 {
		fmt.Println("❌ Bonus tidak boleh negatif.")
		return
	}

	pegawai := Pegawai{
		Nama:      nama,
		Posisi:    posisi,
		GajiPokok: gajiPokok,
		Bonus:     bonus,
	}

	totalGaji := pegawai.GajiPokok + pegawai.Bonus
	p := message.NewPrinter(message.MatchLanguage("id"))

	fmt.Println("\n🧾 Slip Gaji Sederhana")
	fmt.Println("------------------------")
	fmt.Println("Nama     :", pegawai.Nama)
	fmt.Println("Posisi   :", pegawai.Posisi)
	p.Printf("Gaji Pokok : Rp%d\n", int(pegawai.GajiPokok))
	p.Printf("Bonus      : Rp%d\n", int(pegawai.Bonus))
	fmt.Println("------------------------")
	p.Printf("Total Gaji : Rp%d\n", int(totalGaji))
}
