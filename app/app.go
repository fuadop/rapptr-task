package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fuadop/rapptr-task/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *gin.Engine
	Model  *model.Model
	DB     *sql.DB
}

func NewApp() *App {
	router := gin.Default() // prepare api framework
	return &App{Router: router}
}

func (a *App) Start() {
	a.Initialize()
	a.RegisterRoutes()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080" // default port
	}
	log.Fatalln(a.Router.Run(fmt.Sprint(":", port)))
}
