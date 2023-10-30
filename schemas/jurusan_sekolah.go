package schemas

type JurusanSekolah struct {
	IDJurusan      int    `gorm:"column:id_jurusan;primaryKey;autoIncrement:true" json:"id_jurusan"`
	IDJenisSekolah int    `gorm:"column:id_jenis_sekolah;not null" json:"id_jenis_sekolah"`
	NamaJurusan    string `gorm:"column:nama_jurusan" json:"nama_jurusan"`
}

func (*JurusanSekolah) TableName() string {
	return "jurusan_sekolah"
}
