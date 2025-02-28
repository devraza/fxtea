package main

import (
	"fmt"
	"time"
  "strings"
  "strconv"

	"fxtea/fx"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/textinput"
)

const (
	dotChar           = " • "
)

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	keyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

func main() {
  ti := textinput.New()
  ti.Focus()
  ti.Width = 20

	initialModel := model{0, false, 10, false, ti}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type (
	tickMsg  struct{}
	frameMsg struct{}
)

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

type model struct {
	Choice   int
	Chosen   bool
	Frames   int
	Quitting bool
	TextInput textinput.Model
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	if !m.Chosen {
		return updateChoices(msg, m)
	}

  m.TextInput, cmd = m.TextInput.Update(msg)
  return m, cmd
}

func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n  Quitting...\n"
	}
	if !m.Chosen {
		s = choicesView(m)
	} else {
	  switch m.Choice {
	    case 0:
        s = quadraticView(m)
	    default:
		    s = quadraticView(m)
    }
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.Choice < 2 {
			  m.Choice++
			}
		case "k", "up":
			if m.Choice > 0 {
			  m.Choice--
			}
		case "enter":
			m.Chosen = true
			return m, frame()
		}
	}

	return m, nil
}

func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		m.Frames++
		return m, frame()
  }


	return m, nil
}

func choicesView(m model) string {
	c := m.Choice

	tpl := "Run what?\n\n"
	tpl += "%s\n\n"

  splits := []string{help("↑ /k", "up"), help("↓ /j", "down"), help("q", "quit")}

	tpl += strings.Join(splits, dotStyle)

  choices := fmt.Sprintf(
		"%s\n%s\n%s",
		checkbox("Quadratic", c == 0),
		checkbox("Poisson", c == 1),
		checkbox("Chai", c == 2),
	)

	return fmt.Sprintf(tpl, choices)
}

func quadraticView (m model) string {
	arguments := strings.Split(m.TextInput.Value(), " ")
  m.TextInput.Placeholder = "a, b, c"

  var result string

  if len(arguments) == 3 {
    var coefficients []float64

    for i := range arguments {
      parsed, _ := strconv.ParseFloat(arguments[i], 64)
      coefficients = append(coefficients, parsed)
    }

    roots := fx.Quadratic(coefficients[0], coefficients[1], coefficients[2])
    result = fmt.Sprintf(
      "The roots are %v and %v",
      keywordStyle.Render(fx.FormatFloat(roots[0])),
      keywordStyle.Render(fx.FormatFloat(roots[1])),
    )
  }

  return fmt.Sprintf(
    "%s\n\n%v",
    m.TextInput.View(),
    result,
  )
}

func help(key string, label string) string {
  return keyStyle.Render(key) + " " + subtleStyle.Render(label)
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("• " + label)
	}
	return fmt.Sprintf("  %s", label)
}
