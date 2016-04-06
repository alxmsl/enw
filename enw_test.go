package enw

import (
	"runtime"

	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type EnwSuite struct{}

var _ = Suite(&EnwSuite{})

func (e *EnwSuite) TestStartAndStopWatching(c *C) {
	n := runtime.NumGoroutine()
	Watch("VAR", func(string, string) {})
	c.Assert(runtime.NumGoroutine(), Equals, n+1)

	Forget("VAR")
	runtime.Gosched()
	c.Assert(runtime.NumGoroutine(), Equals, n)
}
