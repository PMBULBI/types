package schemas

type JalurProdi struct {
	ID           int `gorm:"column:id;primaryKey" json:"id"`
	IDJalur      int `gorm:"column:id_jalur" json:"id_jalur"`
	ProgramStudi int `gorm:"column:prodi" json:"prodi"`
}

type JalurProdiWithProdi struct {
	ID               int    `gorm:"column:id;primaryKey" json:"id"`
	IDJalur          int    `gorm:"column:id_jalur" json:"id_jalur"`
	KodeProgramStudi int    `gorm:"column:prodi" json:"prodi"`
	ProgramStudi     string `gorm:"column:program_studi" json:"program_studi"`
}

func (*JalurProdi) TableName() string {
	return "jalur_prodi"
}
