package tree_sitter_gmpl_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_gmpl "github.com/florian-walter/tree-sitter-gmpl/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gmpl.Language())
	if language == nil {
		t.Errorf("Error loading GNU MathProg Language grammar")
	}
}
