package schemas

type WilayahKabupaten struct {
	IDKabupaten   string `gorm:"column:id_kabupaten;primaryKey" json:"id_kabupaten"`
	IDProvinsi    string `gorm:"column:id_provinsi" json:"id_provinsi"`
	NamaKabupaten string `gorm:"column:nama_kabupaten" json:"nama_kabupaten"`
}

func (*WilayahKabupaten) TableName() string {
	return "wilayah_kota"
}
