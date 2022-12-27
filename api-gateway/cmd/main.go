package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/allrested/product/api-gateway/pkg/auth"
	"github.com/allrested/product/api-gateway/pkg/config"
	"github.com/allrested/product/api-gateway/pkg/order"
	"github.com/allrested/product/api-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://127.0.0.1:5533"},
	//	AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "http://127.0.0.1:5533"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	err = r.Run(c.Port)
	if err != nil {
		log.Println("api gateway run failed!")
	}
}
