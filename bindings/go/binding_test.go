package tree_sitter_fluent_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-fluent"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_fluent.Language())
	if language == nil {
		t.Errorf("Error loading Fluent grammar")
	}
}
