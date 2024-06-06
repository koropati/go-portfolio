package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/tokenutil"
	"github.com/koropati/go-portfolio/internal/validator"
	"github.com/koropati/go-portfolio/middleware"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	UserUsecase         domain.UserUsecase
	AccessTokenUsecase  domain.AccessTokenUsecase
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Config              *bootstrap.Config
	Cryptos             cryptos.Cryptos
	Validator           *validator.Validator
}

func (lc *LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginUser

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = lc.Validator.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	user, err := lc.UserUsecase.GetByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: "Wrong email or password", Success: false})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: "Wrong email or password", Success: false})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: "User is not active", Success: false})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(&user, lc.Config.AccessTokenSecret, lc.Config.AccessTokenExpiryHour, lc.AccessTokenUsecase)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	refreshToken, err := tokenutil.CreateRefreshToken(&user, lc.Config.RefreshTokenSecret, lc.Config.RefreshTokenExpiryHour, accessToken, lc.RefreshTokenUsecase)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	middleware.SetAuthCookies(accessToken, refreshToken)

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Login Successful",
		Success: true,
		Data: domain.UserTokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
