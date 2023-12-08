package schemas

type JalurProdi struct {
	ID           int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur        string `gorm:"column:jalur" json:"jalur"`
	ProgramStudi int    `gorm:"column:prodi" json:"prodi"`
}

type JalurProdiWithProdi struct {
	ID               int    `gorm:"column:id;primaryKey" json:"id"`
	Jalur            string `gorm:"column:jalur" json:"jalur"`
	KodeProgramStudi int    `gorm:"column:prodi" json:"prodi"`
	ProgramStudi     string `gorm:"column:program_studi" json:"program_studi"`
}

// ini untuk mengubah nama tabel
func (*JalurProdi) TableName() string {
	return "jalur_prodi"
}
