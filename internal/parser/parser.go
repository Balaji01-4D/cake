package parser

import (
	"fmt"
	"io"
	"strings"

	goMake "github.com/unmango/go-make"
	"github.com/unmango/go-make/ast"
)

func ParseMakefile(r io.Reader) ([]*MakeTarget, error) {
	p := goMake.NewParser(r, nil)
	makeAst, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	targetMap := make(map[string]*MakeTarget)
	targetOrder := []string{}

	for _, obj := range makeAst.Contents {
		rule, isRUle := obj.(*ast.Rule)
		if !isRUle {
			continue
		}

		for _, targetExpr := range rule.Targets {
			targetName := fmt.Sprint(targetExpr)

			if strings.HasPrefix(targetName, ".") {
				continue
			}

			preReqs := toStringSlice(rule.PreReqs)
			recipe := recipieLines(rule.Recipes)

			if existingTarget, exists := targetMap[targetName]; exists {
				existingTarget.Prerequisites = append(existingTarget.Prerequisites, preReqs...)
				if len(recipe) > 0 {
					existingTarget.Recipe = recipe
				}
			} else {
				targetOrder = append(targetOrder, targetName)
				targetMap[targetName] = NewMakeTarget(
					targetName, "", preReqs, recipe,
				)
			}
		}
	}

	var targets []*MakeTarget
	for _, targetName := range targetOrder {
		targets = append(targets, targetMap[targetName])
	}

	return targets, nil
}


func recipieLines(recipes []*ast.Recipe) []string {
	var lines []string
	for _, recipe := range recipes {
		lines = append(lines, fmt.Sprint(recipe))
	}
	return lines	
}

func toStringSlice(exprs []ast.Expr) []string {
	var result []string
	for _, expr := range exprs {
		result = append(result, fmt.Sprint(expr))
	}
	return result
}