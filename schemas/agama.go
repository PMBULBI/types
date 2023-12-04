package schemas

type Agama struct {
	ID    int    `gorm:"column:id_agama;primaryKey;autoIncrement:true" json:"id_agama"`
	Agama string `gorm:"column:agama;not null" json:"agama"`
}

func (*Agama) TableName() string {
	return "agama"
}
