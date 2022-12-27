package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/allrested/product/api-gateway/pkg/auth/routes"
	"github.com/allrested/product/api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/auth")
	routesGroup.POST("/register", svc.Register)
	routesGroup.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
