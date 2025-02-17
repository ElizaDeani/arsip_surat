package dto

import "time"

type GetSuratMasukByKodeSuratRequest struct {
	KodeSurat int64 `json:"kode_surat" validate:"required"`
}

type CreateSuratMasukRequest struct {
	KodeSurat    int64     `json:"kode_surat" validate:"required"`
	WaktuMasuk   time.Time `json:"tanggal_masuk" validate:"required"`
	NoSurat      string    `json:"no_surat" validate:"required"`
	TanggalSurat string    `json:"tanggal_surat" validate:"required"`
	Pengirim     string    `json:"pengirim" validate:"required"`
	Kepada       string    `json:"kepada" validate:"required"`
	Perihal      string    `json:"perihal" validate:"required"`
	Lampiran     string    `json:"lampiran" validate:"required"`
}

type UpdateSuratMasukRequest struct {
	KodeSurat    int64     `json:"kode_surat" validate:"required"`
	WaktuMasuk   time.Time `json:"tanggal_masuk" validate:"required"`
	NoSurat      string    `json:"no_surat" validate:"required"`
	TanggalSurat string    `json:"tanggal_surat" validate:"required"`
	Pengirim     string    `json:"pengirim" validate:"required"`
	Kepada       string    `json:"kepada" validate:"required"`
	Perihal      string    `json:"perihal" validate:"required"`
	Lampiran     string    `json:"lampiran" validate:"required"`
}
