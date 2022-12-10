package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	UI "gitlab.com/cxss/bubbles/ui"
)

type Post struct {
	id    string
	title string
}

type postsState struct {
	posts []Post
}

func Posts(props *UI.Props) *UI.Component {
	state := postsState{
		posts: make([]Post, 0),
	}

	// TODO: figure out how to source these?
	state.posts = append(state.posts, Post{
		id:    "hello-world",
		title: "Hello world this is my first post",
	}, Post{
		id:    "another-post",
		title: "This is the second post",
	})

	return &UI.Component{
		Init: func() tea.Cmd {
			return nil
		},
		Update: func(msg tea.Msg) tea.Cmd {
			return nil
		},
		View: func(width int) string {
			s := ""

			for i, post := range state.posts {
				s += fmt.Sprintln(i, post.id, post.title)
			}

			return s
		},
		Destroy: func() {},
	}
}
