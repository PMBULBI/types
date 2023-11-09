package schemas

import (
	"github.com/golang-module/carbon/v2"
)

type ProvinsiResponse struct {
	IDProvinsi   string `json:"id_provinsi"`
	NamaProvinsi string `json:"nama_provinsi"`
}

type KotaResponse struct {
	IDKota     string `json:"id_kota"`
	IDProvinsi string `json:"id_provinsi"`
	NamaKota   string `json:"nama_kota"`
}

type ResponseLogin struct {
	Email string `json:"email_mhs"`
}

type ResponsePendaftaran struct {
	ID              int             `json:"id"`
	NamaMhs         string          `json:"nama_mhs"`
	AsalSekolah     string          `json:"asal_sekolah"`
	EmailMhs        string          `json:"email_mhs"`
	HpMhs           string          `json:"hp_mhs"`
	ProvinsiSekolah string          `json:"provinsi_sekolah"`
	KotaSekolah     string          `json:"kota_sekolah"`
	Password        string          `json:"password"`
	StatusMhs       int             `json:"status_mhs"`
	UsernameAdmin   string          `json:"username_admin"`
	TglDaftarMhs    carbon.DateTime `json:"tgl_daftar_mhs"`
}
