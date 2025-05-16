package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/RoseSecurity/terrafetch/internal"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Logo constant comes from logo.go.

// uiModel implements Bubble‑Tea's Model interface.
type uiModel struct {
	rootDir   string
	analytics []internal.Analytics
}

// ─── Styles ────────────────────────────────────────────────────────────────
var (
	borderColor = lipgloss.Color("99") // purple
	container   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(borderColor)

	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")) // magenta header
	keyStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))            // white
	valStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))           // light‑gray values
	logoStyle   = lipgloss.NewStyle().Foreground(borderColor)                     // purple logo
)

// NewUIModel returns a model that renders once and exits.
func NewUIModel(rootDir string, data []internal.Analytics) tea.Model {
	return uiModel{rootDir: rootDir, analytics: data}
}

// Init immediately quits after first frame.
func (m uiModel) Init() tea.Cmd { return tea.Quit }

// Update is a no‑op (renders once).
func (m uiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

// View assembles the bordered layout with analytics anchored bottom‑right.
func (m uiModel) View() string {
	logoBlock := buildPaddedLogo()                    // left column
	statsBlock := renderStats(m.rootDir, m.analytics) // right column

	// Align stats block to the bottom of the logo column and close border with trailing newline.
	body := lipgloss.JoinHorizontal(lipgloss.Top, logoBlock, statsBlock)
	return container.Render(body) + "\n"
}

// buildPaddedLogo pads each line of logo to equal width so right column aligns neatly.
func buildPaddedLogo() string {
	lines := strings.Split(strings.TrimSuffix(logo, "\n"), "\n")
	max := 0
	for _, l := range lines {
		if w := lipgloss.Width(l); w > max {
			max = w
		}
	}
	for i, l := range lines {
		pad := strings.Repeat(" ", max-lipgloss.Width(l))
		lines[i] = l + pad
	}
	return logoStyle.Render(strings.Join(lines, "\n"))
}

// renderStats lays out key/value stats; later aligned to bottom‑right by View().
func renderStats(rootDir string, data []internal.Analytics) string {
	if len(data) == 0 {
		return "No analytics found."
	}
	a := data[0]

	dir := filepath.Base(rootDir)

	padKey := func(k string) string { return keyStyle.Render(fmt.Sprintf("%-10s:", k)) }

	lines := []string{
		headerStyle.Render(dir),
		headerStyle.Render(strings.Repeat("-", len(dir))),
		fmt.Sprintf("%s %s", padKey("Variables"), valStyle.Render(fmt.Sprint(a.VariableCount))),
		fmt.Sprintf("%s %s", padKey("Resources"), valStyle.Render(fmt.Sprint(a.ResourceCount))),
		fmt.Sprintf("%s %s", padKey("Outputs"), valStyle.Render(fmt.Sprint(a.OutputCount))),
		fmt.Sprintf("%s %s", padKey("Modules"), valStyle.Render(fmt.Sprint(a.ModuleCount))),
		fmt.Sprintf("%s %s", padKey("Providers"), valStyle.Render(fmt.Sprint(a.ProviderCount))),
	}

	return lipgloss.NewStyle().PaddingLeft(2).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}
