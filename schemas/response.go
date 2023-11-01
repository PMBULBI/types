package schemas

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
