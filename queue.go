package main

import (
  "fmt"
  "strconv"

	"github.com/charmbracelet/lipgloss"
)

var (
	list    = lipgloss.NewStyle().Bold(true)
	padding = lipgloss.NewStyle().PaddingTop(1).PaddingLeft(2)
	red     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))
	green   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
)

func Enqueue(queue []int64, number string) []int64 {
	integer, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if len(queue) < 5 {
		queue = append([]int64{integer}, queue...)
	}

	return queue
}

func Dequeue(queue []int64) []int64 {
	if len(queue) > 0 {
		queue = queue[:len(queue)-1]
	}
	return queue
}

func Check(queue []int64) string {
	if len(queue) == 0 {
		return fmt.Sprintf("The queue is currently %v", green.Render("empty"))
	} else if len(queue) == 5 {
		return fmt.Sprintf("The queue is currently %v", red.Render("full"))
	} else {
		return fmt.Sprintf("The queue is currently neither %v nor %v", green.Render("empty"), red.Render("full"))
	}
}
