package product

import (
	"github.com/gin-gonic/gin"
	"github.com/allrested/product/api-gateway/pkg/auth"
	"github.com/allrested/product/api-gateway/pkg/config"
	"github.com/allrested/product/api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/product")
	routesGroup.Use(a.AuthRequired)
	routesGroup.POST("/", svc.CreateProduct)
	routesGroup.GET("/:id", svc.FindOne)
	routesGroup.POST("/:id", svc.DecreaseStock)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) DecreaseStock(ctx *gin.Context) {
	routes.DecreaseStock(ctx, svc.Client)
}
