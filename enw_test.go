package enw

import (
        "testing"
        . "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type EnwSuite struct{}

var _ = Suite(&EnwSuite{})
