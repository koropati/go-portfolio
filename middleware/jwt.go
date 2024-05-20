package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/tokenutil"
	"github.com/koropati/go-portfolio/internal/urlutil"
)

const (
	UserIDContext   = "x-user-id"
	UserRoleContext = "x-user-role"
	RoleSuperAdmin  = "super_admin"
	RoleAdmin       = "admin"
	RoleStaff       = "staff"
	RoleAnonymous   = "anonymous"
	RefreshToken    = "refresh_token"
	AccessToken     = "access_token"
)

func JwtAuthMiddleware(secret string, casbinEnforcer *casbin.Enforcer, accessTokenUsecase domain.AccessTokenUsecase, refreshTokenUsecase domain.RefreshTokenUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		authToken, err := parseAuthorizationHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}

		if !accessTokenUsecase.IsValid(c, authToken) {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Token tidak valid atau telah kadaluarsa"})
			c.Abort()
			return
		}

		userID, userRole, err := tokenutil.ExtractIDFromToken(authToken, secret, AccessToken, accessTokenUsecase, refreshTokenUsecase)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}

		if userRole == "" {
			userRole = RoleAnonymous
		}

		setUserContext(c, userID, userRole)

		if err := enforceCasbinRules(c, casbinEnforcer, userRole); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Message})
			c.Abort()
			return
		}

		c.Next()
	}
}

func parseAuthorizationHeader(authHeader string) (string, error) {
	t := strings.Split(authHeader, " ")
	if len(t) != 2 {
		return "", errors.New("not authorized")
	}
	return t[1], nil
}

func setUserContext(c *gin.Context, userID, userRole string) {
	c.Set(UserIDContext, userID)
	c.Set(UserRoleContext, userRole)
}

func enforceCasbinRules(c *gin.Context, casbinEnforcer *casbin.Enforcer, userRole string) *domain.ErrorResponse {
	pathUrl := urlutil.RemoveAPIVersionMiddleware(c.Request.URL.Path)
	res, err := casbinEnforcer.EnforceSafe(userRole, pathUrl, c.Request.Method)
	if err != nil {
		return &domain.ErrorResponse{Message: err.Error()}
	}
	if !res {
		return &domain.ErrorResponse{Message: "unauthorized"}
	}
	return nil
}
