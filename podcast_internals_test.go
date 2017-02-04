package podcast

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestStringError(t *testing.T) {
	t.Parallel()

	// arrange
	e := "TestEncodeError error result"
	p := Podcast{}
	p.encode = func(w io.Writer, o interface{}) error {
		return errors.New(e)
	}

	// act
	r := p.String()

	// assert
	assert.Contains(t, r, e)
}

func TestEncodeError(t *testing.T) {
	t.Parallel()

	// arrange
	p := New("", "", "", nil, nil)
	b := []byte{}
	w := bytes.NewBuffer(b)
	c := new(chan bool)

	// act
	err := p.encode(w, c)

	// assert
	assert.Error(t, err)
}
