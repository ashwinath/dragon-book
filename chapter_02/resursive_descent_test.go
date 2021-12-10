package two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringToRuneList(value string) []rune {
	return []rune(value)
}

func TestPostiveNegativeA(t *testing.T) {
	var tests = []struct {
		name       string
		testString string
		hasError   bool
	}{
		{
			name:       "nominal",
			testString: "a\n",
			hasError:   false,
		},
		{
			name:       "nominal",
			testString: "+++--+-+-+-+-+++-a-aaaaaaaaaaaaaaaaaa\n",
			hasError:   false,
		},
		{
			name:       "nominal short",
			testString: "+aa\n",
			hasError:   false,
		},
		{
			name:       "failure short",
			testString: "+a\n",
			hasError:   true,
		},
		{
			name:       "failure",
			testString: "c\n",
			hasError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser{
				value: stringToRuneList(tt.testString),
			}
			err := p.PositiveNegativeA()
			if !tt.hasError {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestMatchingBrackets(t *testing.T) {
	var tests = []struct {
		name       string
		testString string
		hasError   bool
	}{
		{
			name:       "nominal empty",
			testString: "\n",
			hasError:   false,
		},
		{
			name:       "nominal",
			testString: "()\n",
			hasError:   false,
		},
		{
			name:       "nominal nested",
			testString: "(((((()())))))\n",
			hasError:   false,
		},
		{
			name:       "nominal chained",
			testString: "()()()()\n",
			hasError:   false,
		},
		{
			name:       "failure",
			testString: "c\n",
			hasError:   true,
		},
		{
			name:       "failure unmatching short",
			testString: "(\n",
			hasError:   true,
		},
		{
			name:       "failure nested long",
			testString: "((((())))\n",
			hasError:   true,
		},
		{
			name:       "failure consecutive long",
			testString: "()(()\n",
			hasError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser{
				value: stringToRuneList(tt.testString),
			}
			err := p.Brackets()
			if !tt.hasError {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
