package router

import (
	"net/http"

	"github.com/ElizaDeani/archivio/internal/http/handler"
	"github.com/ElizaDeani/archivio/pkg/route"
)

var (
	adminOnly = []string{"ADMIN"}
	allRoles  = []string{"ADMIN", "USER"}
)

func PublicRoutes(
	userHandler handler.UserHandler,
	suratMasukHandler handler.SuratMasukHandler,
	suratKeluarHandler handler.SuratKeluarHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: userHandler.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},

		{
			Method:  http.MethodGet,
			Path:    "/suratKeluar",
			Handler: suratKeluarHandler.GetSuratKeluar,
		},
		{
			Method:  http.MethodGet,
			Path:    "/suratKeluar/:kode_surat",
			Handler: suratKeluarHandler.GetSuratKeluarByKodeSurat,
		},
		{
			Method:  http.MethodPost,
			Path:    "/suratKeluar",
			Handler: suratKeluarHandler.CreateSuratKeluar,
		},
		{
			Method:  http.MethodPut,
			Path:    "/suratKeluar/:kode_surat",
			Handler: suratKeluarHandler.UpdateSuratKeluar,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/suratKeluar/:kode_surat",
			Handler: suratKeluarHandler.DeleteSuratKeluar,
		},
		{
			Method:  http.MethodPost,
			Path:    "/suratMasuk",
			Handler: suratMasukHandler.CreateSuratMasuk,
		},
	}
}

func PrivateRoutes(
	userHandler handler.UserHandler,
	suratMasukHandler handler.SuratMasukHandler,
	suratKeluarHandler handler.SuratKeluarHandler) []route.Route {
	return []route.Route{

		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.GetUsers,
			Roles:   allRoles,
		},
		{

			Method:  http.MethodGet,
			Path:    "/users/:id",
			Handler: userHandler.GetUsers,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodGet,
			Path:    "/suratMasuk",
			Handler: suratMasukHandler.GetSuratMasuk,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodGet,
			Path:    "/suratMasuk/:kode_surat",
			Handler: suratMasukHandler.GetSuratMasukByKodeSurat,
			Roles:   allRoles,
		},

		{
			Method:  http.MethodPut,
			Path:    "/suratMasuk/:kode_surat",
			Handler: suratMasukHandler.UpdateSuratMasuk,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/suratMasuk/:kode_surat",
			Handler: suratMasukHandler.DeleteSuratMasuk,
			Roles:   allRoles,
		},
	}

}
