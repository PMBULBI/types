package schemas

type JalurBerkas struct {
	ID     int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur  int    `gorm:"column:id_jalur" json:"id_jalur"`
	Berkas string `gorm:"column:berkas" json:"berkas"`
}

func (*JalurBerkas) TableName() string {
	return "jalur_berkas"
}
