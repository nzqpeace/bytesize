// Package bytesize_test provides ...
package bytesize_test

import (
	"testing"

	"github.com/nzqpeace/bytesize"
	"github.com/stretchr/testify/assert"
)

func parse(t *testing.T, s string) int64 {
	n, err := bytesize.Parse(s)
	if err != nil {
		t.Error(err)
		return 0
	}
	return n
}

func TestStringToByteSize(t *testing.T) {
	assert.Equal(t, parse(t, "1"), bytesize.B)
	assert.Equal(t, parse(t, "1B"), bytesize.B)
	assert.Equal(t, parse(t, "10KB"), 10*bytesize.KB)
	assert.Equal(t, parse(t, "123mB"), 123*bytesize.MB)
	assert.Equal(t, parse(t, "56Gb"), 56*bytesize.GB)
	assert.Equal(t, parse(t, "1067t"), 1067*bytesize.TB)
	assert.Equal(t, parse(t, "123P"), 123*bytesize.PB)
}

func TestFloatSize(t *testing.T) {
	v := 1.1 * 1024 * 1024
	assert.Equal(t, parse(t, "1.1MB"), int64(v))
}

func TestError(t *testing.T) {
	v := 1.4 * 1024 * 1024
	assert.NotEqual(t, parse(t, "1.4M"), v)
	assert.NotEqual(t, parse(t, "1Mb"), bytesize.MB+1)
}
