package schemas

type Pekerjaan struct {
	IDPekerjaan   int    `gorm:"column:id_pekerjaan;primaryKey;autoIncrement:true" json:"id_pekerjaan"`
	NamaPekerjaan string `gorm:"column:nama_pekerjaan" json:"nama_pekerjaan"`
}

func (*Pekerjaan) TableName() string {
	return "pekerjaan"
}
