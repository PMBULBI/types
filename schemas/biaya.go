package schemas

import (
	"github.com/golang-module/carbon/v2"
)

type Biaya struct {
	IDBiaya      int           `gorm:"column:id_biaya;primaryKey;autoIncrement:true" json:"id_biaya"`
	Prodi        string        `gorm:"column:prodi" json:"prodi"`
	Jalur        string        `gorm:"column:jalur" json:"jalur"`
	NamaBiaya    string        `gorm:"column:nama_biaya" json:"nama_biaya"`
	DppBiaya     string        `gorm:"column:dpp_biaya" json:"dpp_biaya"`
	SppBiaya     string        `gorm:"column:spp_biaya" json:"spp_biaya"`
	JumlahBiaya  string        `gorm:"column:jumlah_biaya" json:"jumlah_biaya"`
	ExpiredBiaya carbon.Carbon `gorm:"column:expired_biaya" json:"expired_biaya"`
}

func (*Biaya) TableName() string {
	return "biaya"
}
