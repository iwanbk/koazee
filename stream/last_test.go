package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/internal/last"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Last(t *testing.T) {
	s := stream.New([]string{"hello", "my", "friend"})
	last := s.Last().Val()
	assert.Equal(t, "friend", last)
}

func TestStream_Last_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.EmptyStream(last.OpCode, "It can not be taken an element from an empty Stream"),
		stream.New(nil).Last().Err())

	assert.Equal(
		t,
		errors.EmptyStream(last.OpCode, "It can not be taken an element from an empty Stream"),
		stream.New([]int{}).Last().Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode,
		stream.New([]int{}).Add("home").Last().Err().Operation())

}
