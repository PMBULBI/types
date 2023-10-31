package schemas

type TahunAkademik struct {
	IDTahunAkademik int    `gorm:"column:id_tahun_akademik;primaryKey;autoIncrement:true" json:"id_tahun_akademik"`
	TahunAkademik   string `gorm:"column:tahun_akademik" json:"tahun_akademik"`
	KodeTahun       int    `gorm:"column:kode_tahun" json:"kode_tahun"`
}

func (*TahunAkademik) TableName() string {
	return "tahun_akademik"
}
