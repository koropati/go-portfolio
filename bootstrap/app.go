package bootstrap

import (
	"github.com/casbin/casbin"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/validator"
	"gorm.io/gorm"
)

type Application struct {
	Config         *Config
	DB             *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	Cryptos        cryptos.Cryptos
	Validator      *validator.Validator
}

type AppFunc func(*Application)

func defaultApp() Application {
	myConfig := NewConfig()
	return Application{
		Config:         myConfig,
		DB:             NewDatabase(myConfig),
		CasbinEnforcer: NewCasbinEnforcer(myConfig),
		Cryptos:        NewCryptos(myConfig),
		Validator:      NewValidator(),
	}
}

func NewApp(opts ...AppFunc) *Application {
	app := defaultApp()
	for _, fn := range opts {
		fn(&app)
	}
	return &app
}

func App() Application {
	app := &Application{}
	app.Config = NewConfig()
	app.DB = NewDatabase(app.Config)
	app.CasbinEnforcer = NewCasbinEnforcer(app.Config)
	app.Cryptos = NewCryptos(app.Config)
	app.Validator = NewValidator()

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDatabase(app.DB)
}
