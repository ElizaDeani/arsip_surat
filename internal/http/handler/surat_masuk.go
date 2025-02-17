package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ElizaDeani/archivio/internal/http/dto"
	"github.com/ElizaDeani/archivio/internal/service"
	"github.com/ElizaDeani/archivio/pkg/response"
	"github.com/labstack/echo/v4"
)

type SuratMasukHandler struct {
	suratMasukService service.SuratMasukService
}

func NewSuratMasukHandler(suratMasukService service.SuratMasukService) SuratMasukHandler {
	return SuratMasukHandler{suratMasukService}
}

func (h *SuratMasukHandler) GetSuratMasuk(ctx echo.Context) error {
	surat_masuk, err := h.suratMasukService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing surat masuk", surat_masuk))
}

func (h *SuratMasukHandler) GetSuratMasukByKodeSurat(ctx echo.Context) error {
	var req dto.GetSuratMasukByKodeSuratRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// Panggil service dengan kodeSurat yang sudah dalam bentuk int64
	surat_masuk, err := h.suratMasukService.GetByKodeSurat(ctx.Request().Context(), req.KodeSurat)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing surat", surat_masuk))
}

func (h *SuratMasukHandler) CreateSuratMasuk(ctx echo.Context) error {
	var req dto.CreateSuratMasukRequest
	if err := ctx.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	if req.WaktuMasuk.IsZero() {
		req.WaktuMasuk = time.Now()
	}

	err := h.suratMasukService.Create(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("succesfuly create surat masuk", nil))
}

func (h *SuratMasukHandler) UpdateSuratMasuk(ctx echo.Context) error {
	var req dto.UpdateSuratMasukRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	err := h.suratMasukService.Update(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("succesfuly update surat masuk", nil))
}

func (h *SuratMasukHandler) DeleteSuratMasuk(ctx echo.Context) error {
	var req dto.GetSuratMasukByKodeSuratRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	ticket, err := h.suratMasukService.GetByKodeSurat(ctx.Request().Context(), req.KodeSurat)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.suratMasukService.Delete(ctx.Request().Context(), ticket)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("succesfuly delete surat masuk", nil))
}

func (h *SuratMasukHandler) UploadLampiran(ctx echo.Context) error {
	// Ambil file dari request
	file, err := ctx.FormFile("lampiran")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "failed to get file"))
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to open file"))
	}
	defer src.Close()

	// Buat folder uploads jika belum ada
	destinationDir := "uploads"
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to create upload directory"))
	}

	// Tentukan lokasi penyimpanan file
	destination := filepath.Join(destinationDir, file.Filename)
	dst, err := os.Create(destination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to create file"))
	}
	defer dst.Close()

	// Simpan file
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to save file"))
	}

	// Ambil kode_surat dari form value dan konversi ke int64
	kodeSuratStr := ctx.FormValue("kode_surat")
	if kodeSuratStr == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "kode_surat is required"))
	}

	kodeSurat, err := strconv.ParseInt(kodeSuratStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid kode_surat format"))
	}

	err = h.suratMasukService.UpdateLampiran(ctx.Request().Context(), kodeSurat, destination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	// Kirim respons sukses
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Lampiran successfully uploaded", map[string]string{"file_path": destination}))
}

func (h *SuratMasukHandler) UploadSurat(ctx echo.Context) error {
	// Ambil file dari request
	file, err := ctx.FormFile("surat")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "failed to get file"))
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to open file"))
	}
	defer src.Close()

	// Validasi ukuran file (maksimal 10MB)
	const maxSize = 10 * 1024 * 1024
	if file.Size > maxSize {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "file size exceeds 10MB limit"))
	}

	// Validasi ekstensi file
	allowedExtensions := map[string]bool{".jpg": true, ".pdf": true, ".doc": true, ".docx": true}
	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid file type, only PDF, DOC, DOCX allowed"))
	}

	// Ambil kode_surat dari form value dan konversi ke int64
	kodeSuratStr := ctx.FormValue("kode_surat")
	if kodeSuratStr == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "kode_surat is required"))
	}

	kodeSurat, err := strconv.ParseInt(kodeSuratStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid kode_surat format"))
	}

	// Periksa apakah kodeSurat ada di database
	suratMasuk, err := h.suratMasukService.GetByKodeSurat(ctx.Request().Context(), kodeSurat)
	if err != nil || suratMasuk == nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "kode_surat not found"))
	}

	// Buat folder uploads jika belum ada
	destinationDir := "uploads/surat"
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to create upload directory"))
	}

	// Buat nama file unik
	newFileName := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	destination := filepath.Join(destinationDir, newFileName)
	dst, err := os.Create(destination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to create file"))
	}
	defer dst.Close()

	// Simpan file
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "failed to save file"))
	}

	// Update database dengan lokasi file
	err = h.suratMasukService.UpdateLampiran(ctx.Request().Context(), kodeSurat, destination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	// Kirim respons sukses
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Surat successfully uploaded", map[string]string{"file_path": destination}))
}
