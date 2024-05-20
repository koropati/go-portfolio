package routes

import (
	"time"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"gorm.io/gorm"
)

const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleStaff      = "staff"
	DeviceSwitch   = "switch"
	DeviceHps      = "hps"
	DeviceIr       = "ir"
	DeviceAc       = "ac"
)

type SetupConfig struct {
	Config         *bootstrap.Config
	Timeout        time.Duration
	DB             *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	Cryptos        cryptos.Cryptos
	Gin            *gin.Engine
}

func Setup(config *SetupConfig) {
	config.Gin.Static("v1/assets", "./templates/assets")
	config.Gin.LoadHTMLGlob("./templates/html/*.html")

	// All Public APIs
	publicRouterV1 := config.Gin.Group("/v1")

	NewRegisterRouter(config, publicRouterV1)

}
