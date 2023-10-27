package schemas

type WilayahKota struct {
	IDKota     string `gorm:"column:id_kota;primaryKey" json:"id_kota"`
	IDProvinsi string `gorm:"column:id_provinsi" json:"id_provinsi"`
	NamaKota   string `gorm:"column:nama_kota" json:"nama_kota"`
}

func (*WilayahKota) TableName() string {
	return "wilayah_kota"
}
