package tui_management

import (
	"time"

	yt "github.com/Github-Reneon/yt-tui/src/api_manager"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

type Model struct {
	choices   []string
	cursor    int
	uploaded  bool
	state     string
	textInput textinput.Model
	videos    []yt.Video
}

func initialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Video Title"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		choices:   []string{"Search"},
		uploaded:  false,
		state:     "HOME",
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(Tick(), tea.EnterAltScreen, textinput.Blink)
}

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
