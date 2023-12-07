package schemas

type TahunLulusan struct {
	Tahun int `gorm:"column:tahun" json:"tahun"`
}

func (*TahunLulusan) TableName() string {
	return "tahun_lulusan"
}
