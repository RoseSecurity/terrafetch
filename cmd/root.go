package cmd

import (
	"fmt"

	"github.com/RoseSecurity/terrafetch/internal"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var directory string

var rootCmd = &cobra.Command{
	Use:   "terrafetch",
	Short: "The Neofetch of Terraform",
	Long:  `Turning infrastructure repository statistics into stylish terminal outputs since 2025.`,
	Run:   fetchInfo,
}

func init() {
	// Optional input directory for location of Terraform code.
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory of the Terraform configurations")
	// Disable auto generated string from documentation so that documentation is cleanly built and updated.
	rootCmd.DisableAutoGenTag = true
}

func fetchInfo(cmd *cobra.Command, args []string) {
	analytics, err := internal.AnalyzeRepository(directory)
	if err != nil {
		log.Error(internal.ErrFailedToFetch, err)
	}

	for _, a := range analytics {
		fmt.Printf("Variables: %d\n", a.VariableCount)
		fmt.Printf("Resources: %d\n", a.ResourceCount)
		fmt.Printf("Outputs: %d\n", a.OutputCount)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}
