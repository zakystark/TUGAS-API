package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main__() {
	textManipulation()
	numberManipulation()
	dateManipulation()
}

func dateManipulation() {
	now := time.Now() // mengambil tanggal terbaru / sekarang
	fmt.Println(now)
	fmt.Println("Formatted", now.Format("02 Jan 2006 3:04:05 PM")) // format tanggal
	// 02 : kode tanggal
	// Jan : kode nama bulan
	// 01 : kode urutan bulan
	// 06 : tahun (2 angka terakhir)
	// 2006 : tahun (lengkap)
	// 15 : jam (dalam format 24 Jam)
	// 3 : jam (dalam format 12 Jam)
	// 04 : menit
	// 05 : detik
	// PM : AM/PM
	fmt.Println("Parsing Tanggal")
	tanggal := "01-01-01 01:01:01" // tahun-tanggal-bulan jam:menit:detik (format 24 Jam)
	layout := "06-02-01 15:04:05"  // layout tahun-tanggal-bulan jam:menit:detik (format 24 Jam)
	date, err := time.Parse(layout, tanggal)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(date.Format("02 Jan 2006 3:04:05 PM"))
	}

	oneDay := 24 * time.Hour

	fmt.Println("tambah 1 hari", now.Add(oneDay))
	fmt.Println("kurang 3 hari", now.Add(-3*oneDay))

	birthDay, err := time.Parse(layout, "00-20-12 13:50:12")
	umur := now.Sub(birthDay)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(umur.Hours() / 24 / 30 / 12)
	}

	location, err := time.LoadLocation("asia/kuala_lumpur")
	if err != nil { // ada error
		fmt.Println(err.Error())
	} else { // tidak ada error
		fmt.Println(now.In(location))
	}
}

func numberManipulation() {
	fmt.Println(math.Pi)        // pi = 22/7 = 3.14
	fmt.Println(math.Sqrt(64))  // akar kuadrat
	fmt.Println(math.Pow(3, 5)) // pangkat = 3^5
	fmt.Println(math.Pow(2, 3)) // pangkat = 2^3
	fmt.Println(math.Log(25))   // Logaritma
	fmt.Println(math.Sin(15))   // sinus
	fmt.Println(math.Cos(15))   // Cosinus
	fmt.Println(math.Tan(15))   // Tangen

	fmt.Println("Pembulatan")
	fmt.Println("Round", math.Round(3.1), math.Round(3.9))
	fmt.Println("Ceil", math.Ceil(3.1), math.Ceil(3.9))
	fmt.Println("Floor", math.Floor(3.1), math.Floor(3.9))

	fmt.Println("Random")
	fmt.Println(rand.Intn(200)) // random int 0 - 199
	fmt.Println(rand.Float64()) // random float 0.0 - 1.0

	fmt.Println("Data terbesar", math.Max(3, 5)) // 5
	fmt.Println("Data terkecil", math.Min(3, 5)) // 3
}

func textManipulation() {
	name := "Achmad Fulan"
	lowerName := strings.ToLower(name) // mengubah menjadi lower case
	upperName := strings.ToUpper(name) // mengubah menjadi upper case
	titleName := strings.ToTitle(name) // mengubah menjadi title case = upper case

	fmt.Println("LowerCase : ", lowerName)
	fmt.Println("upperCase : ", upperName)
	fmt.Println("titleCase : ", titleName)

	nama2 := "            Fulanah"
	fmt.Println(name, nama2)
	nama2 = strings.TrimSpace(nama2) // menghapus spasi diawal dan akhir kalimat/string
	fmt.Println(name, nama2)

	buah := strings.Split("jeruk,apel", ",") // memisahkan string = array/slice string
	fmt.Println(buah[0])
	fmt.Println(buah[1])

	listBuah := strings.Join(buah, "-") // untuk menggabungkan array/slice string menjadi string
	fmt.Println(listBuah)

	if strings.Contains(strings.ToLower(name), "achmad") { // untuk mencari substring
		fmt.Println("nama mengandung achmad")
	} else {
		fmt.Println("nama tidak mengandung achmad")
	}

	hello := strings.ReplaceAll("Hello World!", "World", "Dunia") // untuk replace / mengganti string
	fmt.Println(hello)

	fmt.Println(strings.Replace("aaaaa", "a", "b", 3))
	fmt.Println(strings.ReplaceAll("aaaaa", "a", "b"))
}
