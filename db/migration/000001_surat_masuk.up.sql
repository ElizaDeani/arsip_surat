BEGIN;

CREATE TABLE IF NOT EXISTS surat_masuk (
    kode_surat INT AUTO_INCREMENT PRIMARY KEY,
    waktu_masuk DATE NULL,
    no_surat VARCHAR(255) NOT NULL,
    tanggal_surat DATE NULL,
    perihal VARCHAR(255) NOT NULL,
    pengirim VARCHAR(255) NOT NULL,
    kepada VARCHAR(255) NOT NULL,
    lampiran VARCHAR(255) NOT NULL
);

COMMIT;
