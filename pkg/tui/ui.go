package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/RoseSecurity/terrafetch/internal"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// NOTE: ASCII art logo is defined in logo.go as a public var `Logo`.

// uiModel implements Bubble-Tea's Model interface.
type uiModel struct {
	rootDir   string
	analytics []internal.Analytics
}

// ─── Styles ────────────────────────────────────────────────────────────────
var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")) // magenta header
	keyStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("111"))           // cyan keys
	valStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))           // light-gray values
	logoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))            // purple logo
)

// NewUIModel returns an initialised model ready to render once and quit.
func NewUIModel(rootDir string, data []internal.Analytics) tea.Model {
	return uiModel{rootDir: rootDir, analytics: data}
}

// Init immediately issues a Quit command so Bubble-Tea renders one frame then exits.
func (m uiModel) Init() tea.Cmd { return tea.Quit }

// Update is a no-op because we're quitting in Init.
func (m uiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

// View assembles the Neofetch-style output.
func (m uiModel) View() string {
	left := logoStyle.Render(logo)               // ASCII logo column (from logo.go)
	right := renderStats(m.rootDir, m.analytics) // Stats column

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

// renderStats lays out key/value stats to mimic Neofetch alignment.
func renderStats(rootDir string, data []internal.Analytics) string {
	if len(data) == 0 {
		return "No analytics found."
	}
	a := data[0]

	dirName := filepath.Base(rootDir)

	// Helper to pad keys so the colons align.
	padKey := func(k string) string {
		return keyStyle.Render(fmt.Sprintf("%-11s:", k)) // 11 left-padded so the longest key fits
	}

	// Compose lines.
	header := headerStyle.Render(dirName)
	underline := headerStyle.Render(strings.Repeat("-", len(dirName)))

	stats := []string{
		header,
		underline,
		fmt.Sprintf("%s %s", padKey("Variables"), valStyle.Render(fmt.Sprint(a.VariableCount))),
		fmt.Sprintf("%s %s", padKey("Resources"), valStyle.Render(fmt.Sprint(a.ResourceCount))),
		fmt.Sprintf("%s %s", padKey("Outputs"), valStyle.Render(fmt.Sprint(a.OutputCount))),
		fmt.Sprintf("%s %s", padKey("Modules"), valStyle.Render(fmt.Sprint(a.ModuleCount))),
		fmt.Sprintf("%s %s", padKey("Providers"), valStyle.Render(fmt.Sprint(a.ProviderCount))),
	}

	// Pad left so the stats sit a bit away from the logo.
	return lipgloss.NewStyle().PaddingLeft(2).Render(lipgloss.JoinVertical(lipgloss.Left, stats...))
}
