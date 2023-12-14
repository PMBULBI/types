package schemas

type JalurBerkas struct {
	ID      int    `gorm:"column:id;primaryKey" json:"id"`
	IDJalur int    `gorm:"column:id_jalur" json:"id_jalur"`
	Berkas  string `gorm:"column:berkas" json:"berkas"`
}

type JalurBerkasJoin struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	IDJalur   int    `gorm:"column:id_jalur" json:"id_jalur"`
	NamaJalur string `gorm:"column:nama_jalur" json:"nama_jalur"`
	Berkas    string `gorm:"column:berkas" json:"berkas"`
}

func (*JalurBerkas) TableName() string {
	return "jalur_berkas"
}
