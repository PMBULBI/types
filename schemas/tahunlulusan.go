package schemas

type TahunLulusan struct {
	Tahun int `gorm:"column:tahun" json:"tahun"`
}
