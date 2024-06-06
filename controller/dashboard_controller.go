package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/validator"
)

type DashboardPageController struct {
	Config    *bootstrap.Config
	Cryptos   cryptos.Cryptos
	Validator *validator.Validator
}

func (rc *DashboardPageController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.tmpl", nil)
}
