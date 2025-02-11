package entity

type SuratKeluar struct {
	KodeSurat     int64 `gorm:"primaryKey"`
	TanggalKeluar string
	NoSurat       string
	TanggalSurat  string
	Pengirim      string
	Kepada        string
	Perihal       string
}

func (SuratKeluar) TableName() string {
	return "surat_keluar"
}
