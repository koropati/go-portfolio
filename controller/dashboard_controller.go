package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/validator"
)

type DashboardController struct {
	Config    *bootstrap.Config
	Cryptos   cryptos.Cryptos
	Validator *validator.Validator
}

func (ctr *DashboardController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.tmpl", nil)
}
