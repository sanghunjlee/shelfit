package window

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type tagInputModel struct {
	textInput textinput.Model
	tags      []string
	quitting  bool
	err       error
}

func (m tagInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m tagInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// case tea.KeyRunes:
		// 	switch msg.String() {
		// 	case ",":

		// 	}
		case tea.KeyDelete, tea.KeyBackspace:
			if m.textInput.Value() == "" {
				l := len(m.tags)
				if l > 0 {
					m.tags = m.tags[:l-1]
				}
			}
		case tea.KeyEnter:
			if m.textInput.Value() == "" {
				m.quitting = true
				return m, tea.Quit
			}
			fallthrough
		case tea.KeyTab, tea.KeySpace:
			m.tags = append(m.tags, m.textInput.Value())
			m.textInput.Reset()
			return m, cmd
		case tea.KeyCtrlC, tea.KeyEsc:
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

func (m tagInputModel) View() string {
	if m.quitting {
		return strings.Join(m.tags, " ")
	}

	tags := make([]string, 0)
	for _, t := range m.tags {
		tags = append(tags, "#"+t)
	}

	return fmt.Sprintf(
		"%s\n\n%s %s\n\n%s",
		"Enter tags:",
		strings.Join(tags, " "),
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func TagInput(tags []string) (string, error) {
	ti := textinput.New()
	ti.Placeholder = "new-tag"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	tim := tagInputModel{
		textInput: ti,
		tags:      tags,
		err:       nil,
	}

	p := tea.NewProgram(tim)
	m, err := p.Run()

	return m.View(), err
}
