package parser

import (
	"strings"
)

// MakeTarget represents a target in a Makefile
type MakeTarget struct {
	Name          string   // name of the target, for fuzzy filtering
	Comment       string   // comment describing the target, for fuzzy filtering
	Prerequisites []string // for preview
	Recipe        []string // for preview
}

func NewMakeTarget(name, comment string, prerequisites, recipe []string) *MakeTarget {
	return &MakeTarget{
		Name:          name,
		Comment:       comment,
		Prerequisites: prerequisites,
		Recipe:        recipe,
	}
}

func (t *MakeTarget) Title() string {
	return t.Name
}

func (t *MakeTarget) Description() string {
	return t.Comment
}

func (t *MakeTarget) FilterValue() string {
	return t.Name
}

func (t *MakeTarget) String() string {
	var s strings.Builder

	if t.Comment != "" {
		s.WriteString("# ")
		s.WriteString(t.Comment)
		s.WriteString("\n")
	}

	s.WriteString(t.Name)
	s.WriteString(": ")
	if len(t.Prerequisites) > 0 {
		s.WriteString(strings.Join(t.Prerequisites, " "))
	}
	s.WriteString("\n")

	for _, line := range t.Recipe {
		s.WriteString("\t")
		s.WriteString(line)
		s.WriteString("\n")
	}

	return s.String()
}
