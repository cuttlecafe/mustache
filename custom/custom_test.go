package main

import (
	"testing"

	"github.com/cuttlecafe/mustache"
)

func TestCustomDelimiter(t *testing.T) {
	var (
		in  = "hello <% name %>"
		out = "hello world"
	)
	mustache.TagStart = "<%"
	mustache.TagEnd = "%>"
	s, err := mustache.ParseString(in)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]interface{}{
		"name": "world",
	}
	result, err := s.Render(vars)
	if err != nil {
		t.Fatal(err)
	}
	if result != out {
		t.Errorf("wont %s, but got %s", out, result)
	}
}
