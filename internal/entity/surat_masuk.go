package entity

import "time"

type SuratMasuk struct {
	KodeSurat    int64 `gorm:"primaryKey"`
	WaktuMasuk   time.Time
	NoSurat      string
	TanggalSurat string
	Pengirim     string
	Kepada       string
	Perihal      string
	Lampiran     string
}

func (SuratMasuk) TableName() string {
	return "surat_masuk"
}
