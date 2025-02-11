package dto

type GetSuratKeluarByKodeSuratRequest struct {
	KodeSurat int64 `json:"kode_surat" validate:"required"`
}

type CreateSuratKeluarRequest struct {
	KodeSurat     int64  `json:"kode_surat" validate:"required"`
	TanggalKeluar string `json:"tanggal_keluar" validate:"required"`
	NoSurat       string `json:"no_surat" validate:"required"`
	TanggalSurat  string `json:"tanggal_surat" validate:"required"`
	Pengirim      string `json:"pengirim" validate:"required"`
	Kepada        string `json:"kepada" validate:"required"`
	Perihal       string `json:"perihal" validate:"required"`
}

type UpdateSuratKeluarRequest struct {
	KodeSurat     int64  `json:"kode_surat" validate:"required"`
	TanggalKeluar string `json:"tanggal_keluar" validate:"required"`
	NoSurat       string `json:"no_surat" validate:"required"`
	TanggalSurat  string `json:"tanggal_surat" validate:"required"`
	Pengirim      string `json:"pengirim" validate:"required"`
	Kepada        string `json:"kepada" validate:"required"`
	Perihal       string `json:"perihal" validate:"required"`
}
