package main

import (
	"good/goodservice/pkg/models/postgreSQL"
	"log"
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

	shopWS := &goodshopWebService{shop: &recordsShop{records: &postgreSQL.GoodShopModel{DB: db}}}
	newgoodshopWebServer(shopWS)
	shopWS.start()

}
