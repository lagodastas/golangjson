package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Weight int `json:"weight"`
	Length int `json:"length"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Weight: rand.Intn(30) + 1,  // weight от 1 до 30
			Length: rand.Intn(100) + 1, // length от 1 до 100
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("Echo")

	http.ListenAndServe(":33500", nil)
}
