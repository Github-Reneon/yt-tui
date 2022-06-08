package tui_management

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices) {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				if m.cursor == len(m.choices)-1 {
					m.uploaded = !m.uploaded
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		}
	}

	return m, nil
}

func Start() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

}
