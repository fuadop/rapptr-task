package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/fuadop/rapptr-task/docs"
)

//	@title			Todo List API
//	@version		1.0
//	@description	Todo list API

//	@contact.name	Fuad Olatunji
//	@contact.url	https://nodelayer.xyz
//	@contact.email	fuad@nodelayer.xyz

// @BasePath					/api/v1
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func (a *App) RegisterRoutes() {
	// Prepare docs
	host := "localhost"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080" // default port
	}
	if _host, err := os.Hostname(); err == nil {
		if strings.Contains(_host, "local") {
			host = "localhost"	
		} else if strings.Contains(_host, ".") { // check if it's a domain name
			host = _host
		}
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	a.Router.Use(cors.Default()) // allow all origins

	router := a.Router.Group("/api/v1")

	todos := router.Group("/todos")
	todos.Use(a.MustAuth())

	todos.GET("/", a.GetTodos)         // get all user todos
	todos.GET("/:id", a.GetTodo)       // get a todo
	todos.POST("/", a.AddTodo)         // add a todo
	todos.PATCH("/:id", a.EditTodo)    // edit a todo
	todos.DELETE("/:id", a.DeleteTodo) // delete a todo

	auth := router.Group("/auth")

	auth.GET("/me", a.MustAuth(), a.GetMe) // get current user - based on auth header

	auth.POST("/login", a.Login)       // login a user - get jwt
	auth.POST("/register", a.Register) // register a user

	// Documentation routes
	a.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
