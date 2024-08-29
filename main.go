package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the model
type model struct {
	name       string
	inputting  bool
}

// Initial model
func initialModel() model {
	return model{inputting: true} // Start in inputting mode
}

// Init function - required by the tea.Model interface
func (m model) Init() tea.Cmd {
	return nil // No initial command needed
}

// Update function - required by the tea.Model interface
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.inputting = false // Stop inputting after pressing Enter
			return m, nil
		default:
			if m.inputting {
				m.name += msg.String() // Append character to the name
			}
		}
	}

	return m, nil
}

// View function - required by the tea.Model interface
func (m model) View() string {
	if m.inputting {
		return fmt.Sprintf("What's your name? %s", m.name)
	}
	return fmt.Sprintf("Hello, %s!\nPress q to quit.", m.name)
}

// Main function
func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
