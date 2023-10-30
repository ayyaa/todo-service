package main

import (
	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/handler"
	awsSees "github.com/ayyaa/todo-services/lib/helper"
	utils "github.com/ayyaa/todo-services/lib/validator"
	"github.com/ayyaa/todo-services/repository"
	"github.com/ayyaa/todo-services/services"
	"github.com/labstack/echo/v4"
)

// //go:embed migrations
// var migrations embed.FS

func main() {

	cfg := config.LoadConfig()

	e := echo.New()

	sess, err := awsSees.NewSession(cfg.AWS)
	if err != nil {
		cfg.AWS.AwsFlag = false
	}

	cfg.AWS.AWSSession = sess

	// Initialize the server
	s := newServer(cfg)

	// Group route
	l := e.Group("/lists") // Base path for the group

	// Routes under the group
	l.GET("/:id", s.GetListByID)
	l.POST("", s.CreateList)
	l.DELETE("/:id", s.DeleteList)
	l.PUT("/:id", s.EditList)
	l.GET("", s.GetLists)

	// Group route
	sl := e.Group("/sublists") // Base path for the group

	// Routes under the group
	sl.GET("/:id", s.GetSubListByID)
	sl.POST("", s.CreateSubList)
	sl.DELETE("/:id", s.DeleteSubList)
	sl.PUT("/:id", s.EditSubList)
	sl.GET("", s.GetSubLists)

	e.Logger.Fatal(e.Start(":5656"))
}

func newServer(cfg config.Config) *handler.Server {
	utils.InitValidator()

	// initialize repository layer
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Cfg: cfg,
	})

	// Initialize service layer
	var srvc services.ServiceInterface = services.NewService(services.NewServiceOptions{
		Repository: repo,
		Cfg:        cfg,
	})

	opts := handler.NewServerOptions{
		Repository: repo,
		Service:    srvc,
		Cfg:        cfg,
	}

	return handler.NewServer(opts)
}
