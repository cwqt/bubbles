package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	Component "gitlab.com/cxss/bubbles/components"
	UI "gitlab.com/cxss/bubbles/ui"
)

func main() {
	// Define a router with paths & the components that should render on those paths
	Router := UI.CreateRouter(UI.RoutingTable{
		"index": {
			Create: Component.App,
			Children: UI.RoutingTable{
				"posts": {
					Create: Component.Posts,
				},
				// Route parameters
				// "posts/:id": {
				// 	Create: func(props *UI.Props) {
				// 		return Component.Post(props.Params.id)
				// 	},
				// },
			},
		},
	}, []string{"index", "posts"})

	p := tea.NewProgram(UI.Create(&Router), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
