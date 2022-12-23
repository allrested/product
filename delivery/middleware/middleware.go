package middleware

import (
	"github.com/allrested/product/utils/jwt"
	"github.com/allrested/product/utils/logger"
)

// Middleware ...
type Middleware struct {
	jwtSvc jwt.JWTService
	logger logger.Logger
}

// NewMiddleware will create new Middleware object
func NewMiddleware(jwtSvc jwt.JWTService, logger logger.Logger) *Middleware {
	return &Middleware{
		jwtSvc: jwtSvc,
		logger: logger,
	}
}
