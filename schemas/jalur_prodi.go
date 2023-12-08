package schemas

type JalurProdi struct {
	ID           int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur        string `gorm:"column:jalur" json:"jalur"`
	ProgramStudi int    `gorm:"column:prodi" json:"prodi"`
}

type JalurProdiWithProdi struct {
	ID               int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur            string `gorm:"column:jalur" json:"jalur"`
	KodeProgramStudi int    `gorm:"column:kode_prodi" json:"kode_prodi"`
	ProgramStudi     int    `gorm:"column:prodi" json:"prodi"`
}

func (*JalurProdi) TableName() string {
	return "jalur_prodi"
}
