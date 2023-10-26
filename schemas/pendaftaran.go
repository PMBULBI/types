package schemas

import "github.com/golang-module/carbon/v2"

type Pendaftaran struct {
	ID            int           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NamaMhs       string        `gorm:"column:nama_mhs;not null" json:"nama_mhs"`
	AsalSekolah   string        `gorm:"column:asal_sekolah;not null" json:"asal_sekolah"`
	EmailMhs      string        `gorm:"column:email_mhs;not null" json:"email_mhs"`
	HpMhs         string        `gorm:"column:hp_mhs;not null" json:"hp_mhs"`
	Password      string        `gorm:"column:password;not null" json:"password"`
	StatusMhs     int           `gorm:"column:status_mhs;not null;default:0" json:"status_mhs"`
	UsernameAdmin string        `gorm:"column:username_admin" json:"username_admin"`
	TglDaftarMhs  carbon.Carbon `gorm:"column:tgl_daftar_mhs;not null" json:"tgl_daftar_mhs"`
}

func (*Pendaftaran) TableName() string {
	return "pendaftaran"
}
