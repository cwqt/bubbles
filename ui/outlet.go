package UI

import (
	tea "github.com/charmbracelet/bubbletea"
)

type outletState struct {
	active map[string]*Component
	logger *Logger
}

type Render struct{}

func CreateOutlet(routes Routes, paths *[]string, depth int) *Outlet {
	state := outletState{
		active: map[string]*Component{},
	}

	return &Outlet{
		State: &state,
		Init: func(logger *Logger) tea.Cmd {
			state.logger = logger

			return func() tea.Msg {
				return Render{}
			}
		},
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := Cmds()

			if 0 < len(*paths) && depth < len(*paths) {
				head := (*paths)[depth]

				if route, ok := state.active[head]; ok {
					cmds.Append(route.Update(msg))
				} else {
					for path, route := range state.active {
						if path != head {
							route.Destroy()
							delete(state.active, path)
						}
					}

					if route, ok := routes[head]; ok {
						state.active[head] = route.Component(&Props{
							Outlet: CreateOutlet(route.Children, paths, depth+1),
							Logger: *state.logger,
						})

						cmds.Append(state.active[head].Init())
					}
				}
			}

			return cmds.AsCmd()
		},
		View: func(width int) string {
			s := ""

			for _, path := range *paths {
				if component, ok := state.active[path]; ok {
					s += component.View(width)
				}
			}

			return s
		},
		Destroy: func() {
			for _, component := range state.active {
				component.Destroy()
			}
		},
	}
}
