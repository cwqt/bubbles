# bubbles

```go
package main

import (
	"fmt"
	"os"
	Component "app/components"

	Bus "sprak/bus"
	Data "sprak/data"
	UI "sprak/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  // Define a router with paths & the components that should render on those paths
	Router := UI.CreateRouter(UI.RoutingTable{
		"index": {
			Create: Component.App,
			Children: UI.RoutingTable{
				"posts": {
					Create: Component.PostList,
				},
        // Route parameters
        "posts/:id": {
          Create: func(id string) {
            return Component.Post(id)
          },
        },
			},
		},
	}, []string{"index", "menu"})

	p := tea.NewProgram(UI.Create(&Router), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
```
