package tregexp

import (
	"testing"
	"regexp"
)

func TestRegexp(t *testing.T) {
	t.Log(regexp.QuoteMeta("[aaa]"))
}
