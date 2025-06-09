package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Activity struct {
	Activity     string  `json:"activity"`
	Type         string  `json:"type"`
	Participants int     `json:"participants"`
	Price        float64 `json:"price"`
}

func fetchActivity(id int, ch chan<- Activity, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("https://bored-api.appbrewery.com/random")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var a Activity
	if err := json.NewDecoder(resp.Body).Decode(&a); err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	ch <- a
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan Activity, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fetchActivity(i, ch, &wg)
	}

	wg.Wait()
	close(ch)

	fmt.Println("=== Ide Aktivitas Saat Bosan ===")
	for act := range ch {
		fmt.Printf("- %s (Tipe: %s, Peserta: %d, Harga: %.1f)\n",
			act.Activity, act.Type, act.Participants, act.Price)
	}
}
