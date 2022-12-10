package UI

import (

	tea "github.com/charmbracelet/bubbletea"
)

func CreateProgram(router *Router) *tea.Program {
	app := CreateApp(router)
	p := tea.NewProgram(app, tea.WithAltScreen())

	return p
}
