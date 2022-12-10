package UI

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type LogEvent struct {
	Level   string
	Message string
}

type Logger struct {
	Info      func(message string)
	Error     func(message string)
	Important func(message string)
}

type loggerState struct {
	logs []string
}

var levelStyleMap = map[string]lipgloss.Style{
	"info":      lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
	"render":    lipgloss.NewStyle().Foreground(lipgloss.Color("105")),
	"error":     lipgloss.NewStyle().Foreground(lipgloss.Color("161")),
	"important": lipgloss.NewStyle().Foreground(lipgloss.Color("161")),
}

func CreateLogger(bus *Bus) *Component {
	state := loggerState{
		logs: make([]string, 0),
	}

	unsubscribe := bus.Subscribe("log", func(event Event) {
		// FIXME: pager for logging
		if len(state.logs) > 10 {
			state.logs = state.logs[1:]
		}

		if event.Topic == "log" {
			if event, ok := event.Data.(LogEvent); ok {
				state.logs = append(
					state.logs,
					levelStyleMap[event.Level].Render(event.Message),
				)
			}
		} else if event.Topic == "re:render" {
			state.logs = append(state.logs, levelStyleMap["render"].Render("RE-RENDER!"))
		} else {
			state.logs = append(state.logs, fmt.Sprintf("%+v", event))
		}
	})

	return &Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func(width int) string {
			v := "\n"

			if len(state.logs) == 0 {
				v += "No logs (yet)"
			} else {
				// Add all logs onto s, reverse order: newest at top
				for i := len(state.logs) - 1; i >= 0; i-- {
					v += fmt.Sprintf("%s\n", state.logs[i])
				}
			}

			return v
		},
		Destroy: func() {
			unsubscribe()
		},
	}
}
