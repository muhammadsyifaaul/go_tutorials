package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchWebsite(name string, wg *sync.WaitGroup) {
	defer wg.Done() // ketika selesai, beri tahu WaitGroup

	fmt.Printf("Mengambil data dari %s...\n", name)
	time.Sleep(2 * time.Second) // simulasi delay
	fmt.Printf("Selesai ambil data dari %s\n", name)
}

func main() {
	var wg sync.WaitGroup

	websites := []string{
		"https://tokopedia.com",
		"https://shopee.co.id",
		"https://bukalapak.com",
	}

	// Tambah jumlah tugas yang akan dijalankan
	wg.Add(len(websites))

	// Jalankan semua pengambilan data secara paralel
	for _, site := range websites {
		go fetchWebsite(site, &wg)
	}

	// Tunggu semua selesai
	wg.Wait()
	fmt.Println("✅ Semua data selesai diambil. Siap untuk diproses!")
}
