package UI

import (
	tea "github.com/charmbracelet/bubbletea"
)

func CreateProgram(router *Router) *tea.Program {
	app := CreateApp(router)
	p := tea.NewProgram(app, tea.WithAltScreen())

	// Re-render the app whenever a logging message is sent to the bus
	app.bus.Subscribe("log", func(event Event) {
		p.Send(Render{})
	})

	return p
}
