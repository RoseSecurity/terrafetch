package cmd

import (
	"github.com/RoseSecurity/terrafetch/internal"
	"github.com/RoseSecurity/terrafetch/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
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

	p := tea.NewProgram(tui.NewUIModel(directory, analytics))
	_, err = p.Run()
	if err != nil {
		log.Error("failed to start TUI", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}
