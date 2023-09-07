package main

import (
	"k8s-softroce-device-plugin/cmd"
	"k8s-softroce-device-plugin/utils"
)

func main() {
	utils.Must(cmd.Execute())
}
