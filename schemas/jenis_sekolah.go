package schemas

type JenisSekolah struct {
	IDJenisSekolah   int32  `gorm:"column:id_jenis_sekolah;primaryKey;autoIncrement:true" json:"id_jenis_sekolah"`
	NamaJenisSekolah string `gorm:"column:nama_jenis_sekolah" json:"nama_jenis_sekolah"`
}

func (*JenisSekolah) TableName() string {
	return "jenis_sekolah"
}
