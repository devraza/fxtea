package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	dotChar = " • "
)

// Choices
const (
	choiceQuadratic    = iota
	choicePoisson      = iota
	choiceChi          = iota
	choiceBinarySearch = iota
	choiceFibonacci    = iota
	choiceLen          = iota // prevent the menu scrolling past this point
)

func main() {
	ti := textinput.New()
	ti.Width = 20

	initialModel := model{choiceQuadratic, false, false, ti}

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type model struct {
	Choice    int
	Chosen    bool
	Quitting  bool
	TextInput textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.Quitting {
		return m, tea.Quit
	}

	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()

		if k == "q" && m.Chosen {
			m.Chosen = false
		} else if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.ExitAltScreen
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
		return positiveStyle.Render("See you later!")
	}
	if !m.Chosen {
		s = choicesView(m)
	} else {
		switch m.Choice {
		case choiceQuadratic:
			s = quadraticView(m)
		case choicePoisson:
			s = poissonView(m)
		case choiceChi:
			s = chiView(m)
		case choiceBinarySearch:
			s = binarySearchView(m)
		case choiceFibonacci:
			s = fibonacciView(m)
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
			if m.Choice == choiceLen-1 {
				m.Choice = 0
			} else {
				m.Choice++
			}
		case "k", "up":
			if m.Choice == 0 {
				m.Choice = choiceLen - 1
			} else {
				m.Choice--
			}
		case "enter":
			m.Chosen = true
      m.TextInput.SetValue("")
      m.TextInput.Focus()
		}
	}

	return m, nil
}

func choicesView(m model) string {
	c := m.Choice

	tpl := titleStyle.Render("Welcome to ")
	tpl += codeStyle.Render("fxtea")
	tpl += "\n\n%s\n\n"

	splits := []string{help("↑ /k", "up"), help("↓ /j", "down"), help("enter", "confirm"), help("q", "quit")}

	tpl += strings.Join(splits, dotStyle)

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s",
		checkbox("Quadratic", c == choiceQuadratic),
		checkbox("Poisson", c == choicePoisson),
		checkbox("Chai", c == choiceChi),
		checkbox("Binary Search", c == choiceBinarySearch),
		checkbox("Fibonacci", c == choiceFibonacci),
	)

	return fmt.Sprintf(tpl, choices)
}

func header(s string, help []string) string {
	helpText := infoStyle.Render(s)
	helpText += "\n\n%s\n\n"
	helpText += strings.Join(help, dotStyle)

	return helpText
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
