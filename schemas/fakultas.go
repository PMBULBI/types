package schemas

type Fakultas struct {
	IDFakultas   int    `gorm:"column:id_fakultas;primaryKey" json:"id_fakultas"`
	NamaFakultas string `gorm:"column:nama_fakultas" json:"nama_fakultas"`
}

func (*Fakultas) TableName() string {
	return "fakultas"
}
