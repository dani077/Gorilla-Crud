package router

import (
	con "connectionGorilla/item/controller"

	"github.com/gorilla/mux"
)

func setItemRouter(router *mux.Router) *mux.Router {
	//perintah membuat sebuah EndPoint
	router.HandleFunc("/item", con.GetAll).Methods("GET")
	router.HandleFunc("/item-select/{ID}", con.SelectItem).Methods("GET")
	router.HandleFunc("/item-update/{ID}", con.UpdateItem).Methods("PUT")
	router.HandleFunc("/item-delete/{ID}", con.DeleteItem).Methods("DELETE")
	router.HandleFunc("/item-insert", con.InsertItem).Methods("POST")
	return router
}
