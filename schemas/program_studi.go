package schemas

type ProgramStudi struct {
	ID               int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Fakultas         int    `gorm:"column:fakultas" json:"fakultas"`
	KodeProgramStudi string `gorm:"column:kode_program_studi" json:"kode_program_studi"`
	ProgramStudi     string `gorm:"column:program_studi" json:"program_studi"`
}

func (*ProgramStudi) TableName() string {
	return "program_studi"
}
