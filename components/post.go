package components

import (
	tea "github.com/charmbracelet/bubbletea"
	UI "gitlab.com/cxss/bubbles/ui"
)

func PostComponent(postId string) UI.FC {
	return func(props *UI.Props) *UI.Component {

		return &UI.Component{
			Init: func() tea.Cmd {
				return nil
			},
			Update: func(msg tea.Msg) tea.Cmd {
				return nil
			},
			View: func(width int) string {
				s := ""

				s += postId

				return s
			},
			Destroy: func() {},
		}

	}

}
