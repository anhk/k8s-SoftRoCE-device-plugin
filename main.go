package main

import (
	"k8s-softroce-device-plugin/app"
	"k8s-softroce-device-plugin/pkg/utils"
)

func main() {
	utils.Must(app.Execute())
}
