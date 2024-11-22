package main

import "fmt"

func __main() {
	// operator aritmatika
	a := 10
	b := 3
	jumlah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulus := a % b

	fmt.Println("a + b = ", jumlah)
	fmt.Println("a - b = ", kurang)
	fmt.Println("a * b = ", kali)
	fmt.Println("a / b = ", bagi)
	fmt.Println("a % b =", modulus)

	// operator perbandingan (boolean benar / salah)
	fmt.Println("10 == 10", 10 == 10)
	fmt.Println("10 != 10", 10 != 10)
	fmt.Println("10 > 5", 10 > 5)
	fmt.Println("10 >= 10", 10 >= 10)
	fmt.Println("10 < 5", 10 < 5)
	fmt.Println("10 <= 20", 10 <= 20)

	// Operator logika
	fmt.Println("AND", true && false) // semua harus true = true
	fmt.Println("OR", true || false)  // salah satu true = true
	fmt.Println("NOT", !false)        // kebalikan
}
