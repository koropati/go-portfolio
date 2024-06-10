package controller

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/validator"
)

type LogoutController struct {
	UserUsecase         domain.UserUsecase
	AccessTokenUsecase  domain.AccessTokenUsecase
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Config              *bootstrap.Config
	Cryptos             cryptos.Cryptos
	Validator           *validator.Validator
}

func (lc *LogoutController) Logout(c *gin.Context) {
	session := sessions.Default(c) // Get the current session

	// Clear the session data
	session.Clear()

	// Save the changes to the session store
	err := session.Save()
	if err != nil {
		// Handle error saving the session
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// Redirect to the login page or other appropriate location
	c.Redirect(http.StatusFound, "/login")
}