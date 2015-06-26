package yaml_test

import (
	. "github.com/dvln/check"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})
