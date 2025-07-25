package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Data struct {
	Weight int `json:"weight"`
	Length int `json:"length"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Weight: rand.Intn(30) + 1,  // weight от 1 до 30
			Length: rand.Intn(100) + 1, // length от 1 до 100
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("Echo new 4")

	http.ListenAndServe(":33500", nil)
}
