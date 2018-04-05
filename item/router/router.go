package router

import (
	"github.com/gorilla/mux"
)

//ini adalah function untuk memanggil muxRouter
func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	//dan sekarang kita membutuhkan sebuah fungsi untuk membuat EndPoint
	router = setItemRouter(router)
	return router
}
