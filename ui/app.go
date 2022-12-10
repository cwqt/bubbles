package UI

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type appState struct {
	router *Router
	outlet *Component
	logger *Logger
	logs   *Component
}

func CreateApp(router *Router) appState {
	bus := CreateBus()

	logger := &Logger{
		Info: func(message string) {
			bus.Publish("log", LogEvent{Level: "info", Message: fmt.Sprintf("%v", message)})
		},
		Error: func(message string) {
			bus.Publish(
				"log",
				LogEvent{Level: "error", Message: fmt.Sprintf("%v", message)},
			)
		},
		Important: func(message string) {
			bus.Publish(
				"log",
				LogEvent{Level: "important", Message: fmt.Sprintf("%v", message)},
			)
		},
	}

	return appState{
		router: router,
		outlet: router.Outlet,
		logger: logger,
		logs:   CreateLogger(&bus),
	}
}

func (state appState) Init() tea.Cmd {
	return state.outlet.Init(state.logger)
}

func (m appState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// All route change messages bubble up to this point
	case RouteChange:
		m.router.Navigate(msg.Path...)
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			return m, func() tea.Msg {
				return RouteChange{
					Path: []string{"index", "menu"},
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	cmds := Cmds()
	cmds.Append(m.logs.Update(msg))
	cmds.Append(m.outlet.Update(msg))

	return m, cmds.AsCmd()
}

func (m appState) View() string {
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

	s += m.outlet.View(PROGRAM_WIDTH)
	s += m.logs.View(PROGRAM_WIDTH)

	return s
}
