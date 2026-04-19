package ui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/balaji01-4d/cake/internal/parser"
)

func TestEnterSelectsTargetAndSetsCommand(t *testing.T) {
	target := parser.NewMakeTarget("hello", "", nil, nil)
	m := New([]*parser.MakeTarget{target}, `make -f "temp.mk"`).(Model)

	updatedModel, _ := m.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter}))
	updated := updatedModel.(Model)

	if updated.CurrentTarget == nil || updated.CurrentTarget.Name != "hello" {
		t.Fatalf("expected selected target to be hello, got %#v", updated.CurrentTarget)
	}

	if updated.FinalCmd != `make -f "temp.mk" hello` {
		t.Fatalf("expected final command to be %q, got %q", `make -f "temp.mk" hello`, updated.FinalCmd)
	}
}

func TestShiftEnterPrefillsEditableCommand(t *testing.T) {
	target := parser.NewMakeTarget("hello", "", nil, nil)
	m := New([]*parser.MakeTarget{target}, `make -f "temp.mk"`).(Model)

	updatedModel, _ := m.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter, Mod: tea.ModShift}))
	updated := updatedModel.(Model)

	if !updated.editCmdDialog {
		t.Fatal("expected edit command dialog to be open")
	}

	if updated.Input.Value() != `make -f "temp.mk" hello` {
		t.Fatalf("expected prefilled command %q, got %q", `make -f "temp.mk" hello`, updated.Input.Value())
	}
}
