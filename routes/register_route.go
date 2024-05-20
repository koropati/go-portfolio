package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/controller"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/repository"
	"github.com/koropati/go-portfolio/usecase"
)

func NewRegisterRouter(cfg *SetupConfig, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(cfg.DB, domain.UserTable, cfg.Config.DefaultPageNumber, cfg.Config.DefaultPageSize)
	sc := controller.RegisterController{
		UserUsecase: usecase.NewUserUsecase(ur, cfg.Timeout),
		Config:      cfg.Config,
		Cryptos:     cfg.Cryptos,
	}
	group.POST("/register", sc.Register)
}
