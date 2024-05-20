package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
)

type RegisterController struct {
	UserUsecase domain.UserUsecase
	Config      *bootstrap.Config
	Cryptos     cryptos.Cryptos
}

func (rc *RegisterController) Register(c *gin.Context) {
	var request domain.RegisterUser

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userData, err := request.ToUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = rc.UserUsecase.Create(c, userData)
	if err == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Registration Successful",
	})
}
