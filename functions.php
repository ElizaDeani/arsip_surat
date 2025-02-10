<?php
$conn = mysqli_connect("localhost","root","","db_archivio");
function query($query){
	global $conn;
	$result = mysqli_query($conn,$query);
	$rows =[];
	while ($row = mysqli_fetch_assoc($result)) {
		$rows[]= $row;
	}
	return $rows;
}
function tambah($surat) {
	global $conn;
	$kode = $surat["kode_surat"];
	$waktu = $surat["waktu_masuk"];
	$nomor = $surat["nomor_surat"];
	$tanggal = $surat["tanggal_surat"];
    $perihal = $surat["perihal"];
    $pengirim = $surat["pengirim"];
    $kepada = $surat["kepada"];
	$query = "INSERT INTO surat_masuk VALUES ('$kode','$waktu','$nomor','$tanggal','$perihal','$pengirim','$kepada')";
	mysqli_query($conn, $query);
	return mysqli_affected_rows($conn);
}
function ubah($surat) {
    global $conn;

    $kode_surat = htmlspecialchars($surat["kode_surat"]);
    $waktu_masuk = htmlspecialchars($surat["waktu_masuk"]);
    $nomor_surat = htmlspecialchars($surat["nomor_surat"]);
    $tanggal_surat = htmlspecialchars($surat["tanggal_surat"]);
    $perihal = htmlspecialchars($surat["perihal"]);
    $pengirim = htmlspecialchars($surat["pengirim"]);
    $kepada = htmlspecialchars($surat["kepada"]);

    $query = "UPDATE surat_masuk SET 
              waktu_masuk = '$waktu_masuk',
              nomor_surat = '$nomor_surat',
              tanggal_surat = '$tanggal_surat',
              perihal = '$perihal',
              pengirim = '$pengirim',
              kepada = '$kepada'
              WHERE kode_surat = '$kode_surat'";

    $result = mysqli_query($conn, $query);

    if (!$result) {
        die("Query Error: " . mysqli_error($conn));
    }

    return mysqli_affected_rows($conn);
}

function hapus($kode){
	global $conn;
	mysqli_query($conn, "DELETE FROM surat_masuk WHERE kode_surat=$kode");
	return mysqli_affected_rows($conn);
}
function cari($keyword){
$query = "SELECT * FROM surat_masuk WHERE perihal LIKE '$keyword' ";
return query($query);
}
?>