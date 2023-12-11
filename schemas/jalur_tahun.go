package schemas

type JalurTahun struct {
	ID    int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur string `gorm:"column:jalur" json:"jalur"`
	Tahun int    `gorm:"column:tahun" json:"tahun"`
}

type JalurTahunJoin struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur     string `gorm:"column:jalur" json:"jalur"`
	NamaJalur string `gorm:"column:nama_jalur" json:"nama_jalur"`
	Tahun     int    `gorm:"column:tahun" json:"tahun"`
}

func (*JalurTahun) TableName() string {
	return "jalur_tahun"
}
