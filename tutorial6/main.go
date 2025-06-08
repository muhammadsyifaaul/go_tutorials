// package main

// import "fmt"

// func main() {
//     ch := make(chan string) // buat channel untuk string

//     // goroutine yang mengirim data ke channel
//     go func() {
//         ch <- "Halo dari goroutine!"
//     }()

//     // menerima data dari channel (blocking sampai ada data)
//     msg := <-ch
//     fmt.Println(msg)
// }
package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	time.Sleep(time.Duration(id) * time.Second) // simulasi kerja berbeda-beda
	ch <- fmt.Sprintf("Halo dari worker %d", id)
}

func main() {
	ch := make(chan string) // channel tanpa buffer

	// Jalankan 3 goroutine worker
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Terima 3 pesan dari channel
	for i := 1; i <= 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}
