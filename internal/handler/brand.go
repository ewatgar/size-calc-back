package handler

import "net/http"

func GetBrands(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBrands"))
}

func CreateBrand(http.ResponseWriter, *http.Request) {

}

func GetBrand(http.ResponseWriter, *http.Request) {

}

func DeleteBrand(http.ResponseWriter, *http.Request) {

}
