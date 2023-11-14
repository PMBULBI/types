package schemas

import "github.com/golang-module/carbon/v2"

type Admin struct {
	IDAdmin       int           `gorm:"column:id_admin;primaryKey;autoIncrement:true" json:"id_admin"`
	NamaAdmin     string        `gorm:"column:nama_admin;not null" json:"nama_admin"`
	UsernameAdmin string        `gorm:"column:username;not null" json:"username"`
	Email         string        `gorm:"column:email;not null" json:"email"`
	NoHp          string        `gorm:"column:no_hp;not null" json:"no_hp"`
	Password      string        `gorm:"column:password;not null" json:"password"`
	TglBuatAkun   carbon.Carbon `gorm:"column:tgl_buat_akun;not null" json:"tgl_buat_akun"`
	IsAktif       bool          `gorm:"column:aktif" json:"aktif"`
}

func (*Admin) TableName() string {
	return "admin"
}
