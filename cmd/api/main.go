package main

import (
	"github.com/allrested/product/utils/crypto"
	"github.com/allrested/product/utils/jwt"
	"net/http"
	"time"

	_ "github.com/allrested/product/docs"
	"github.com/allrested/product/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/allrested/product/config"
	httpDelivery "github.com/allrested/product/delivery/http"
	appMiddleware "github.com/allrested/product/delivery/middleware"
	"github.com/allrested/product/infrastructure/datastore"
	pgsqlRepository "github.com/allrested/product/repository/pgsql"
	redisRepository "github.com/allrested/product/repository/redis"
	"github.com/allrested/product/usecase"
	"github.com/allrested/product/utils/logger"
)

// @title Product Api Gateway
// @version 1.0.4
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey JwtToken
// @in header
// @name Authorization
func main() {
	// Load config
	configApp := config.LoadConfig()

	// Setup logger
	appLogger := logger.NewApiLogger(configApp)
	appLogger.InitLogger()

	// Setup infra
	dbInstance, err := datastore.NewDatabase(configApp.DatabaseURL)
	utils.PanicIfNeeded(err)

	cacheInstance, err := datastore.NewCache(configApp.CacheURL)
	utils.PanicIfNeeded(err)

	// Setup repository
	redisRepo := redisRepository.NewRedisRepository(cacheInstance)
	todoRepo := pgsqlRepository.NewPgsqlTodoRepository(dbInstance)
	userRepo := pgsqlRepository.NewPgsqlUserRepository(dbInstance)

	// Setup Service
	cryptoSvc := crypto.NewCryptoService()
	jwtSvc := jwt.NewJWTService(configApp.JWTSecretKey)

	// Setup usecase
	ctxTimeout := time.Duration(configApp.ContextTimeout) * time.Second
	todoUC := usecase.NewTodoUsecase(todoRepo, redisRepo, ctxTimeout)
	authUC := usecase.NewAuthUsecase(userRepo, cryptoSvc, jwtSvc, ctxTimeout)

	// Setup app middleware
	appMiddleware := appMiddleware.NewMiddleware(jwtSvc, appLogger)

	// Setup route engine & middleware
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(appMiddleware.RequestID())
	e.Use(appMiddleware.Logger())
	e.Use(middleware.Recover())

	// Setup handler
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "i am alive")
	})

	httpDelivery.NewTodoHandler(e, appMiddleware, todoUC)
	httpDelivery.NewAuthHandler(e, appMiddleware, authUC)

	e.Logger.Fatal(e.Start(":" + configApp.ServerPORT))
}
