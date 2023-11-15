package schemas

type AdminLevel struct {
	IDLevel   int    `gorm:"column:id_level;primaryKey;autoIncrement:true" json:"id_level,omitempty"`
	NamaLevel string `gorm:"column:nama_level;not null" json:"nama_level,omitempty"`
}

func (*AdminLevel) TableName() string {
	return "admin_level"
}
