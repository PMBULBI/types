package schemas

type JalurProdi struct {
	ID           int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur        string `gorm:"column:jalur" json:"jalur"`
	ProgramStudi int    `gorm:"column:prodi" json:"prodi"`
}

func (*JalurProdi) TableName() string {
	return "jalur_prodi"
}
