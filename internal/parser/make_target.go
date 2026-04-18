package parser

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
