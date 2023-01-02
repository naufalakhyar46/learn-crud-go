package pasien

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/naufalakhyar46/learn-crud-go/entities"
	"github.com/naufalakhyar46/learn-crud-go/libraries"
	"github.com/naufalakhyar46/learn-crud-go/models"
)

var validation = libraries.NewValidation()
var pasienModel = models.NewPasienModel()

func Index(response http.ResponseWriter, request *http.Request) {
	// fmt.Println("It's Working")
	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("nohp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil {
			data["pasien"] = pasien
			data["validasi"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil disimpan"
			pasienModel.Create(pasien)
		}
		// fmt.Println(pasien)
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
		// http.Redirect(response, request, "/", http.StatusFound)
		// return
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var pasien entities.Pasien
		pasienModel.Find(id, &pasien)

		data := map[string]interface{}{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var pasien entities.Pasien
		pasien.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("nohp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil {
			data["pasien"] = pasien
			data["validasi"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil diperbarui"
			pasienModel.Update(pasien)
		}
		temp, _ := template.ParseFiles("views/pasien/edit.html")
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()

	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	pasienModel.Delete(id)

	http.Redirect(response, request, "/pasien", http.StatusSeeOther)

}
