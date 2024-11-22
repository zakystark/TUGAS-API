package main

import "fmt"

func ___main() {
	// percabangan if
	// Grade :
	// 80 - 100 : A
	// 60 - 79 : B
	// sisa : C
	nilai := 59

	if nilai >= 80 && nilai <= 100 { // 80 - 100
		fmt.Println("Grade A")
	} else if nilai >= 60 && nilai <= 79 { // 60 - 79
		fmt.Println("Grade B")
	} else { // selain semua kondisi diatas
		fmt.Println("Grade C")
	}

	// looping normal
	for i := 0; i <= 5; i++ { // kondisi awal; kondisi akhir; perubahan data
		fmt.Println("Looping normal : ", i)
	}

	// looping dengan kondisi saja
	angka := 0       // kondisi awal
	for angka <= 5 { // kondisi akhir
		fmt.Println("Looping kondisi saja : ", angka)
		angka++ // perubahan data
	}

	count := 0
	for { // forever looping (perulangan selamanya)
		count++ // 1, 2, 3, 4
		if count == 3 {
			continue
		}
		fmt.Println("forever looping", count)
		if count == 5 {
			break
		}
	}

	bulan := 13

	switch bulan {
	case 1:
		fmt.Println("januari")
	case 2:
		fmt.Println("februari")
	case 3:
		fmt.Println("maret")
	case 4:
		fmt.Println("april")
	case 5:
		fmt.Println("mei")
	case 6:
		fmt.Println("juni")
	case 7:
		fmt.Println("juli")
	case 8:
		fmt.Println("agustus")
	case 9:
		fmt.Println("september")
	case 10:
		fmt.Println("oktober")
	case 11:
		fmt.Println("november")
	case 12:
		fmt.Println("desember")
	default:
		fmt.Println("bulan tidak ditemukan")
	}

	if bulan == 1 {
		fmt.Println("januari")
	} else if bulan == 2 {
		fmt.Println("februari")
	}
}
