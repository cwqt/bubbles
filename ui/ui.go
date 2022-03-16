package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	router *Router
	outlet *Component
}

type Render struct{}

func Create(router *Router) model {
	return model{
		router: router,
		outlet: router.Outlet,
	}
}

func (m model) Init() tea.Cmd {
	return m.outlet.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// All route change messages bubble up to this point
	case RouteChangeMsg:
		m.router.Navigate(msg.Path...)
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			return m, func() tea.Msg {
				return RouteChangeMsg{
					Path: []string{"index", "menu"},
				}
			}
		case "ctrl+c", "q":
			fmt.Println("Goodbye!")
			return m, tea.Quit
		}
	}

	return m, m.outlet.Update(msg)
}

func (m model) View() string {
	s := ""

	// Show the current routing path
	s += "[ "
	path := *m.router.GetPath()
	for index, segment := range path {
		s += segment
		if index != len(path)-1 {
			s += " / "
		} else {
			s += " ]\n"
		}
	}

	return s + m.outlet.View(PROGRAM_WIDTH)
}
