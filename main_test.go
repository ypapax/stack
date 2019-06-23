package stack

import "testing"

func TestCut(t *testing.T) {
	type input struct {
		text        string
		skip, limit int
		expected    string
	}
	for _, inp := range []input{
		{`1\n2`, 2, 1, `1\n2`}, // should not panic
		{`1\n2`, 2, 0, `1\n2`}, // should not panic
		{`1\n2`, 1, 1, `2`},
		{`1\n2`, 1, 2, `1\n2`}, // should not panic
	} {
		Cut(inp.text, inp.skip, inp.limit) // test pass if no panic
	}
}
