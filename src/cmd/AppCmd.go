package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func ExecuteAppCmd() {
	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Application Description",
	}
	appCmd.AddCommand(CreateServeCmd())

	if err := appCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
