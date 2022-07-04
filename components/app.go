package components

import (
	UI "gitlab.com/cxss/bubbles/ui"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type appModel struct {
}

var Wrapper = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder())
	// BorderForeground(lipgloss.Color(Style.DarkGray))

func App(props *UI.Props) *UI.Component {
	m := appModel{}

	// props.Logger.Info("Hello World!")

	return &UI.Component{
		Model: m,
		Init:  props.Outlet.Init,
		Update: func(msg tea.Msg) tea.Cmd {
			cmds := UI.Cmds()
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
