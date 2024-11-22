package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main_() {
	array()
	slice()
	hashTable()
	fulan := User{
		Name:  "Fulan",
		Email: "fulan@gmail.com",
	}
	fmt.Println(fulan.Name, fulan.Email)
	slice()
	hashTable()
	looping()
}

func looping() {
	var buah = []string{"Pisang", "Salak", "Melon"}
	user := map[string]string{ // nama : email
		"Fulan":   "fulan@gmail.com",
		"Fulanah": "fulanah@gmail.com",
	}
	for key, value := range buah {
		fmt.Println(key, value)
	}
	for key, value := range user {
		fmt.Println(key, value)
	}
}

func hashTable() {
	var scores map[string]int
	scores = make(map[string]int)
	scores["Fulan"] = 100
	scores["Fulanah"] = 90
	scores["Ahmad"] = 80

	fmt.Println(scores["Fulan"])

	// A, B, C
	// (go) rune = char
	grades := map[string]rune{
		"Fulan":   'A',
		"Fulanah": 'b',
	}
	fmt.Println(string(grades["Fulan"]))

	for nama, grade := range grades {
		fmt.Println(nama, string(grade))
	}
}

func slice() {
	// memori yang digunakan = memori yang tersimpan
	var user []string
	user = append(user, "Fulan")
	user = append(user, "Fulanah")
	fmt.Println(user[0], user[1])

	// looping slice & array
	for _, name := range user {
		fmt.Println(name)
	}

	ages := []int{1, 4, 2, 5, 7, 10}
	fmt.Println(ages)
}

func array() {
	// memori yang digunakan : 2 (fulan, fulanah)
	// memori yang tersimpan : 3
	var user [3]string
	user[0] = "Fulan"   //
	user[1] = "Fulanah" //
	user[2] = ""

	fmt.Println(user[0], user[1], user[2])

	colors := [2]string{"red", "blue"}
	fmt.Println(colors)
}
