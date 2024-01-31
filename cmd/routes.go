package main

import (
	"fmt"
	"log"
	"net/http"
)

type goodshopWebService struct {
	shop   *recordsShop
	router *http.ServeMux
}

func newgoodshopWebServer(sw *goodshopWebService) *goodshopWebService {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sw.shop.home)
	mux.HandleFunc("/records", sw.shop.showRecords)
	sw.router = mux
	return sw
}

func (sw *goodshopWebService) start() {
	fmt.Printf("Запуск сервера на localhost:4000")
	err := http.ListenAndServe(":4000", sw.router)
	log.Fatal(err)
}
