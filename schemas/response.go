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
	ID              int             `json:"id,omitempty"`
	NamaMhs         string          `json:"nama_mhs,omitempty"`
	Tahun           string          `json:"tahun,omitempty"`
	AsalSekolah     string          `json:"asal_sekolah,omitempty"`
	EmailMhs        string          `json:"email_mhs,omitempty"`
	HpMhs           string          `json:"hp_mhs,omitempty"`
	ProvinsiSekolah string          `json:"provinsi_sekolah,omitempty"`
	KotaSekolah     string          `json:"kota_sekolah,omitempty"`
	Password        string          `json:"password,omitempty"`
	StatusMhs       int             `json:"status_mhs,omitempty"`
	UsernameAdmin   string          `json:"username_admin,omitempty"`
	TglDaftarMhs    carbon.DateTime `json:"tgl_daftar_mhs,omitempty"`
}

type ResponseAdmin struct {
	IDAdmin       int             `json:"id_admin,omitempty"`
	NamaAdmin     string          `json:"nama_admin,omitempty"`
	UsernameAdmin string          `json:"username,omitempty"`
	Email         string          `json:"email,omitempty"`
	NoHp          string          `json:"no_hp,omitempty"`
	Password      string          `json:"password,omitempty"`
	TglBuatAkun   carbon.DateTime `json:"tgl_buat_akun,omitempty"`
	IsAktif       bool            `json:"aktif,omitempty"`
	Level         int             `json:"level,omitempty"`
}
