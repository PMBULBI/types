package schemas

type DaftarSekolah struct {
	ID          int    `gorm:"column:id_sekolah;primaryKey;autoIncrement:true" json:"id"`
	NamaSekolah string `gorm:"column:nama_sekolah;" json:"nama_sekolah"`
}

func (*DaftarSekolah) TableName() string {
	return "daftar_sekolah"
}
