package schemas

type WilayahProvinsi struct {
	IDProvinsi   string `gorm:"column:id_provinsi;primaryKey" json:"id_provinsi"`
	NamaProvinsi string `gorm:"column:nama_provinsi" json:"nama_provinsi"`
}

func (*WilayahProvinsi) TableName() string {
	return "wilayah_provinsi"
}
