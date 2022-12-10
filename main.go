package main

import (
	"fmt"
	"os"

	Components "gitlab.com/cxss/bubbles/components"
	UI "gitlab.com/cxss/bubbles/ui"
)

func main() {
	// Define a Router with paths & the components that should render on those paths
	router := UI.CreateRouter(UI.Routes{
		"index": {
			Component: Components.App,
			// Children: UI.Routes{
			// 	"posts": {
			// 		Create: Component.Posts,
			// 	},
			// 	// Route parameters
			// 	"posts/:id": {
			// 		Create: func(props *UI.Props) *UI.Component {
			// 			return Component.PostComponent(props.Params["id"])(props)
			// 		},
			// 	},
			// },
		},
	}, []string{"index", "posts"})

	// Start the Program
	if err := UI.CreateProgram(&router).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
