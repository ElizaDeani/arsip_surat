package handler

import (
	"net/http"

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
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.suratMasukService.Create(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("succesfuly create a surat masuk", nil))
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
	return ctx.JSON(http.StatusOK, response.SuccessResponse("succesfuly update surat masuks", nil))
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
	return ctx.JSON(http.StatusOK, response.SuccessResponse("succesfuly delete a movie", nil))

}
