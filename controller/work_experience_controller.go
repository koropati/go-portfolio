package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/validator"
)

type WorkExperienceController struct {
	WorkExperienceUsecase domain.WorkExperienceUsecase
	AccessTokenUsecase    domain.AccessTokenUsecase
	RefreshTokenUsecase   domain.RefreshTokenUsecase
	Config                *bootstrap.Config
	Cryptos               cryptos.Cryptos
	Validator             *validator.Validator
}

func (ctr *WorkExperienceController) Create(c *gin.Context) {

	var request domain.WorkExperience

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = ctr.Validator.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = ctr.WorkExperienceUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Success Create Data",
		Success: true,
	})
}
