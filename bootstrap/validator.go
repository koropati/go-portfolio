package bootstrap

import (
	"sync"

	"github.com/koropati/go-portfolio/internal/validator"
)

var (
	myValidatorOnce sync.Once
	myValidator     *validator.Validator
)

// NewRedisClient creates a new Redis client connection
func NewValidator() *validator.Validator {
	myValidatorOnce.Do(func() {
		myValidator = validator.NewValidator()
	})

	return myValidator
}
