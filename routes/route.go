package routes

import (
	"time"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/validator"
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
	Validator      *validator.Validator
}

func Setup(config *SetupConfig) {

	config.Gin.Static("assets", "./templates/assets")
	config.Gin.LoadHTMLGlob("./templates/*.tmpl")
	// All Public APIs
	publicRouter := config.Gin.Group("/")
	NewLandingPageRouter(config, publicRouter)
	NewRegisterRouter(config, publicRouter)

}
