package main

import "fmt"

type Manusia struct {
	Nama   string
	Umur   int
	Alamat string
	Gender string
}

func (this Manusia) Makan(makanan string) {
	this.Tampilkan(this.Nama + " makan " + makanan)
}

func (this Manusia) Tampilkan(pesan string) {
	fmt.Println(pesan)
}

// non pointer (hasil yang diubah tidak mempengaruhi variable diluar fungsi)
func (this Manusia) UbahNama(namabaru string) {
	this.Nama = namabaru
}

// pointer (hasil yang diubah akan mempengaruhi variable diluar fungsi)
func (this *Manusia) UbahNamaPointer(namabaru string) {
	this.Nama = namabaru
}

// jika menggunakan embedded type
// maka keseluruhan fungsi dapat dipakai oleh struct/object utama
type Embedded struct {
	Manusia // Embedded Type (Langsung Tipe Data) (nama atribut = nama object)
}

type NonEmbedded struct {
	Manusia Manusia
}

type Hewan interface { // interface = kerangka
	bersuara() // memiliki fungsi bersuara
}

type Kucing struct { // menggunakan interface = menyusun kerangka
}

func (k Kucing) bersuara() { // mendeklarasikan fungsi bersuara (menyusun kerangka hewan kucing)
	fmt.Println("Meong")
}

func (k Kucing) terbang() {
	fmt.Println("Tidak bisa terbang")
}

type Sapi struct {
}

func (s Sapi) bersuara() {
	fmt.Println("Moo")
}

func main() {
	// hewan kucing dengan nama nonik
	var nonik Hewan
	nonik = Kucing{}
	nonik.bersuara()

	// hewan sapi dengan nama paijo
	var paijo Hewan
	paijo = Sapi{}
	paijo.bersuara()

	fulan := Manusia{
		Nama:   "Fulan",
		Umur:   23,
		Alamat: "Malang",
		Gender: "laki-laki",
	}

	fulan.UbahNama("Ahmad") // mengirim value (berupa hanya data)
	fulan.Makan("indomie")

	fulan.UbahNamaPointer("Ahmad") // mengirim alamat (berupa variable)
	fulan.Makan("indomie")
	embedded := Embedded{
		Manusia: Manusia{
			Nama:   "Fulanah",
			Umur:   23,
			Alamat: "Bandung",
			Gender: "Perempuan",
		},
	}
	nonEmbedded := NonEmbedded{
		Manusia: Manusia{
			Nama:   "Zakyah",
			Umur:   23,
			Alamat: "Jombang",
			Gender: "Perempuan",
		},
	}
	embedded.Makan("Kacang")
	fmt.Println(embedded.Nama)

	nonEmbedded.Manusia.Makan("Kacang")
	fmt.Println(nonEmbedded.Manusia.Nama)
}
