package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"fxtea/fx"
)

var (
	command = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Background(lipgloss.Color("235"))
	list    = lipgloss.NewStyle().Bold(true)
	padding = lipgloss.NewStyle().PaddingTop(1).PaddingLeft(2)
	red     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))
	green   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
)

func enqueue(queue []int64, number string) []int64 {
	integer, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if len(queue) < 5 {
		queue = append([]int64{integer}, queue...)
	}

	return queue
}

func dequeue(queue []int64) []int64 {
	if len(queue) > 0 {
		queue = queue[:len(queue)-1]
	}
	return queue
}

func check(queue []int64) string {
	if len(queue) == 0 {
		return fmt.Sprintf("The queue is currently %v", green.Render("empty"))
	} else if len(queue) == 5 {
		return fmt.Sprintf("The queue is currently %v", red.Render("full"))
	} else {
		return fmt.Sprintf("The queue is currently neither %v nor %v", green.Render("empty"), red.Render("full"))
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	queue     []int64
	status    string
	show      bool
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "enqueue 9"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func quadraticModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			arg := os.Args[1]
			arguments := strings.Split(m.textInput.Value(), " ")
			if arg == "queue" {
				switch arguments[0] {
				case "enqueue":
					m.show = false
					m.queue = enqueue(m.queue, arguments[1])
				case "dequeue":
					m.show = false
					m.queue = dequeue(m.queue)
				case "check":
					m.show = true
					m.status = check(m.queue)
				default:
					m.show = true
					m.status = red.Render("This action is unsupported")
				}
			} else {
				var variables []float64
				for i := range arguments {
					parsed, _ := strconv.ParseFloat(arguments[i], 64)
					variables = append([]float64{parsed}, variables...)
				}
				m.status = fmt.Sprintf("%v", fx.Quadratic(variables[0], variables[1], variables[2]))
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.show {
		return padding.Render(fmt.Sprintf(
			"%v\n\n%s",
			m.status,
			m.textInput.View(),
		)) + "\n"
	} else {
		return padding.Render(fmt.Sprintf(
			"%v\n\n%s",
			m.status,
			m.textInput.View(),
		)) + "\n"
	}
}

func main() {
	arg := os.Args[1]
	switch arg {
	case "queue":
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	case "quadratic":
		p := tea.NewProgram(quadraticModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
