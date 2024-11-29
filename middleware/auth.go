package middleware

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/tokenutil"
)

const (
	LoginUrlRedirect     = "/login"
	DashboardUrlRedirect = "/dashboard"
)

func AuthMiddleware(secret string, casbinEnforcer *casbin.Enforcer, cryptos cryptos.Cryptos, accessTokenUsecase domain.AccessTokenUsecase, refreshTokenUsecase domain.RefreshTokenUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := GetAuthContext(c, cryptos, "access")
		if err != nil {
			c.Redirect(http.StatusFound, LoginUrlRedirect)
			return
		}

		if !accessTokenUsecase.IsValid(c, authToken) {
			accessTokenUsecase.Delete(c, authToken)
			refreshToken, err := GetAuthContext(c, cryptos, "refresh")
			if err != nil {
				refreshTokenUsecase.Delete(c, refreshToken)
			}
			c.Redirect(http.StatusFound, LoginUrlRedirect)
			return
		}

		userID, userRole, err := tokenutil.ExtractIDFromToken(authToken, secret, AccessToken, accessTokenUsecase, refreshTokenUsecase)
		if err != nil {
			c.Redirect(http.StatusFound, LoginUrlRedirect)
			return
		}

		if userRole == "" {
			userRole = RoleAnonymous
		}

		SetUserContext(c, cryptos, userID, userRole)

		if err := enforceCasbinRules(c, casbinEnforcer, userRole); err != nil {
			c.Redirect(http.StatusFound, LoginUrlRedirect)
			return
		}

		c.Next()
	}
}

func AuthPublicMiddleware(secret string, casbinEnforcer *casbin.Enforcer, cryptos cryptos.Cryptos, accessTokenUsecase domain.AccessTokenUsecase, refreshTokenUsecase domain.RefreshTokenUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, _ := GetAuthContext(c, cryptos, "access")
		if c.Request.URL.Path == "/" || c.Request.URL.Path == "/logout" {
			c.Next()
		}
		if authToken != "" {
			c.Redirect(http.StatusFound, DashboardUrlRedirect)
		} else {
			c.Next()
		}

	}
}
