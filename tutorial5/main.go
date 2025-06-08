package main

import (
	"fmt"
	"time"
)

func cetakPesan() {
	fmt.Println("Halo dari Goroutine!")
}

func main() {
	go cetakPesan() // dijalankan sebagai goroutine (concurrent)
	fmt.Println("Halo dari main!")
	

	time.Sleep(1 * time.Second) // beri waktu agar goroutine selesai
	go func() {
		fmt.Println("Ini dari anonymous goroutine")
	}()
	time.Sleep(500 * time.Millisecond)
}
