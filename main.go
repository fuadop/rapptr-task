package main

import (
	"github.com/fuadop/rapptr-task/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load() // load envs
}

func main() {
	a := app.NewApp()
	a.Start()
}
