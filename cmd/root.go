package cmd

import (
	"os"

	"github.com/RoseSecurity/terrafetch/internal"
	"github.com/RoseSecurity/terrafetch/pkg/tui"
	log "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var directory string

var rootCmd = &cobra.Command{
	Use:          "terrafetch",
	Short:        "The Neofetch of Terraform",
	Long:         `Turning infrastructure repository statistics into stylish terminal outputs since 2025.`,
	RunE:         fetchInfo,
	SilenceUsage: true,
}

func init() {
	// ‑‑directory / -d flag to point at Terraform code (defaults to current dir)
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory containing the Terraform configurations")
	rootCmd.DisableAutoGenTag = true // keep generated docs clean
}

// fetchInfo gathers repository analytics and hands them to the TUI.
// It bubbles any failure up to Cobra so the CLI exits with non‑zero status.
func fetchInfo(cmd *cobra.Command, args []string) error {
	analytics, err := internal.AnalyzeRepository(directory)
	if err != nil {
		log.Error(internal.ErrFailedToFetch, err)
		return err
	}

	if len(analytics) == 0 {
		return internal.ErrNoTerraformFiles
	}

	tui.DisplayInfo(directory, analytics[0])
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
