package schemas

type Agama struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Agama string `gorm:"column:agama;not null" json:"agama"`
}

func (*Agama) TableName() string {
	return "agama"
}
