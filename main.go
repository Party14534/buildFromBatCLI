package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
  Bold(true).
  Foreground(lipgloss.Color("#FAFAFA")).
  Background(lipgloss.Color("#7D56F4")).
  PaddingTop(2).
  PaddingLeft(4).
  Width(22)

type model struct {
  options []string
  cursor int
  selected int
}

func initModel() model {
  return model {
    options: []string{
      style.Render("Build Project"),
      style.Render("Etc"),
      style.Render("Quit"),
    },
    selected: 0,
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    
    case "up", "k":
      m.cursor++ 
      m.cursor = m.cursor % len(m.options)

    case "down", "j":
      m.cursor--
      m.cursor = m.cursor % len(m.options)

    case "enter", " ":
      m.selected = m.cursor
    }
  }

  return m, nil
}

func (m model) View() string {
  s := "Thank you for using buildFromBat!\n"

  for i, option := range m.options {
    cursor := " "
    if m.cursor == i {
      cursor = ">"
    }

    checked := " "
    if i == m.selected {
      checked = "x"
    }

    s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, option)
  }

  s += "\nPress q to quit.\n"

  return s
}

func main() {
  fmt.Println(style.Render("Hello kitty"))
  p := tea.NewProgram(initModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v",err)
    os.Exit(1)
  }

}
