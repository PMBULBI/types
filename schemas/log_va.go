package schemas

import (
	"github.com/golang-module/carbon/v2"
)

type LogVA struct {
	IDLog           int64         `gorm:"column:id_log;primaryKey;autoIncrement:true" json:"id_log"`
	Email           string        `gorm:"column:email" json:"email"`
	TrxID           string        `gorm:"column:trx_id" json:"trx_id"`
	VirtualAccount  string        `gorm:"column:virtual_account" json:"virtual_account"`
	TrxAmount       int           `gorm:"column:trx_amount" json:"trx_amount"`
	JenisBayar      string        `gorm:"column:jenis_bayar" json:"jenis_bayar"`
	TglCreate       carbon.Carbon `gorm:"column:tgl_create" json:"tgl_create"`
	TglExpiredVa    carbon.Carbon `gorm:"column:tgl_expired_va" json:"tgl_expired_va"`
	StatusBayar     string        `gorm:"column:status_bayar;default:BELUM" json:"status_bayar"`
	StatusExtendExp string        `gorm:"column:status_extend_exp" json:"status_extend_exp"`
}

func (*LogVA) TableName() string {
	return "log_va"
}
