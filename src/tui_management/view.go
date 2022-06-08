package tui_management

import (
	"fmt"
)

func (m Model) View() string {
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
			checked := " "

			if _, ok := m.selected[i]; ok {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		} else {
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}
	s += "\nPress q to quit.\n"

	return s
}
