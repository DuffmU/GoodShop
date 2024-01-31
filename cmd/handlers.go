package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (shop *recordsShop) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ааааааааааааа"))
}

func (shop *recordsShop) showRecords(w http.ResponseWriter, r *http.Request) {
	records, err := shop.records.GET()
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(records)
}
