package pasien

import (
	"html/template"
	"net/http"

	"github.com/naufalakhyar46/learn-crud-go/entities"
	"github.com/naufalakhyar46/learn-crud-go/models"
)

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

		// fmt.Println(pasien)
		pasienModel.Create(pasien)
		data := map[string]interface{}{
			"pesan": "Data berhasil disimpan",
		}

		temp, _ := template.ParseFiles("views/pasien/index.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
