package main

import (
	router "connectionGorilla/item/router"
	"log"
	"net/http"
)

func main() {
	//kita membutuhkan sebuah router untuk dapat memanggil EndPoint
	router := router.InitRouter()

	//buat konfigurasi server
	log.Fatal(http.ListenAndServe(":8080", router))
}
