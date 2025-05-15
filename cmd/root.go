package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var directory string

var rootCmd = &cobra.Command{
	Use:   "terrafetch",
	Short: "The Neofetch of Terraform",
	Long:  `Turning infrastructure repository statistics into stylish terminal outputs sinc 2025.`,
	Run:   fetchInfo,
}

func init() {
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory of the Terraform configurations")
	// Disable auto generated string from documentation so that documentation is cleanly built and updated
	rootCmd.DisableAutoGenTag = true
}

func fetchInfo(cmd *cobra.Command, args []string) {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}
