package schemas

type Konsentrasi struct {
	IDKonsentrasi   int32  `gorm:"column:id_konsentrasi;primaryKey;autoIncrement:true" json:"id_konsentrasi"`
	NamaKonsentrasi string `gorm:"column:nama_konsentrasi" json:"nama_konsentrasi"`
}

func (*Konsentrasi) TableName() string {
	return "konsentrasi"
}
