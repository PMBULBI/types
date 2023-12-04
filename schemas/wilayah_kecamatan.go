package schemas

type WilayahKecamatan struct {
	IDKecamatan   string `gorm:"column:id_kecamatan;primaryKey" json:"id_kecamatan"`
	IDKota        string `gorm:"column:id_kota" json:"id_kota"`
	NamaKecamatan string `gorm:"column:nama_kecamatan" json:"nama_kecamatan"`
}

func (*WilayahKecamatan) TableName() string {
	return "wilayah_kecamatan"
}
