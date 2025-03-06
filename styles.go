package main

import "github.com/charmbracelet/lipgloss"

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	keyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("246")).Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("203")).Italic(true)
	infoStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	listStyle     = lipgloss.NewStyle().Bold(true)
	resultStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("156"))
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Bold(true)
	positiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Margin(1).MarginLeft(2)
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	codeStyle     = lipgloss.NewStyle().Background(lipgloss.Color("237")).PaddingLeft(1).PaddingRight(1).Bold(true)
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)
