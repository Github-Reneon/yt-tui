package tui_management

import (
	"fmt"
)

func (m Model) View() string {
	s := ""

	switch m.state {
	case "HOME":
		s = HomeView(&m)
	case "SEARCH":
		s = SearchView(&m)
	default:
		s = HomeView(&m)
	}

	return s
}

func SearchView(m *Model) string {
	s := ""

	s = fmt.Sprintf("YouTube video search Title: %s",
		m.textInput.View())

	s += "\n"

	s += "Press ctrl-c to go back to home.\n"

	return s
}

func HomeView(m *Model) string {
	s := ""

	if m.uploaded {
		s += "Uploaded!\n"
	} else {
		s += "What should we blah blah\n"
	}

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}
		if i != len(m.choices)-1 {
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		} else {
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}
	s += "\nPress q to quit.\n"

	return s
}
