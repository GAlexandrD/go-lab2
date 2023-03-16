package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToPrefix(c *C) {
    c.Assert("AAA", Equals, "AAA")
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
