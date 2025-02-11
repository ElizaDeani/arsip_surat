BEGIN;

CREATE TABLE IF NOT EXISTS surat_keluar (
    kode_surat INT AUTO_INCREMENT PRIMARY KEY,
    tanggal_keluar DATE NOT NULL,
    no_surat VARCHAR(255) NOT NULL,
    tanggal_surat DATE NOT NULL,
    pengirim VARCHAR(255) NOT NULL,
    kepada VARCHAR(255) NOT NULL,
    perihal VARCHAR(255) NOT NULL
);

COMMIT;