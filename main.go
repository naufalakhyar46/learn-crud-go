package main

import (
	"net/http"

	"github.com/naufalakhyar46/learn-crud-go/controllers/pasien"
)

func main() {
	http.HandleFunc("/", pasien.Index)
	http.HandleFunc("/pasien", pasien.Index)
	http.HandleFunc("/pasien/index", pasien.Index)
	http.HandleFunc("/pasien/add", pasien.Add)
	http.HandleFunc("/pasien/edit", pasien.Edit)
	http.HandleFunc("/pasien/delete", pasien.Delete)

	http.ListenAndServe(":4000", nil)
}
