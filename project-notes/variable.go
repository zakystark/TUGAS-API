package main

import "fmt"

func _main() {
	var nama string = "Fulan" // eksplisit
	var umur = 30.1           // implisit
	isActive := true          // deklarasi singkat

	fmt.Println(nama, umur, isActive)

	nama = "Ahmad"
	umur = 20
	isActive = false

	fmt.Println(nama, umur, isActive)

	const alamat = "Jl. Cisitu Indah VI no 6, Bandung" // konstanta implisit
	fmt.Println(alamat)

	const kota string = "Malang" // konstanta eksplisit
	fmt.Println(kota)
}
