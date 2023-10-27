package schemas

type ProvinsiResponse struct {
	IDProvinsi   string `json:"id_provinsi"`
	NamaProvinsi string `json:"nama_provinsi"`
}

type KotaRequest struct {
	IDKabupaten   string `gorm:"column:id_kabupaten;primaryKey" json:"id_kabupaten"`
	IDProvinsi    string `gorm:"column:id_provinsi" json:"id_provinsi"`
	NamaKabupaten string `gorm:"column:nama_kabupaten" json:"nama_kabupaten"`
}
