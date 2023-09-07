package app

import (
	"k8s-softroce-device-plugin/pkg/log"
)

type App struct {
	devicePlugin *SoftRoceDevicePlugin
}

func NewApp() *App {
	dp := NewSoftRoceDevicePlugin()
	return &App{devicePlugin: dp}
}

func (app *App) Run() error {
	log.Infof("Kubernetes Device Plugin for SoftRoce start")
	app.devicePlugin.Start()

	select {}
}
