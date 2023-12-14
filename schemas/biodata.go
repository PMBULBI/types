package schemas

import (
	"database/sql"
	"github.com/golang-module/carbon/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BiodataMaster struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	IdJalur       string       `gorm:"column:id_jalur" json:"id_jalur"`
	IdProdi       string       `gorm:"column:id_prodi" json:"id_prodi"`
	IdDataDiri    string       `gorm:"column:id_data_diri" json:"id_data_diri"`
	IdDataOrtu    string       `gorm:"column:id_data_ortu" json:"id_data_ortu"`
	IdDataSekolah string       `gorm:"column:id_data_sek" json:"id_data_sek"`
	IdDataBerkas  string       `gorm:"column:id_data_berkas" json:"id_data_berkas"`
	TglDaftar     sql.NullTime `gorm:"column:tgl_daftar" json:"tgl_daftar"`
}

type BiodataJalur struct {
	IdHash     string        `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	freezed    bool          `gorm:"column:freezed" json:"freezed"`
	IDJalur    int           `gorm:"column:id_jalur" json:"id_jalur"`
	IDJalur2   sql.NullInt16 `gorm:"column:id_jalur2" json:"id_jalur2"`
	TahunLulus int           `gorm:"column:tahun_lulus;not null" json:"tahun_lulus"`
	KodeRef    string        `gorm:"column:kode_ref;not null" json:"kode_ref"`
}

type BiodataProdi struct {
	IdHash      string        `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	freezed     bool          `gorm:"column:freezed" json:"freezed"`
	IDFakultas  int           `gorm:"column:id_fakultas" json:"id_fakultas"`
	Prodi1      int           `gorm:"column:prodi1;not null" json:"prodi1"`
	Prodi2      int           `gorm:"column:prodi2;not null" json:"prodi2"`
	Konsentrasi sql.NullInt16 `gorm:"column:konsentrasi" json:"konsentrasi"`
}

type BiodataDataDiri struct {
	IdHash           string         `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	freezed          bool           `gorm:"column:freezed" json:"freezed"`
	JenisKelamin     string         `gorm:"column:jenis_kelamin;not null" json:"jenis_kelamin"`
	Nik              string         `gorm:"column:nik" json:"nik"`
	TempatLahir      string         `gorm:"column:tempat_lahir;not null" json:"tempat_lahir"`
	Agama            string         `gorm:"column:agama;not null" json:"agama"`
	Alamat           string         `gorm:"column:alamat;not null" json:"alamat"`
	Provinsi         string         `gorm:"column:provinsi;not null" json:"provinsi"`
	Kota             string         `gorm:"column:kota;not null" json:"kota"`
	Kecamatan        string         `gorm:"column:kecamatan" json:"kecamatan"`
	Kelurahan        string         `gorm:"column:kelurahan" json:"kelurahan"`
	KodePos          string         `gorm:"column:kode_pos" json:"kode_pos"`
	TglDaftar        sql.NullTime   `gorm:"column:tgl_daftar" json:"tgl_daftar"`
	StatusKelulusan  sql.NullString `gorm:"column:status_kelulusan" json:"status_kelulusan"`
	TglDaftarUlang   sql.NullTime   `gorm:"column:tgl_daftar_ulang" json:"tgl_daftar_ulang"`
	NoPendaftaranKip sql.NullString `gorm:"column:no_pendaftaran_kip" json:"no_pendaftaran_kip"`
	JumlahSks        sql.NullInt16  `gorm:"column:jumlah_sks" json:"jumlah_sks"`
	Npm              sql.NullString `gorm:"column:npm" json:"npm"`
	TanggalLahir     carbon.Carbon  `gorm:"column:tanggal_lahir;not null" json:"tanggal_lahir"`
}

type BiodataDataOrtu struct {
	IdHash                  string `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	freezed                 bool   `gorm:"column:freezed" json:"freezed"`
	NamaAyahKandung         string `gorm:"column:nama_ayah_kandung;not null" json:"nama_ayah_kandung"`
	HpAyahKandung           string `gorm:"column:hp_ayah_kandung;not null" json:"hp_ayah_kandung"`
	NamaIbuKandung          string `gorm:"column:nama_ibu_kandung;not null" json:"nama_ibu_kandung"`
	HpIbuKandung            string `gorm:"column:hp_ibu_kandung" json:"hp_ibu_kandung"`
	PekerjaanOrangTuaWali   string `gorm:"column:pekerjaan_orang_tua_wali;not null" json:"pekerjaan_orang_tua_wali"`
	AlamatOrangTuaWali      string `gorm:"column:alamat_orang_tua_wali;not null" json:"alamat_orang_tua_wali"`
	PenghasilanOrangTuaWali string `gorm:"column:penghasilan_orang_tua_wali;not null" json:"penghasilan_orang_tua_wali"`
	SumberDana              string `gorm:"column:sumber_dana" json:"sumber_dana"`
}

type BiodataDataSekolah struct {
	IdHash            string         `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	freezed           bool           `gorm:"column:freezed" json:"freezed"`
	Nisn              string         `gorm:"column:nisn" json:"nisn,omitempty"`
	AsalJurusan       string         `gorm:"column:asal_jurusan" json:"asal_jurusan"`
	AsalSekolah       string         `gorm:"column:asal_sekolah;not null" json:"asal_sekolah"`
	AlamatSekolah     string         `gorm:"column:alamat_sekolah;not null" json:"alamat_sekolah"`
	KotaSekolah       string         `gorm:"column:kota_sekolah;not null" json:"kota_sekolah"`
	ProvinsiSekolah   string         `gorm:"column:provinsi_sekolah;not null" json:"provinsi_sekolah"`
	KecamatanSekolah  string         `gorm:"column:kecamatan_sekolah" json:"kecamatan_sekolah"`
	KelurahanSekolah  string         `gorm:"column:kelurahan_sekolah" json:"kelurahan_sekolah"`
	KodePosSekolah    string         `gorm:"column:kode_pos_sekolah" json:"kode_pos_sekolah"`
	JenisSekolah      string         `gorm:"column:jenis_sekolah" json:"jenis_sekolah"`
	Jurusan           string         `gorm:"column:jurusan" json:"jurusan"`
	AkreditasiSekolah string         `gorm:"column:akreditasi_sekolah" json:"akreditasi_sekolah"`
	TahunLulus        int32          `gorm:"column:tahun_lulus" json:"tahun_lulus"`
	GuruBk            string         `gorm:"column:guru_bk" json:"guru_bk,omitempty"`
	HpGuruBk          string         `gorm:"column:hp_guru_bk" json:"hp_guru_bk,omitempty"`
	JenjangTerakhir   sql.NullString `gorm:"column:jenjang_terakhir" json:"jenjang_terakhir"`
	LamaBekerja       sql.NullString `gorm:"column:lama_bekerja" json:"lama_bekerja"`
	BekerjaDimana     sql.NullString `gorm:"column:bekerja_dimana" json:"bekerja_dimana"`
	AsalKampus        sql.NullString `gorm:"column:asal_kampus" json:"asal_kampus"`
	AsalProdi         sql.NullString `gorm:"column:asal_prodi" json:"asal_prodi"`
	AlamatKampus      sql.NullString `gorm:"column:alamat_kampus" json:"alamat_kampus"`
	ProvinsiKampus    sql.NullString `gorm:"column:provinsi_kampus" json:"provinsi_kampus"`
	KotaKampus        sql.NullString `gorm:"column:kota_kampus" json:"kota_kampus"`
	KodePosKampus     sql.NullString `gorm:"column:kode_pos_kampus" json:"kode_pos_kampus"`
	Ipk               sql.NullString `gorm:"column:ipk" json:"ipk"`
}

type BiodataDataBerkas struct {
	IdHash                         string         `gorm:"column:id_hash;primaryKey" json:"id_hash"`
	freezed                        bool           `gorm:"column:freezed" json:"freezed"`
	Nilai                          sql.NullString `gorm:"column:nilai" json:"nilai"`
	NilaiUtbk                      sql.NullString `gorm:"column:nilai_utbk" json:"nilai_utbk"`
	SuratUndangan                  sql.NullString `gorm:"column:surat_undangan" json:"surat_undangan"`
	RaportSemester1                sql.NullString `gorm:"column:raport_semester_1" json:"raport_semester_1"`
	RaportSemester2                sql.NullString `gorm:"column:raport_semester_2" json:"raport_semester_2"`
	RaportSemester3                sql.NullString `gorm:"column:raport_semester_3" json:"raport_semester_3"`
	RaportSemester4                sql.NullString `gorm:"column:raport_semester_4" json:"raport_semester_4"`
	RaportSemester5                sql.NullString `gorm:"column:raport_semester_5" json:"raport_semester_5"`
	RaportSemester6                sql.NullString `gorm:"column:raport_semester_6" json:"raport_semester_6"`
	Photo                          sql.NullString `gorm:"column:photo;not null" json:"photo"`
	JenisPrestasi                  sql.NullString `gorm:"column:jenis_prestasi" json:"jenis_prestasi"`
	SuratRekomendasiSekolah        sql.NullString `gorm:"column:surat_rekomendasi_sekolah" json:"surat_rekomendasi_sekolah"`
	PeringkatParalel               sql.NullString `gorm:"column:peringkat_paralel" json:"peringkat_paralel"`
	SuratRekomendasiKemitraan      sql.NullString `gorm:"column:surat_rekomendasi_kemitraan" json:"surat_rekomendasi_kemitraan"`
	SuratPernyataanRegistrasi1     sql.NullString `gorm:"column:surat_pernyataan_registrasi_1" json:"surat_pernyataan_registrasi_1"`
	SuratPernyataanRegistrasi2     sql.NullString `gorm:"column:surat_pernyataan_registrasi_2" json:"surat_pernyataan_registrasi_2"`
	SuratPernyataanRegistrasi3     sql.NullString `gorm:"column:surat_pernyataan_registrasi_3" json:"surat_pernyataan_registrasi_3"`
	IjazahSttb                     sql.NullString `gorm:"column:ijazah_sttb" json:"ijazah_sttb"`
	Ktp                            sql.NullString `gorm:"column:ktp" json:"ktp"`
	Skhun                          sql.NullString `gorm:"column:skhun" json:"skhun"`
	SuratKelakuanBaik              sql.NullString `gorm:"column:surat_kelakuan_baik" json:"surat_kelakuan_baik"`
	SuratKeteranganBebasNarkoba    sql.NullString `gorm:"column:surat_keterangan_bebas_narkoba" json:"surat_keterangan_bebas_narkoba"`
	PasFotoTerbaru                 sql.NullString `gorm:"column:pas_foto_terbaru" json:"pas_foto_terbaru"`
	AktaKelahiran                  sql.NullString `gorm:"column:akta_kelahiran" json:"akta_kelahiran"`
	SuratKeteranganKepegawaianPos  sql.NullString `gorm:"column:surat_keterangan_kepegawaian_pos" json:"surat_keterangan_kepegawaian_pos"`
	KartuKeluarga                  sql.NullString `gorm:"column:kartu_keluarga" json:"kartu_keluarga"`
	IjazahSklSementara             sql.NullString `gorm:"column:ijazah_skl_sementara" json:"ijazah_skl_sementara"`
	SkhunSklSementara              sql.NullString `gorm:"column:skhun_skl_sementara" json:"skhun_skl_sementara"`
	Kip                            sql.NullString `gorm:"column:kip" json:"kip"`
	Sktm                           sql.NullString `gorm:"column:sktm" json:"sktm"`
	PengajuanKipKuliah             sql.NullString `gorm:"column:pengajuan_kip_kuliah" json:"pengajuan_kip_kuliah"`
	RekeningListrik                sql.NullString `gorm:"column:rekening_listrik" json:"rekening_listrik"`
	SlipGaji                       sql.NullString `gorm:"column:slip_gaji" json:"slip_gaji"`
	FotoKeluarga                   sql.NullString `gorm:"column:foto_keluarga" json:"foto_keluarga"`
	FotoRumahLuar                  sql.NullString `gorm:"column:foto_rumah_luar" json:"foto_rumah_luar"`
	FotoRumahDalam                 sql.NullString `gorm:"column:foto_rumah_dalam" json:"foto_rumah_dalam"`
	KartuPesertaKip                sql.NullString `gorm:"column:kartu_peserta_kip" json:"kartu_peserta_kip"`
	SertifikatUtbk                 sql.NullString `gorm:"column:sertifikat_utbk" json:"sertifikat_utbk"`
	NoIjazah                       sql.NullString `gorm:"column:no_ijazah" json:"no_ijazah"`
	NilaiMatematikaSmt1            sql.NullString `gorm:"column:nilai_matematika_smt_1" json:"nilai_matematika_smt_1"`
	NilaiMatematikaSmt2            sql.NullString `gorm:"column:nilai_matematika_smt_2" json:"nilai_matematika_smt_2"`
	NilaiMatematikaSmt3            sql.NullString `gorm:"column:nilai_matematika_smt_3" json:"nilai_matematika_smt_3"`
	NilaiMatematikaSmt4            sql.NullString `gorm:"column:nilai_matematika_smt_4" json:"nilai_matematika_smt_4"`
	NilaiMatematikaSmt5            sql.NullString `gorm:"column:nilai_matematika_smt_5" json:"nilai_matematika_smt_5"`
	NilaiIndonesiaSmt1             sql.NullString `gorm:"column:nilai_indonesia_smt_1" json:"nilai_indonesia_smt_1"`
	NilaiIndonesiaSmt2             sql.NullString `gorm:"column:nilai_indonesia_smt_2" json:"nilai_indonesia_smt_2"`
	NilaiIndonesiaSmt3             sql.NullString `gorm:"column:nilai_indonesia_smt_3" json:"nilai_indonesia_smt_3"`
	NilaiIndonesiaSmt4             sql.NullString `gorm:"column:nilai_indonesia_smt_4" json:"nilai_indonesia_smt_4"`
	NilaiIndonesiaSmt5             sql.NullString `gorm:"column:nilai_indonesia_smt_5" json:"nilai_indonesia_smt_5"`
	NilaiInggrisSmt1               sql.NullString `gorm:"column:nilai_inggris_smt_1" json:"nilai_inggris_smt_1"`
	NilaiInggrisSmt2               sql.NullString `gorm:"column:nilai_inggris_smt_2" json:"nilai_inggris_smt_2"`
	NilaiInggrisSmt3               sql.NullString `gorm:"column:nilai_inggris_smt_3" json:"nilai_inggris_smt_3"`
	NilaiInggrisSmt4               sql.NullString `gorm:"column:nilai_inggris_smt_4" json:"nilai_inggris_smt_4"`
	NilaiInggrisSmt5               sql.NullString `gorm:"column:nilai_inggris_smt_5" json:"nilai_inggris_smt_5"`
	KeteranganPeringkat            sql.NullString `gorm:"column:keterangan_peringkat" json:"keterangan_peringkat"`
	KemampuanPenalaranUmum         sql.NullString `gorm:"column:kemampuan_penalaran_umum" json:"kemampuan_penalaran_umum"`
	KemampuanMemahamiBacaanMenulis sql.NullString `gorm:"column:kemampuan_memahami_bacaan_menulis" json:"kemampuan_memahami_bacaan_menulis"`
	PengetahuanPemahamanUmum       sql.NullString `gorm:"column:pengetahuan_pemahaman_umum" json:"pengetahuan_pemahaman_umum"`
	PengetahuanKuantitatif         sql.NullString `gorm:"column:pengetahuan_kuantitatif" json:"pengetahuan_kuantitatif"`
	SuratKeteranganPindah          sql.NullString `gorm:"column:surat_keterangan_pindah" json:"surat_keterangan_pindah"`
	CurriculumVitae                sql.NullString `gorm:"column:curriculum_vitae" json:"curriculum_vitae"`
	SertifikatPelatihan            sql.NullString `gorm:"column:sertifikat_pelatihan" json:"sertifikat_pelatihan"`
	SertifikatPengalamanBekerja    sql.NullString `gorm:"column:sertifikat_pengalaman_bekerja" json:"sertifikat_pengalaman_bekerja"`
	SuratRekomendasiPimpinan       sql.NullString `gorm:"column:surat_rekomendasi_pimpinan" json:"surat_rekomendasi_pimpinan"`
	PortofolioLainnya              sql.NullString `gorm:"column:portofolio_lainnya" json:"portofolio_lainnya"`
	SlipGajiSaudara                sql.NullString `gorm:"column:slip_gaji_saudara" json:"slip_gaji_saudara"`
	MotivationLetter               sql.NullString `gorm:"column:motivation_letter" json:"motivation_letter"`
	BuktiDtks                      sql.NullString `gorm:"column:bukti_dtks" json:"bukti_dtks"`
	SuratPernyataanKip             sql.NullString `gorm:"column:surat_pernyataan_kip" json:"surat_pernyataan_kip"`
	Sertifikasi                    sql.NullString `gorm:"column:sertifikasi" json:"sertifikasi"`
	Skmk                           sql.NullString `gorm:"column:skmk" json:"skmk"`
	SuratRekomendasiPartai         sql.NullString `gorm:"column:surat_rekomendasi_partai" json:"surat_rekomendasi_partai"`
	SuratKeteranganBekerja         sql.NullString `gorm:"column:surat_keterangan_bekerja" json:"surat_keterangan_bekerja"`
	SuratKeteranganSehat           sql.NullString `gorm:"column:surat_keterangan_sehat" json:"surat_keterangan_sehat"`
	SkPengangkatanKaryawan         sql.NullString `gorm:"column:sk_pengangkatan_karyawan" json:"sk_pengangkatan_karyawan"`
}

type BerkasPendaftaran struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Nilai         *[]DataNilaiSemester `bson:"nilai" json:"nilai"`
	FilesSemester *[]FileSemester      `bson:"files_semester" json:"files_semester"`
	FileGeneric   *[]FileGeneric       `bson:"file_generic" json:"file_generic"`
}

type DataNilaiSemester struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Semester      int                `bson:"semester" json:"id_semester"`
	MataPelajaran string             `bson:"mata_pelajaran" json:"mata_pelajaran"`
	Nilai         string             `bson:"nilai" json:"nilai"`
}

func (*BiodataMaster) TableName() string {
	return "biodata_master"
}

func (*BiodataJalur) TableName() string {
	return "biodata_jalur"
}

func (*BiodataProdi) TableName() string {
	return "biodata_prodi"
}

func (*BiodataDataDiri) TableName() string {
	return "biodata_data_diri"
}

func (*BiodataDataOrtu) TableName() string {
	return "biodata_data_ortu"
}

func (*BiodataDataSekolah) TableName() string {
	return "biodata_data_sekolah"
}
