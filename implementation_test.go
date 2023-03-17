package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func TestPostfixToPrefix(t *testing.T) { TestingT(t) }

type PostfixToPrefixSuite struct{}

var _ = Suite(&PostfixToPrefixSuite{})

func (s *PostfixToPrefixSuite) TestPostfixToPrefix(c *C) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{"short", "2 2 +", "+ 2 2", nil},
		{"long", "4 2 - 3 * 5 +", "+ * - 4 2 3 5", nil},
		{"complex", "2 20 * 2 / 3 4 + 3 2 ^ * + 6 - 15 +", "+ - + / * 2 20 2 * + 3 4 ^ 3 2 6 15", nil},
		{"empty", "", "", InvalidExpressionError{}},
		{"not enough operators", "2 5 7 +", "", InvalidExpressionError{}},
		{"too many operators", "- + / * 1 2 3", "", InvalidExpressionError{}},
		{"invalid token", " \\ 2 2", "", InvalidExpressionError{}},
	}

	for _, test := range tests {

		c.Logf("CASE: %s", test.name)

		output, err := PostfixToPrefix(test.input)

		status := c.Check(output, Equals, test.output)
		status = status && c.Check(err, DeepEquals, test.err)

		if status {
			c.Logf("PASSED\n")
		} else {
			c.Logf("FAILED\n")
		}
	}
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("4 2 - 3 * 5 +")
	fmt.Println(res)

	// Output:
	// + * - 4 2 3 5
}
