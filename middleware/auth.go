package middleware

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/tokenutil"
)

func AuthMiddleware(secret string, casbinEnforcer *casbin.Enforcer, cryptos cryptos.Cryptos, accessTokenUsecase domain.AccessTokenUsecase, refreshTokenUsecase domain.RefreshTokenUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := parseAuthorizationCookies(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.JsonResponse{Message: err.Error(), Success: false})
			c.Abort()
			return
		}

		if !accessTokenUsecase.IsValid(c, authToken) {
			c.JSON(http.StatusUnauthorized, domain.JsonResponse{Message: "Token tidak valid atau telah kadaluarsa", Success: false})
			c.Abort()
			return
		}

		userID, userRole, err := tokenutil.ExtractIDFromToken(authToken, secret, AccessToken, accessTokenUsecase, refreshTokenUsecase)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.JsonResponse{Message: err.Error(), Success: false})
			c.Abort()
			return
		}

		if userRole == "" {
			userRole = RoleAnonymous
		}

		SetUserContext(c, cryptos, userID, userRole)

		if err := enforceCasbinRules(c, casbinEnforcer, userRole); err != nil {
			c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Message, Success: false})
			c.Abort()
			return
		}

		c.Next()
	}
}

func parseAuthorizationCookies(c *gin.Context) (string, error) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found"})
		return "", err
	}

	return accessToken, nil
}
