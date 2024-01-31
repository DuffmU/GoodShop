package main

import (
	"fmt"
	"good/goodservice/pkg/models/postgreSQL"
	"log"
	"net/http"
)

type recordsShop struct {
	records *postgreSQL.GoodShopModel
}

func main() {

	dbname := "postgres://shop:123@localhost:5432/goodshop?sslmode=disable"

	db, err := openDB(dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shop := &recordsShop{records: &postgreSQL.GoodShopModel{DB: db}}

	srv := &http.Server{
		Addr:    ":4000",
		Handler: shop.routes(),
	}

	fmt.Printf("Запуск сервера на %s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)

}
