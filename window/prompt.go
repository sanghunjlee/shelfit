package window

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type promptModel struct {
	label     string
	textInput textinput.Model
	quitting  bool
	err       error
}

func Prompt(label string, placeholder string) (string, error) {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	pm := promptModel{
		label:     label,
		textInput: ti,
		err:       nil,
	}

	p := tea.NewProgram(pm)
	m, err := p.Run()

	return m.View(), err
}

func (m promptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m promptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m promptModel) View() string {
	if m.quitting {
		return m.textInput.Value()
	}
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.label,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
