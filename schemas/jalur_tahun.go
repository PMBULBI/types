package schemas

type JalurTahun struct {
	ID      int `gorm:"column:id;primaryKey" json:"id"`
	IDJalur int `gorm:"column:id_jalur" json:"id_jalur"`
	Tahun   int `gorm:"column:tahun" json:"tahun"`
}

type JalurTahunJoin struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	IDJalur   int    `gorm:"column:id_jalur" json:"id_jalur"`
	NamaJalur string `gorm:"column:nama_jalur" json:"nama_jalur"`
	Tahun     int    `gorm:"column:tahun" json:"tahun"`
}

func (*JalurTahun) TableName() string {
	return "jalur_tahun"
}
