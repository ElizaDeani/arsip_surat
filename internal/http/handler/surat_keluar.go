package handler

import (
	"net/http"

	"github.com/ElizaDeani/archivio/internal/http/dto"
	"github.com/ElizaDeani/archivio/internal/service"
	"github.com/ElizaDeani/archivio/pkg/response"
	"github.com/labstack/echo/v4"
)

type SuratKeluarHandler struct {
	suratkeluarService service.SuratKeluarService
}

func NewSuratKeluarHandler(suratKeluarService service.SuratKeluarService) SuratKeluarHandler {
	return SuratKeluarHandler{suratKeluarService}
}

func (h *SuratKeluarHandler) GetSuratKeluar(ctx echo.Context) error {
	surat_keluar, err := h.suratkeluarService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing surat keluar", surat_keluar))
}

func (h *SuratKeluarHandler) GetSuratKeluarByKodeSurat(ctx echo.Context) error {
	var req dto.GetSuratKeluarByKodeSuratRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	surat_keluar, err := h.suratkeluarService.GetByKodeSurat(ctx.Request().Context(), req.KodeSurat)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing surat keluar", surat_keluar))
}

func (h *SuratKeluarHandler) CreateSuratKeluar(ctx echo.Context) error {
	var req dto.CreateSuratKeluarRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.suratkeluarService.Create(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully created surat keluar", nil))
}

func (h *SuratKeluarHandler) UpdateSuratKeluar(ctx echo.Context) error {
	var req dto.UpdateSuratKeluarRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.suratkeluarService.Update(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully updated surat keluar", nil))
}

func (h *SuratKeluarHandler) DeleteSuratKeluar(ctx echo.Context) error {
	var req dto.GetSuratKeluarByKodeSuratRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	surat_keluar, err := h.suratkeluarService.GetByKodeSurat(ctx.Request().Context(), req.KodeSurat)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.suratkeluarService.Delete(ctx.Request().Context(), surat_keluar)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully deleted surat keluar", nil))
}
