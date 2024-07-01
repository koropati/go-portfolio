package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/casbin/casbin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/mailer"
	"github.com/koropati/go-portfolio/repository"
	"gopkg.in/robfig/cron.v2"
	"gorm.io/gorm"
)

type SetupConfig struct {
	Config         *bootstrap.Config
	Timeout        time.Duration
	DB             *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	Cryptos        cryptos.Cryptos
	Mailer         mailer.Mailer
}

func InitCron(config *SetupConfig) {
	stopChan := make(chan bool)

	sch := cron.New()

	_, _ = sch.AddFunc("* * * * *", func() {
		log.Print("Start Task TaskRemoveAccessToken()")
		TaskRemoveAccessToken(config)
	})
	_, _ = sch.AddFunc("* * * * *", func() {
		log.Print("Start Task TaskRemoveRefrehToken()")
		TaskRemoveRefrehToken(config)
	})
	_, _ = sch.AddFunc("* * * * *", func() {
		log.Print("Start Task TaskRemoveForgotPasswordToken()")
		TaskRemoveForgotPasswordToken(config)
	})

	sch.Start()
	<-stopChan

}

func TaskRemoveAccessToken(config *SetupConfig) {
	at := repository.NewAccessTokenRepository(config.DB, domain.AccessTokenTable, config.Config.DefaultPageNumber, config.Config.DefaultPageSize)
	err := at.DeleteExpiredToken(context.Background(), time.Now().UTC().UnixNano())
	if err != nil {
		log.Printf("Error Delete Expired Access Token: %v\n", err)
	}
}

func TaskRemoveRefrehToken(config *SetupConfig) {
	rt := repository.NewRefreshTokenRepository(config.DB, domain.RefreshTokenTable, config.Config.DefaultPageNumber, config.Config.DefaultPageSize)
	err := rt.DeleteExpiredToken(context.Background(), time.Now().UTC().UnixNano())
	if err != nil {
		log.Printf("Error Delete Expired Refresh Token: %v\n", err)
	}
}

func TaskRemoveForgotPasswordToken(config *SetupConfig) {
	rt := repository.NewForgotPasswordTokenRepository(config.DB, domain.ForgotPasswordTokenTable, config.Config.DefaultPageNumber, config.Config.DefaultPageSize)
	err := rt.DeleteExpiredToken(context.Background(), time.Now().UTC().UnixNano())
	if err != nil {
		log.Printf("Error Delete Expired Forgot Password Token: %v\n", err)
	}
}
