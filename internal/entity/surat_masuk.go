package entity

type SuratMasuk struct {
	KodeSurat    int64 `gorm:"primaryKey"`
	TanggalMasuk string
	NoSurat      string
	TanggalSurat string
	Pengirim     string
	Kepada       string
	Perihal      string
}

func (SuratMasuk) TableName() string {
	return "surat_masuk"
}
