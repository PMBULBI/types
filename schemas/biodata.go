package schemas

type BiodataJalur struct {
	IdHash     string `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	IDJalur    int    `gorm:"column:id_jalur" json:"id_jalur"`
	TahunLulus int    `gorm:"column:tahun_lulus;not null" json:"tahun_lulus"`
	KodeRef    string `gorm:"column:kode_ref;not null" json:"kode_ref"`
}

type BiodataProdi struct {
	IdHash     string `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	IDFakultas int    `gorm:"column:id_fakultas" json:"id_fakultas"`
	Prodi1     int    `gorm:"column:prodi1;not null" json:"prodi1"`
	Prodi2     int    `gorm:"column:prodi2;not null" json:"prodi2"`
}

func (*BiodataJalur) TableName() string {
	return "biodata_jalur"
}

func (*BiodataProdi) TableName() string {
	return "biodata_prodi"
}
