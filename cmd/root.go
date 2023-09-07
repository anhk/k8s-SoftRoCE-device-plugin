package cmd

import (
	"k8s-softroce-device-plugin/app"
	"k8s-softroce-device-plugin/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "k8s-softroce-device-plugin",
	Long: "kubernetes device plugin for SoftRoce",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApp()
		utils.Must(app.Run())
	},
}

func Execute() error {
	return rootCmd.Execute()
}
