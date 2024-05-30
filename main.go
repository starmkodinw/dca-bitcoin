package main

import (
	"dca-bitcoin/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	services.BuyBitcoin()
	go services.StartCronJob()

	port := 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
