package schemas

import "github.com/golang-module/carbon/v2"

type BiodataJalur struct {
	IdHash     string `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	IDJalur    int    `gorm:"column:id_jalur" json:"id_jalur"`
	TahunLulus int    `gorm:"column:tahun_lulus;not null" json:"tahun_lulus"`
	KodeRef    string `gorm:"column:kode_ref;not null" json:"kode_ref"`
}

type BiodataProdi struct {
	IdHash     string `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	IDFakultas int    `gorm:"column:id_fakultas" json:"id_fakultas"`
	Prodi1     int    `gorm:"column:prodi1;not null" json:"prodi1"`
	Prodi2     int    `gorm:"column:prodi2;not null" json:"prodi2"`
}

type BiodataDataDiri struct {
	IdHash           string      `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	NamaLengkap      string      `gorm:"column:nama_lengkap;not null" json:"nama_lengkap"`
	JenisKelamin     string      `gorm:"column:jenis_kelamin;not null" json:"jenis_kelamin"`
	Nik              string      `gorm:"column:nik" json:"nik"`
	NoPendaftaranKip string      `gorm:"column:no_pendaftaran_kip" json:"no_pendaftaran_kip"`
	Npm              string      `gorm:"column:npm" json:"npm"`
	TempatLahir      string      `gorm:"column:tempat_lahir;not null" json:"tempat_lahir"`
	TanggalLahir     carbon.Date `gorm:"column:tanggal_lahir;not null" json:"tanggal_lahir"`
	Agama            string      `gorm:"column:agama;not null" json:"agama"`
	Alamat           string      `gorm:"column:alamat;not null" json:"alamat"`
	Provinsi         string      `gorm:"column:provinsi;not null" json:"provinsi"`
	Kota             string      `gorm:"column:kota;not null" json:"kota"`
	Kecamatan        string      `gorm:"column:kecamatan" json:"kecamatan"`
	Kelurahan        string      `gorm:"column:kelurahan" json:"kelurahan"`
	KodePos          string      `gorm:"column:kode_pos" json:"kode_pos"`
	Whatsapp         string      `gorm:"column:whatsapp;not null" json:"whatsapp"`
	Email            string      `gorm:"column:email;not null" json:"email"`
}

func (*BiodataJalur) TableName() string {
	return "biodata_jalur"
}

func (*BiodataProdi) TableName() string {
	return "biodata_prodi"
}
func (*BiodataDataDiri) TableName() string {
	return "biodata_data_diri"
}
