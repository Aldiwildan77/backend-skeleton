package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skeleton",
	Short: "Main console runner.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(startCmd, []string{})
		os.Exit(0)
	},
}

// Runner execute the rootCmd
func Runner() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
