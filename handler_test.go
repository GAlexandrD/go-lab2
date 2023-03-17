package lab2

import (
	"bytes"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ComputeHandlerSuite struct{}

var _ = Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestComputeHandler(c *C) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{"valid", "2 2 +", "+ 2 2", nil},
		{"invalid", "2 +", "", InvalidExpressionError{}},
	}
	for _, test := range tests {

		c.Logf("CASE: %s", test.name)

		input := strings.NewReader(test.input)
		output := bytes.NewBuffer(nil)

		handler := &ComputeHandler{input, output}
		err := handler.Compute()

		status := c.Check(output.String(), Equals, test.output)
		status = status && c.Check(err, DeepEquals, test.err)
		if status {
			c.Logf("PASSED\n")
		} else {
			c.Logf("FAILED\n")
		}
	}
}
