package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	database "go-just-portfolio/pkg/database"
	middleware "go-just-portfolio/pkg/middleware"

	gin "github.com/gin-gonic/gin"

	auth "go-just-portfolio/src/auth"
	authhttp "go-just-portfolio/src/auth/delivery/http"
	authmysql "go-just-portfolio/src/auth/repository/mysql"
	authusecase "go-just-portfolio/src/auth/usecase"
	// service "go-just-portfolio/service"
	// servicehttp "go-just-portfolio/service/delivery/http"
	// servicemysql "go-just-portfolio/service/repository/mysql"
	// serviceusecase "go-just-portfolio/service/usecase"
	// market "go-just-portfolio/market"
	// markethttp "go-just-portfolio/market/delivery/http"
	// marketmysql "go-just-portfolio/market/repository/mysql"
	// marketusecase "go-just-portfolio/market/usecase"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase
	// serviceUC service.UseCase
	// marketUC  market.UseCase
}

func NewApp() *App {
	db := database.InitDB()

	userRepo := authmysql.NewUserRepository(db)
	// serviceRepo := servicemysql.NewUserRepository(db)
	// marketRepo := marketmysql.NewUserRepository(db)

	return &App{
		authUC: authusecase.NewAuthUseCase(userRepo),
		// marketUC:  marketusecase.NewAuthUseCase(marketRepo),
		// serviceUC: serviceusecase.NewAuthUseCase(serviceRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.Use(middleware.CORSMiddleware())

	authhttp.RegisterHTTPEndpoints(router, a.authUC)
	// servicehttp.RegisterHTTPEndpoints(router, a.serviceUC)
	// markethttp.RegisterHTTPEndpoints(router, a.marketUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
