package app

import (
	"github.com/fuadop/rapptr-task/model"
)

func (a *App) Initialize() {
	a.DB = ConnectDatabase()
	a.Model = model.NewModel(a.DB)

	a.AutoMigrate()
}
