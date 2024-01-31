package main

import "net/http"

func (shop *recordsShop) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", shop.home)
	mux.HandleFunc("/records", shop.showRecords)

	return mux
}
