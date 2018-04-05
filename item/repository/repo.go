package repository

import (
	model "connectionGorilla/item/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ItemRepository struct {
	DB *sql.DB
}

func GetAll(db *ItemRepository) []model.Item {
	result, err := db.DB.Query("select tblItemId, ItemCode, ItemName, BuyingPrice, SellingPrice, ItemAmount, Pieces from tblItem")
	if err != nil {
		return nil
	}
	defer result.Close()

	item := []model.Item{}
	for result.Next() {
		var p model.Item
		err := result.Scan(&p.ID, &p.ItemCode, &p.ItemName, &p.BuyingPrice, &p.SellingPrice, &p.ItemAmount, &p.Pieces)
		if err != nil {
			return nil
		}
		item = append(item, p) //untuk mengisi slice item
	}
	return item
}

func SelectItem(db *ItemRepository, w http.ResponseWriter, r *http.Request) []model.Item {
	//pertama kita harus mencari ID yang diinginkan
	//ini adalah perintah untuk mencari ID menggunakan mux.Vars
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Fatal(err)
	}

	//lalu kita menarik struct Item
	var p model.Item

	//hasil dari id akan kita simpan kedalam sebuah struct ID
	p.ID = id

	//sekarang kita akan memanggil data menggunakan query SQL
	//dan kita akan menyimpan hasil dari pemanggilan kedalam struct Item
	result, err := db.DB.Query("select tblItemId, ItemCode, ItemName, BuyingPrice, SellingPrice, ItemAmount, Pieces from tblItem where tblItemID=?", &p.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	item := []model.Item{}
	for result.Next() {
		err := result.Scan(&p.ID, &p.ItemCode, &p.ItemName, &p.BuyingPrice, &p.SellingPrice, &p.ItemAmount, &p.Pieces)
		if err != nil {
			return nil
		}
		item = append(item, p) //untuk mengisi slice item
	}
	return item
}

func UpdateItem(db *ItemRepository, w http.ResponseWriter, r *http.Request) []model.Item {
	//pertama kita harus mencari ID nya
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Fatal(err)
	}

	//kita tarik struct item dari model
	var p model.Item

	//decoder berfungsi untuk menterjemahkan biner kedalam bentuk karakter
	//sekarang kita cek apakah sturct item sudah terisi
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&p)
	if err != nil {
		return nil
	}
	defer r.Body.Close()

	//lalu kita inialisasikan ID menggunakan variable p
	//dan kita beri nilai berdasarkan variable params ID
	p.ID = id

	//sekarang kita buat sebuah query untuk update SQL
	item := []model.Item{}
	_, err = db.DB.Exec("update tblItem SET ItemCode=?, ItemName=?, BuyingPrice=?, SellingPrice=?, ItemAmount=?, Pieces=? WHERE tblItemId=?;",
		&p.ItemCode, &p.ItemName, &p.BuyingPrice, &p.SellingPrice, &p.ItemAmount, &p.Pieces, &p.ID)
	if err != nil {
		resultErr := "Error pada perintah result"
		log.Fatal(resultErr)
	}
	item = append(item, p) //untuk mengisi slice item
	return item
}

func DeleteItem(db *ItemRepository, w http.ResponseWriter, r *http.Request) []model.Item {
	//cari ID untuk data yang ingin dihapus
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Fatal(err)
	}

	//sekarang kita menarik sebuah struct item
	var p model.Item

	//lalu kita mendeklarasikan id untuk mencari query nanti
	p.ID = id

	//sekarang kita membutuhkan deklarasi untuk
	//menghapus data dari database
	//kita harus membuat sebuah query nya
	item := []model.Item{}
	_, err = db.DB.Exec("delete from tblItem where tblItemId=?", &p.ID)
	if err != nil {
		resultErr := "Error pada perintah result"
		log.Fatal(resultErr)
	}
	item = append(item, p) //untuk mengisi slice item
	return item
}

func InsertItem(db *ItemRepository, w http.ResponseWriter, r *http.Request) []model.Item {
	//tarik struct Item yang berada di model
	var p model.Item

	//decoder berfungsi untuk menterjemahkan biner kedalam bentuk karakter
	//sekarang kita cek apakah sturct item sudah terisi
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	//if err != nil {
	//	return nil
	//}
	defer r.Body.Close()

	//sekarang buat deklarasi untuk input Item
	item := []model.Item{}
	_, err = db.DB.Exec("insert into tblItem (ItemCode, ItemName, BuyingPrice, SellingPrice, ItemAmount, Pieces) values (?,?,?,?,?,?)",
		&p.ItemCode, &p.ItemName, &p.BuyingPrice, &p.SellingPrice, &p.ItemAmount, &p.Pieces)
	if err != nil {
		resultErr := "Error From Query ..."
		log.Fatal(resultErr)
	}
	item = append(item, p) //untuk mengisi slice item
	return item
}
