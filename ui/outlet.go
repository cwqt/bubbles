package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type outletModel struct {
	active map[string]*Component
}

func CreateOutlet(routing RoutingTable, paths *[]string, depth int) *Component {
	m := outletModel{
		active: map[string]*Component{},
	}

	return &Component{
		Model: &m,
		Init: func() tea.Cmd {
			return func() tea.Msg {
				return Render{}
			}
		},
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := Cmds()

			if 0 < len(*paths) && depth < len(*paths) {
				head := (*paths)[depth]

				if route, ok := m.active[head]; ok {
					cmds.Append(route.Update(msg))
				} else {
					for path, route := range m.active {
						if path != head {
							route.Destroy()
							delete(m.active, path)
						}
					}

					if route, ok := routing[head]; ok {
						m.active[head] = route.Create(&Props{
							Outlet: CreateOutlet(route.Children, paths, depth+1),
						})

						cmds.Append(m.active[head].Init())
					}
				}
			}

			return cmds.AsCmd()
		},
		View: func(width int) string {
			s := ""

			for _, path := range *paths {
				if component, ok := m.active[path]; ok {
					s += component.View(width)
				}
			}

			return s
		},
		Destroy: func() {
			for _, component := range m.active {
				component.Destroy()
			}
		},
	}
}
