package tui_management

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	uploaded bool
}

func initialModel() Model {
	return Model{
		choices: []string{"test1", "test2", "test3", "-- Confirm --"},

		selected: make(map[int]struct{}),
		uploaded: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
