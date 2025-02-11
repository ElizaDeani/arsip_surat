package builder

import (
	"github.com/ElizaDeani/archivio/config"
	"github.com/ElizaDeani/archivio/internal/http/handler"
	"github.com/ElizaDeani/archivio/internal/http/router"
	"github.com/ElizaDeani/archivio/internal/repository"
	"github.com/ElizaDeani/archivio/internal/service"
	"github.com/ElizaDeani/archivio/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	userRepository := repository.NewUserRepository(db)
	suratMasukRepository := repository.NewSuratMasukRepository(db)
	suratKeluarRepository := repository.NewSuratKeluarRepository(db)

	//service
	userService := service.NewUserService(cfg, userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	suratMasukService := service.NewSuratMasukService(suratMasukRepository)
	suratKeluarService := service.NewSuratKeluarService(suratKeluarRepository)

	//handler
	userHandler := handler.NewUserHandler(tokenService, userService)
	suratMasukHandler := handler.NewSuratMasukHandler(suratMasukService)
	suratKeluarHandler := handler.NewSuratKeluarHandler(suratKeluarService)

	return router.PublicRoutes(userHandler, suratMasukHandler, suratKeluarHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	userRepository := repository.NewUserRepository(db)
	suratMasukRepository := repository.NewSuratMasukRepository(db)
	suratKeluarRepository := repository.NewSuratKeluarRepository(db)

	//service
	userService := service.NewUserService(cfg, userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	suratMasukService := service.NewSuratMasukService(suratMasukRepository)
	suratKeluarService := service.NewSuratKeluarService(suratKeluarRepository)

	//handler
	userHandler := handler.NewUserHandler(tokenService, userService)
	suratMasukHandler := handler.NewSuratMasukHandler(suratMasukService)
	suratKeluarHandler := handler.NewSuratKeluarHandler(suratKeluarService)

	return router.PrivateRoutes(userHandler, suratMasukHandler, suratKeluarHandler)
}
