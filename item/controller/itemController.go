package controller

import (
	repo "connectionGorilla/item/repository"
	"encoding/json"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//kita ambil connection dari fungsi DBConnection berserta struct nya
	context := Con{}
	d := DBConnection(context.DB)
	defer d.Close()

	//sekarang kita membutuhkan fungsi untuk mengambil data dari DB
	data := &repo.ItemRepository{d}
	item := repo.GetAll(data)
	//olah error nya
	//tampilkan datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}

func SelectItem(w http.ResponseWriter, r *http.Request) {
	//kita ambil connection dari fungsi DBConnection berserta struct nya
	context := Con{}
	d := DBConnection(context.DB)
	defer d.Close()

	//sekarang kita membutuhkan fungsi untuk mengambil data dari DB
	data := &repo.ItemRepository{d}
	item := repo.SelectItem(data, w, r)
	//olah error nya
	//tampilkan datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	//kita ambil connection dari fungsi DBConnection beserta struct nya
	context := Con{}
	d := DBConnection(context.DB)
	defer d.Close()

	//sekarang kita memanggil fungsi untuk update
	//yaitu function UpdateItem di repository
	data := &repo.ItemRepository{d}
	item := repo.UpdateItem(data, w, r)
	//olah error nya
	//tampilkan datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	//kita membuat connection dari fungsi DBConnection beserta struct nya
	context := Con{}
	d := DBConnection(context.DB)
	defer d.Close()

	//sekarang kita memanggil fungsi untuk delete
	//yaitu function DeleteItem di repository
	data := &repo.ItemRepository{d}
	item := repo.DeleteItem(data, w, r)
	//olah error nya
	//tampilkan datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}

func InsertItem(w http.ResponseWriter, r *http.Request) {
	//kita membuat sebuah connection terhadap database dari fungsi DBConnection
	//dan juga memanggil struct dari Connection
	context := Con{}
	d := DBConnection(context.DB)
	defer d.Close()

	//sekarang kita memanggil fungsi untuk insert
	//yaitu function InsertItem di repository
	data := &repo.ItemRepository{d}
	item := repo.InsertItem(data, w, r)
	//olah error nya
	//tampilkan datanya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}
