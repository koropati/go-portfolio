package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/controller"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/repository"
	"github.com/koropati/go-portfolio/usecase"
)

func NewLoginRouter(cfg *SetupConfig, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(cfg.DB, domain.UserTable, cfg.Config.DefaultPageNumber, cfg.Config.DefaultPageSize)
	at := repository.NewAccessTokenRepository(cfg.DB, domain.AccessTokenTable, cfg.Config.DefaultPageNumber, cfg.Config.DefaultPageSize)
	rt := repository.NewRefreshTokenRepository(cfg.DB, domain.RefreshTokenTable, cfg.Config.DefaultPageNumber, cfg.Config.DefaultPageSize)
	lc := controller.LoginController{
		UserUsecase:         usecase.NewUserUsecase(ur, cfg.Timeout),
		AccessTokenUsecase:  usecase.NewAccessTokenUsecase(at, cfg.Timeout),
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(rt, cfg.Timeout),
		Config:              cfg.Config,
		Cryptos:             cfg.Cryptos,
		Validator:           cfg.Validator,
	}

	group.GET("/login", lc.Index)
	group.POST("/login", lc.Login)
}
