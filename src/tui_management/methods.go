package tui_management

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

}
