package app

import (
	"k8s-softroce-device-plugin/pkg/log"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() error {
	log.Infof("Kubernetes Device Plugin for SoftRoce start")
	return nil
}
