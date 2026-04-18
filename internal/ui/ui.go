package ui

import (
	"strings"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/balaji01-4d/cake/internal/parser"
	"github.com/muesli/termenv"
)

var (
	codeLexer    = "Makefile"
	codeStyle    = "monokai"
	colorProfile = detectTerminalColorProfile()
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2).Italic(true)

	previewBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2)

	dialog = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)
)

type Model struct {
	list          list.Model
	CurrentTarget *parser.MakeTarget
	Input         textinput.Model
	editCmdDialog bool
	FinalCmd      string
	width, height int
}

func New(items []*parser.MakeTarget) tea.Model {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}
	i := textinput.New()
	i.Placeholder = "Edit command..."
	i.CharLimit = 256

	l := list.New(
		listItems, list.NewDefaultDelegate(), 0, 0,
	)
	l.Title = "Targets"
	l.SetFilteringEnabled(false)
	m := Model{
		list:  l,
		Input: i,
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) selectedTarget() *parser.MakeTarget {
	item := m.list.SelectedItem()
	target, ok := item.(*parser.MakeTarget)
	if !ok || target == nil {
		return nil
	}
	return target
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else if msg.String() == "enter" {
			if m.editCmdDialog {
				m.FinalCmd = m.Input.Value()
				return m, tea.Quit
			}
			if target := m.selectedTarget(); target != nil {
				m.CurrentTarget = target
				return m, tea.Quit
			}
		} else if msg.String() == "shift+enter" {
			if m.editCmdDialog {
				return m, nil
			}
			target := m.selectedTarget()
			if target == nil {
				return m, nil
			}
			m.CurrentTarget = target
			m.Input.SetValue("make " + target.Name)
			m.editCmdDialog = true
			m.Input.Focus()
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := docStyle.GetFrameSize()
		leftWidth := (msg.Width - h) / 2
		leftHight := msg.Height - v
		m.list.SetSize(leftWidth, leftHight)
	}

	var cmd tea.Cmd
	if m.editCmdDialog {
		m.Input, cmd = m.Input.Update(msg)
		return m, cmd
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("Initializing...")
	}

	if m.editCmdDialog {
		d := dialog.Render("Edit Command:\n\n" + m.Input.View())
		da := lipgloss.Place(
			m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			d,
		)
		v := tea.NewView(da)
		v.AltScreen = true
		return v
	}

	leftPane := m.list.View()

	var rightPane string

	// Calculate how much space is left for the right pane
	horizontalMargin, verticalMargin := docStyle.GetFrameSize()
	rightPaneWidth := m.width - horizontalMargin - lipgloss.Width(leftPane) - previewBoxStyle.GetHorizontalFrameSize()

	// Calculate height to match the list exactly
	paneHeight := m.height - verticalMargin - previewBoxStyle.GetVerticalFrameSize()

	if rightPaneWidth > 0 && paneHeight > 0 {
		item := m.list.SelectedItem()

		if target, ok := item.(*parser.MakeTarget); ok && target != nil {
			snippet := target.String()
			rightPane = previewBoxStyle.Width(rightPaneWidth).Height(paneHeight).Render(highlightcode(snippet))
		} else {
			rightPane = previewBoxStyle.Width(rightPaneWidth).Height(paneHeight).Render("")
		}
	}

	splitLayout := lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)

	v := tea.NewView(docStyle.Render(splitLayout))
	v.AltScreen = true
	return v
}

func highlightcode(code string) string {
	var s strings.Builder

	err := quick.Highlight(&s, code, codeLexer, colorProfile, codeStyle)
	if err != nil {
		return code // Fallback to unformatted code on error
	}
	return s.String()
}

func detectTerminalColorProfile() string {
	switch termenv.ColorProfile() {
	case termenv.TrueColor:
		return "terminal16m"
	case termenv.ANSI256:
		return "terminal256"
	case termenv.ANSI:
		return "terminal16"
	default:
		return "noop" // Chroma's no-op formatter
	}
}
