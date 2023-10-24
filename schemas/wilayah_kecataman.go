package schemas

type WilayahKecamatan struct {
	IDKecamatan   string `gorm:"column:id_kecamatan;primaryKey" json:"id_kecamatan"`
	IDKabupaten   string `gorm:"column:id_kabupaten" json:"id_kabupaten"`
	NamaKecamatan string `gorm:"column:nama_kecamatan" json:"nama_kecamatan"`
}

func (*WilayahKecamatan) TableName() string {
	return "wilayah_kecamatan"
}
