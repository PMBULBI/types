package schemas

type WilayahKelurahan struct {
	IDKelurahan   string `gorm:"column:id_kelurahan;primaryKey" json:"id_kelurahan"`
	IDKecamatan   string `gorm:"column:id_kecamatan" json:"id_kecamatan"`
	NamaKelurahan string `gorm:"column:nama_kelurahan" json:"nama_kelurahan"`
}

func (*WilayahKelurahan) TableName() string {
	return "wilayah_kelurahan"
}
