package components

import (
	UI "gitlab.com/cxss/bubbles/ui"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type appState struct {
}

var Wrapper = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder())
	// BorderForeground(lipgloss.Color(Style.DarkGray))

func App(props *UI.Props) *UI.Component {
	state := appState{}

	props.Logger.Info("Hello World!")
	props.Logger.Info("Hello World!")
	props.Logger.Info("Hello World!")
	props.Logger.Info("Hello World!")
	props.Logger.Info("Hello World!")

	return &UI.Component{
		State: state,
		Init: func() tea.Cmd {
			return props.Outlet.Init(&props.Logger)
		},
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := UI.Cmds()

			switch msg := msg.(type) {
			case tea.KeyMsg:
				props.Logger.Important(msg.String())
			}

			cmds.Append(props.Outlet.Update(msg))
			return cmds.AsCmd()
		},
		View: func(width int) string {
			s := "bubbles\n"
			s += Wrapper.Width(width).Render(props.Outlet.View(width))
			return s
		},
		Destroy: props.Outlet.Destroy,
	}
}
