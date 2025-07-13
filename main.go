package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type list_item struct {
	task_name, task_desc string
}

func (i list_item) Title() string       { return i.task_name }
func (i list_item) Description() string { return i.task_desc }
func (i list_item) FilterValue() string { return i.task_name }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	tasks := []list.Item{
		list_item{task_name: "brush teeth", task_desc: "clean myself time"},
		list_item{task_name: "code", task_desc: "grind time!!!!"},
		list_item{task_name: "sleep", task_desc: "zzzzzzzzzzzzz"},
		list_item{task_name: "cook", task_desc: "i like food"},
		list_item{task_name: "workout", task_desc: "i get stronk"},
	}

	m := model{list: list.New(tasks, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "my to-do list"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}

}
