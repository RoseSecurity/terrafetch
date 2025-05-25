package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/RoseSecurity/terrafetch/internal"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

var (
	borderColor = lipgloss.Color("99") // purple
	container   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(borderColor)

	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")) // magenta header
	keyBase     = lipgloss.NewStyle().Foreground(lipgloss.Color("111"))           // cyan keys (base style)
	valStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))            // white values
	logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))            // purple logo
)

// padKey dynamically pads and styles a key string based on the provided style
// (whose width is already set to the widest key + 1 for the trailing colon).
func padKey(k string, style lipgloss.Style) string {
	return style.Render(k + ":")
}

func RenderInfo(dir string, a internal.Analytics) string {
	// All label strings in the order they'll be rendered.
	labels := []string{
		"Terraform Files",
		"Documentation",
		"Providers",
		"Module Calls",
		"Resources",
		"Data Sources",
		"Variables",
		"Sensitive Variables",
		"Outputs",
		"Sensitive Outputs",
	}

	// Find the widest label using runewidth (handles double‑width runes correctly).
	max := 0
	for _, l := range labels {
		if w := runewidth.StringWidth(l); w > max {
			max = w
		}
	}

	// Clone the base key style and set its width dynamically.
	keyStyle := keyBase.Width(max + 1) // +1 for the trailing colon

	tfDir := filepath.Base(dir)

	// Helper closure so we don't repeat ourselves.
	pk := func(k string) string { return padKey(k, keyStyle) }

	lines := []string{
		headerStyle.Render(tfDir),
		headerStyle.Render(strings.Repeat("-", runewidth.StringWidth(tfDir))),
		fmt.Sprintf("%s %s", pk("Terraform Files"), valStyle.Render(fmt.Sprint(a.FileCount))),
		fmt.Sprintf("%s %s", pk("Documentation"), valStyle.Render(fmt.Sprint(a.DocCount))),
		fmt.Sprintf("%s %s", pk("Providers"), valStyle.Render(fmt.Sprint(a.ProviderCount))),
		fmt.Sprintf("%s %s", pk("Module Calls"), valStyle.Render(fmt.Sprint(a.ModuleCount))),
		fmt.Sprintf("%s %s", pk("Resources"), valStyle.Render(fmt.Sprint(a.ResourceCount))),
		fmt.Sprintf("%s %s", pk("Data Sources"), valStyle.Render(fmt.Sprint(a.DataSourceCount))),
		fmt.Sprintf("%s %s", pk("Variables"), valStyle.Render(fmt.Sprint(a.VariableCount))),
		fmt.Sprintf("%s %s", pk("Sensitive Variables"), valStyle.Render(fmt.Sprint(a.SensitiveVariableCount))),
		fmt.Sprintf("%s %s", pk("Outputs"), valStyle.Render(fmt.Sprint(a.OutputCount))),
		fmt.Sprintf("%s %s", pk("Sensitive Outputs"), valStyle.Render(fmt.Sprint(a.SensitiveOutputCount))),
	}

	rightColumn := lipgloss.NewStyle().
		PaddingLeft(4).
		PaddingRight(4).
		Render(strings.Join(lines, "\n"))

	logoColumn := logoStyle.Render(logo)

	return container.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, logoColumn, rightColumn),
	)
}

func DisplayInfo(dir string, a internal.Analytics) {
	fmt.Println(RenderInfo(dir, a))
}
