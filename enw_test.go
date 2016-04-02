package enw

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type EnwSuite struct{}

var _ = Suite(&EnwSuite{})
