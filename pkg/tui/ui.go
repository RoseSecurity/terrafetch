package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/RoseSecurity/terrafetch/internal"
	"github.com/charmbracelet/lipgloss"
)

var (
	borderColor = lipgloss.Color("99") // purple
	container   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(borderColor)

	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")) // magenta header
	keyStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("111"))           // cyan keys
	valStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))            // white values
	logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))            // purple logo
)

func padKey(k string) string {
	return keyStyle.Render(fmt.Sprintf("%-10s:", k))
}

func RenderInfo(dir string, a internal.Analytics) string {
	tfDir := filepath.Base(dir)
	lines := []string{
		headerStyle.Render(tfDir),
		headerStyle.Render(strings.Repeat("-", len(tfDir))),
		fmt.Sprintf("%s %s", padKey("Terraform Files"), valStyle.Render(fmt.Sprint(a.FileCount))),
		fmt.Sprintf("%s %s", padKey("Docs"), valStyle.Render(fmt.Sprint(a.DocCount))),
		fmt.Sprintf("%s %s", padKey("Resources"), valStyle.Render(fmt.Sprint(a.ResourceCount))),
		fmt.Sprintf("%s %s", padKey("Modules"), valStyle.Render(fmt.Sprint(a.ModuleCount))),
		fmt.Sprintf("%s %s", padKey("Variables"), valStyle.Render(fmt.Sprint(a.VariableCount))),
		fmt.Sprintf("%s %s", padKey("Outputs"), valStyle.Render(fmt.Sprint(a.OutputCount))),
		fmt.Sprintf("%s %s", padKey("Providers"), valStyle.Render(fmt.Sprint(a.ProviderCount))),
	}

	rightColumn := lipgloss.NewStyle().
		PaddingLeft(4).
		Render(strings.Join(lines, "\n"))
	logoColumn := logoStyle.Render(logo)

	return container.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, logoColumn, rightColumn),
	)
}

func DisplayInfo(dir string, a internal.Analytics) {
	fmt.Println(RenderInfo(dir, a))
}
