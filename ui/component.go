package UI

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Function that returns a Component - a la React FC
type FC = func(props *Props) *Component

// Component itself
type Component struct {
	State   interface{}
	Init    func(logger *Logger) tea.Cmd
	Update  func(msg tea.Msg) tea.Cmd
	View    func(width int) string
	Destroy func()
}

type Props struct {
	Outlet *Component
	Params map[string]string
	Logger Logger
}
