package main

import "fmt"

// func main() {
// 	var myString string = "Hello, World!"
// 	fmt.Println(myString)
// }
type Mahasiswa struct {
    Nama     string
    Umur     int
    AktifDi  []string // Slice of string
	rektor rektor
}
type rektor struct {
	Nama string
	Umur int
}

type People interface {
	GetName() string
	GetAge() int
}

func main() {
	var myArray = [3]int{1, 2, 3}
	// var intSlice = myArray[:]
	intSlice := append(myArray[0:2], 4)
	fmt.Println(myArray)
	fmt.Println(intSlice)

	myArray3 := [3]int{1, 2, 3}
	myArray4 := [3]int{1, 2, 3}
	fmt.Println(myArray3 == myArray4)

	// myArray5 := [3]int{1, 2, 3}
	// myArray6 := [3]int{1, 2, 3}
	var myMap = map[string]int{
		"apple": 1,
		"banana": 2,
		"cherry": 3,
	}
	fmt.Println(myMap["apple"])

	var mahasiswa1 Mahasiswa = Mahasiswa{
		Nama: "John",
		Umur: 20,
		AktifDi: []string{"IT", "Math"},
		rektor: rektor{
			Nama: "Rektor",
			Umur: 50,
		},
	}
	fmt.Println(mahasiswa1)
	fmt.Println(mahasiswa1.GetName())
	fmt.Println(mahasiswa1.GetAge())
	pointer()
}

func (m Mahasiswa) GetName() string {
	return m.Nama
}

func (m Mahasiswa) GetAge() int {
	return m.Umur
}
func pointer() {
	var a int = 10      // variabel biasa
	var p *int = &a     // pointer yang menunjuk ke alamat a

	fmt.Println("Nilai a:", a)
	fmt.Println("Alamat a:", &a)
	fmt.Println("Isi pointer p:", p)
	fmt.Println("Nilai yang ditunjuk p:", *p)
}