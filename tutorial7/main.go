package main

import "fmt"

// Fungsi generik yang menjumlahkan dua nilai dengan tipe T
func Add[T int | float64 | int64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(3, 5))           // int
	fmt.Println(Add(2.5, 3.1))       // float64
	fmt.Println(Add(int64(10), 20))  // int64
}
