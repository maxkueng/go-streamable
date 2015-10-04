package streamable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bytesToString(t *testing.T) {
	byteArray := []byte("yolo")
	str := bytesToString(byteArray)
	assert.Equal(t, "yolo", str)
}
