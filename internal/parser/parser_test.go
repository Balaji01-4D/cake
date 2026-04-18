package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/balaji01-4d/cake/internal/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDataPath = "testdata"


func TestParseMakefile_EmptyFile(t *testing.T) {
	makefilePath := filepath.Join(testDataPath, "empty_makefile.mk")
	r, err := os.Open(makefilePath)
	require.NoError(t, err)
	defer r.Close()

	targets, err := parser.ParseMakefile(r)
	require.NoError(t, err)
	require.Empty(t, targets, "expected no targets for an empty Makefile")
}

func TestParseMakefile_SingleTarget(t *testing.T) {
	makefilePath := filepath.Join(testDataPath, "single_target.mk")

	f, err := os.Open(makefilePath)
	require.NoError(t, err)
	defer f.Close()

	targets, err := parser.ParseMakefile(f)
	require.NoError(t, err)
	require.Len(t, targets, 1, "expected exactly one target")
	assert.Equal(t, "clean", targets[0].Name)

	expectedRecipes := []string{
		"rm -rf bin/",
		"rm -f *.o",
	}
	assert.Equal(t, expectedRecipes, targets[0].Recipe)
}

func TestParseMakefile_MultipleTargets(t *testing.T) {
	makefilepath := filepath.Join(testDataPath, "multiple_targets.mk")

	f, err := os.Open(makefilepath)
	require.NoError(t, err)
	defer f.Close()

	targets, err := parser.ParseMakefile(f)
	require.NoError(t, err, "failed to parse Makefile")
	require.Len(t, targets, 4, "expected exactly four targets")

	execpectedTargets := []parser.MakeTarget{
		{
			Name:          "all",
			Comment:       "",
			Prerequisites: []string{"build", "test"},
			Recipe:        nil,
		},
		{
			Name:          "build",
			Comment:       "",
			Prerequisites: []string{"main.go", "utils.go"},
			Recipe: []string{
				"go build -o bin/app .",
			},
		},
		{
			Name:          "test",
			Comment:       "",
			Prerequisites: []string{"main.go"},
			Recipe: []string{
				"go test -v ./...",
			},
		},
		{
			Name:          "clean",
			Comment:       "",
			Prerequisites: nil,
			Recipe: []string{
				"rm -rf bin/",
				"rm -f *.o",
			},
		},
	}

	for i, expected := range execpectedTargets {
		assert.Equal(t, expected.Name, targets[i].Name, "target name mismatch")
		assert.Equal(t, expected.Comment, targets[i].Comment, "target comment mismatch")
		assert.Equal(t, expected.Prerequisites, targets[i].Prerequisites, "target prerequisites mismatch")
		assert.Equal(t, expected.Recipe, targets[i].Recipe, "target recipe mismatch")
	}
}
