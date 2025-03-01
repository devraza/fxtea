package main

import (
	"fmt"
	"strconv"
	"strings"

	"fxtea/fx"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar = " • "
)

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	keyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("246")).Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("203")).Italic(true)
	infoStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Bold(true)
	positiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Padding(1).PaddingLeft(2)
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	codeStyle     = lipgloss.NewStyle().Background(lipgloss.Color("236")).PaddingLeft(1).PaddingRight(1)
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

func main() {
	ti := textinput.New()
	ti.Focus()
	ti.Width = 20

	initialModel := model{0, false, false, ti}

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
		if k == "q" || k == "esc" || k == "ctrl+c" {
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
		case 0:
			s = quadraticView(m)
		case 1:
			s = poissonView(m)
		case 2:
			s = chaiView(m)
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
		"%s\n%s\n%s",
		checkbox("Quadratic", c == 0),
		checkbox("Poisson", c == 1),
		checkbox("Chai", c == 2),
	)

	return fmt.Sprintf(tpl, choices)
}

func quadraticView(m model) string {
	headerContent := fmt.Sprintf("Enter the values for a quadratic in the form %s", codeStyle.Render("ax² + bx + c"))
	helpText := header(headerContent, []string{help("q", "quit")})

	arguments := strings.Split(m.TextInput.Value(), " ")
	m.TextInput.Placeholder = "a b c"

	var result string

	if len(arguments) == 3 {
		var coefficients []float64

		for i := range arguments {
			parsed, _ := strconv.ParseFloat(arguments[i], 64)
			coefficients = append(coefficients, parsed)
		}

		roots := fx.Quadratic(coefficients[0], coefficients[1], coefficients[2])

		if fx.NaN(roots[0]) || fx.NaN(roots[1]) {
			result = errorStyle.Render("The roots are complex")
		} else if roots[0] != roots[1] {
			result = fmt.Sprintf(
				"The roots are %v and %v",
				keywordStyle.Render(fx.FormatFloat(roots[0])),
				keywordStyle.Render(fx.FormatFloat(roots[1])),
			)
		} else {
			result = fmt.Sprintf(
				"The root is %v",
				keywordStyle.Render(fx.FormatFloat(roots[0])),
			)
		}
	}

	content := fmt.Sprintf(
		"%s\n\n%v",
		m.TextInput.View(),
		result,
	)

	return fmt.Sprintf(helpText, content)
}

func poissonView(m model) string {
	helpText := header("Enter the rate and the value of x", []string{help("q", "quit")})

	arguments := strings.Split(m.TextInput.Value(), " ")
	m.TextInput.Placeholder = "λ x"

	var result string

	if len(arguments) == 2 {
		lambda, _ := strconv.ParseFloat(arguments[0], 64)
		x, _ := strconv.ParseUint(arguments[1], 10, 64)

		result = fmt.Sprintf(
			"The cumulative probability is: %s",
			keywordStyle.Render(fmt.Sprintf("%.4f", fx.PoissonCD(lambda, x))),
		)
	}

	content := fmt.Sprintf(
		"%s\n\n%v",
		m.TextInput.View(),
		result,
	)

	return fmt.Sprintf(helpText, content)
}

func chaiView(m model) string {
	helpText := header("Enter the degrees of freedom and the significance level", []string{help("q", "quit")})

	arguments := strings.Split(strings.TrimSpace(m.TextInput.Value()), " ")
	m.TextInput.Placeholder = "ν α"

	var result string

	if len(arguments) == 2 {
		var floatArgs []float64

		for i := range arguments {
			parsed, _ := strconv.ParseFloat(arguments[i], 64)
			floatArgs = append(floatArgs, parsed)
		}

		critical := fx.ChiCritical(floatArgs[0], floatArgs[1])
		result = fmt.Sprintf(
			"The critical value is %v",
			keywordStyle.Render(fx.FormatFloat(critical)),
		)
	}

	content := fmt.Sprintf("%s\n\n%v", m.TextInput.View(), result)

	return fmt.Sprintf(helpText, content)
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
