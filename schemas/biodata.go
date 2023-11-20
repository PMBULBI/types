package schemas

type BiodataJalur struct {
	IdHash     string `gorm:"column:id_hash;primaryKey;autoIncrement:true" json:"id_hash"`
	IDJalur    int    `gorm:"column:id_jalur" json:"id_jalur"`
	TahunLulus int    `gorm:"column:tahun_lulus;not null" json:"tahun_lulus"`
	KodeRef    string `gorm:"column:kode_ref;not null" json:"kode_ref"`
}

func (*BiodataJalur) TableName() string {
	return "biodata_jalur"
}
