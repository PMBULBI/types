package schemas

import (
	"github.com/golang-module/carbon/v2"
)

type DiskonPeriode struct {
	IDDiskon       int           `gorm:"column:id_diskon;primaryKey;autoIncrement:true" json:"id_diskon"`
	Jalur          string        `gorm:"column:jalur" json:"jalur"`
	JumlahDiskon   string        `gorm:"column:jumlah_diskon" json:"jumlah_diskon"`
	TanggalDiskon  carbon.Carbon `gorm:"column:tanggal_diskon" json:"tanggal_diskon"`
	TanggalExpired carbon.Carbon `gorm:"column:tanggal_expired" json:"tanggal_expired"`
}

func (*DiskonPeriode) TableName() string {
	return "diskon_periode"
}
