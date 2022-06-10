package tui_management

import (
	"fmt"
	"os"
	"os/exec"

	yt "github.com/Github-Reneon/yt-tui/src/api_manager"
	tea "github.com/charmbracelet/bubbletea"
)

func MakeChoices(m *Model) []string {
	var ret []string

	for _, video := range m.videos {
		ret = append(ret, video.Title)
	}

	ret = append(ret, "Search")

	return ret
}

func HomeControl(m *Model, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			// Search button
			if m.cursor == len(m.choices)-1 {
				m.state = "SEARCH"
			} else {
				//start the video
				cmd := exec.Command("mpv", m.videos[m.cursor].Link)
				if err := cmd.Run(); err != nil {
					return tea.Quit
				} else {
					m.videos = []yt.Video{}
					m.choices = MakeChoices(m)
				}
			}
		}
	}
	return nil
}

func SearchControl(m *Model, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.state = "HOME"
		case "enter":
			temp, _ := yt.SearchTitle(m.textInput.Value())
			m.videos = *temp
			m.state = "HOME"
		default:
			m.textInput, cmd = m.textInput.Update(msg)
		}
	}

	return cmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

	switch m.state {
	case "HOME":
		m.choices = MakeChoices(&m)
		cmd = HomeControl(&m, msg)
	case "SEARCH":
		cmd = SearchControl(&m, msg)
	}

	return m, cmd
}

func Start() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
}
