package schemas

type ProvinsiResponse struct {
	IDProvinsi   string `json:"id_provinsi"`
	NamaProvinsi string `json:"nama_provinsi"`
}

type KotaResponse struct {
	IDKabupaten   string `json:"id_kabupaten"`
	IDProvinsi    string `json:"id_provinsi"`
	NamaKabupaten string `json:"nama_kabupaten"`
}
