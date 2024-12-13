package main

import (
  "fmt"
  "log"
  "strconv"
  "strings"

	"github.com/charmbracelet/bubbles/textinput"
  "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var (
  command = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Background(lipgloss.Color("235"))
  list = lipgloss.NewStyle().Bold(true)
  padding = lipgloss.NewStyle().PaddingTop(1).PaddingLeft(2)
  red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))
  green = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
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
      arguments := strings.Split(m.textInput.Value(), " ")
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
    }

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
  command_list := fmt.Sprintf(
    "Valid commands: %s, %s, %s", 
    command.Render(" enqueue "),
    command.Render(" dequeue "),
    command.Render(" check "),
  )

  if m.show {
    return padding.Render(fmt.Sprintf(
		  "%v\n\n%v\n\n%s",
      command_list,
      m.status,
		  m.textInput.View(),
    )) + "\n"
	} else {
    return padding.Render(fmt.Sprintf(
		  "%v\n\n%v\n\n%s",
      command_list,
      list.Render(fmt.Sprintf("%v", m.queue)),
		  m.textInput.View(),
    )) + "\n"
  }
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
