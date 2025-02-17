package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ElizaDeani/archivio/config"
	"github.com/ElizaDeani/archivio/internal/entity"
	"github.com/ElizaDeani/archivio/pkg/response"
	"github.com/ElizaDeani/archivio/pkg/route"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(cfg *config.Config, publicRoutes, privateRoutes []route.Route) *Server {
	e := echo.New()

	// ✅ Tambahkan Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://localhost:5500"}, // Bisa diganti dengan URL frontend spesifik
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1 := e.Group("/api/v1")
	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	if len(privateRoutes) > 0 {
		v1.Use(JWTMiddleware(cfg.JWTSecret))
		for _, route := range privateRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	return &Server{e}
}

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &entity.JWTCustomClaims{}
		},
		SigningKey:    []byte(secretKey),
		SigningMethod: "HS256",
		TokenLookup:   "header:Authorization,query:token",
		ErrorHandler: func(ctx echo.Context, err error) error {
			authHeader := ctx.Request().Header.Get("Authorization")

			// Hapus prefix "Bearer " jika ada
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Debugging log
			fmt.Println("🔴 Token yang diterima:", tokenString)
			fmt.Println("🔴 Error JWT:", err)

			// Cek jika token kosong atau error parsing
			if authHeader == "" {
				return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "token tidak ditemukan, silakan login"))
			}

			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "token tidak valid atau kadaluarsa"))
		},
	})
}
func RBACMiddleware(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*entity.JWTCustomClaims)

			allowed := false
			for _, role := range roles {
				if role == claims.Role {
					allowed = true
					break
				}
			}

			if !allowed {
				return c.JSON(http.StatusForbidden, response.ErrorResponse(http.StatusForbidden, "anda tidak memiliki akses untuk resource ini"))
			}
			return next(c)
		}
	}
}
