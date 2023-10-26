package schemas

type JalurPendaftaran struct {
	IDJalur         int32  `gorm:"column:id_jalur;primaryKey;autoIncrement:true" json:"id_jalur"`
	Jalur           string `gorm:"column:jalur" json:"jalur"`
	NamaJalur       string `gorm:"column:nama_jalur" json:"nama_jalur"`
	KeteranganJalur string `gorm:"column:keterangan_jalur" json:"keterangan_jalur"`
	Status          string `gorm:"column:status;default:nonaktif" json:"status"`
}

func (*JalurPendaftaran) TableName() string {
	return "jalur_pendaftaran"
}
