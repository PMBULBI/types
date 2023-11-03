package schemas

type Request struct {
	NamaMhs         string `gorm:"column:nama_mhs;not null" json:"nama_mhs"`
	AsalSekolah     string `gorm:"column:asal_sekolah;not null" json:"asal_sekolah"`
	EmailMhs        string `gorm:"column:email_mhs;not null" json:"email_mhs"`
	HpMhs           string `gorm:"column:hp_mhs;not null" json:"hp_mhs"`
	ProvinsiSekolah string `gorm:"column:provinsi_sekolah" json:"provinsi_sekolah"`
	KotaSekolah     string `gorm:"column:kota_sekolah" json:"kota_sekolah"`
	UsernameAdmin   string `gorm:"column:username_admin" json:"username_admin"`
}

type RequestDftSekolah struct {
	NamaSekolah string `json:"nama_sekolah"`
}

type RequestNmProvinsi struct {
	NamaProvinsi string `json:"nama_provinsi"`
}

type RequestIdProvinsi struct {
	IdProvinsi string `json:"id_provinsi"`
}

type RequestIdProvinsiNmKota struct {
	IdProvinsi string `json:"id_provinsi"`
	NamaKota   string `json:"nama_kota"`
}

type RequestLogin struct {
	Email    string `json:"email_mhs"`
	Password string `json:"password"`
}
