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

	project "go-just-portfolio/src/project"
	projecthttp "go-just-portfolio/src/project/delivery/http"
	projectmysql "go-just-portfolio/src/project/repository/mysql"
	projectusecase "go-just-portfolio/src/project/usecase"

	categories "go-just-portfolio/src/categories"
	categorieshttp "go-just-portfolio/src/categories/delivery/http"
	categoriesmysql "go-just-portfolio/src/categories/repository/mysql"
	categoriesusecase "go-just-portfolio/src/categories/usecase"
)

type App struct {
	httpServer *http.Server

	authUC       auth.UseCase
	projectUC    project.UseCase
	categoriesUC categories.UseCase
}

func NewApp() *App {
	db := database.InitDB()

	userRepo := authmysql.NewUserRepository(db)
	projectRepo := projectmysql.NewProjectRepository(db)
	categoriesRepo := categoriesmysql.NewСategoriesRepository(db)

	return &App{
		authUC:       authusecase.NewAuthUseCase(userRepo),
		projectUC:    projectusecase.NewprojectUseCase(projectRepo),
		categoriesUC: categoriesusecase.NewСategoriesUseCase(categoriesRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Static("/images", "./images")

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.Use(middleware.CORSMiddleware())

	authhttp.RegisterHTTPEndpoints(router, a.authUC)
	projecthttp.RegisterHTTPEndpoints(router, a.projectUC)
	categorieshttp.RegisterHTTPEndpoints(router, a.categoriesUC)

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
