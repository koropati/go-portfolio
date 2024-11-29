package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/controller"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/repository"
	"github.com/koropati/go-portfolio/usecase"
)

func NewWorkExperienceRouter(cfg *SetupConfig, group *gin.RouterGroup) {
	repo := repository.NewWorkExperienceRepository(cfg.DB, domain.WorkExperienceTable, cfg.Config.DefaultPageNumber, cfg.Config.DefaultPageSize)
	ctrl := controller.WorkExperienceController{
		WorkExperienceUsecase: usecase.NewWorkExperienceUsecase(repo, cfg.Timeout),
		Config:                cfg.Config,
		Cryptos:               cfg.Cryptos,
		Validator:             cfg.Validator,
	}

	group.POST("/work-experience", ctrl.Create)
}
