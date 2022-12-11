package goutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aibor/aoc/goutils"
)

func TestStackEmpty(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		ts := goutils.Stack[string]{}
		assert.Equal(t, "", ts.Pop())
		ts.Push("H")
		assert.Len(t, ts, 1)
		assert.Equal(t, "H", ts.Pop())
		assert.Equal(t, "", ts.Pop())
	})
	t.Run("int", func(t *testing.T) {
		ts := goutils.Stack[int]{}
		assert.Equal(t, 0, ts.Pop())
		ts.Push(5)
		assert.Len(t, ts, 1)
		assert.Equal(t, 5, ts.Pop())
		assert.Equal(t, 0, ts.Pop())
	})
}

func TestStackNotEmpty(t *testing.T) {
	type testStack = goutils.Stack[int]
	tests := []struct {
		name   string
		stack  testStack
		length int
	}{
		{
			name:   "single value",
			stack:  testStack{3},
			length: 1,
		},
		{
			name:   "multiple values",
			stack:  testStack{5, 6, 7},
			length: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var ts testStack
			ts = tc.stack
			assert.Equal(t, tc.stack[tc.length-1], ts.Pop())
			assert.Len(t, ts, tc.length-1)
			ts.Push(42)
			assert.Len(t, ts, tc.length)
			assert.Equal(t, append(tc.stack[:tc.length-1], 42), ts)
			i, h := ts.Top()
			assert.Equal(t, 42, i)
			assert.Equal(t, tc.length-1, h)
			assert.Len(t, ts, tc.length)
			assert.Equal(t, 42, ts.Pop())
		})
	}
}

func TestIteratorEmpty(t *testing.T) {
	iter := goutils.NewIterator([]int{})
	assert.Panics(t, func() { iter.Value() })
	assert.False(t, iter.Next())
	assert.False(t, iter.Prev())
	assert.False(t, iter.Skip(1))
}

func TestIteratorNonEmpty(t *testing.T) {
	iter := goutils.NewIterator([]int{4, 7, 11})
	t.Run("walk", func(t *testing.T) {
		iter := iter
		assert.Equal(t, 4, iter.Value(), "first Value() should return first value")
		assert.True(t, iter.Next(), "first Next() should succeed")
		assert.Equal(t, 7, iter.Value(), "Value() after first Next() should return second value")
		assert.True(t, iter.Next(), "second Next() should succeed")
		assert.Equal(t, 11, iter.Value(), "Value() after second Next()")
		assert.False(t, iter.Next(), "Next() should fail after last element")
		assert.True(t, iter.Prev(), "Prev() from last element shoould succeed")
		assert.Equal(t, 7, iter.Value(), "Value() after first Prev() should return second value")
		assert.True(t, iter.Prev(), "Prev() from second element should succeed")
		assert.Equal(t, 4, iter.Value(), "Value() after second Prev() should return first value")
		assert.False(t, iter.Prev(), "Prev() should fail from first element")
		assert.True(t, iter.Next(), "Next() from first element should succeed again")
		assert.Equal(t, 7, iter.Value(), "Value() should return second value again after next()")
	})
	t.Run("reset", func(t *testing.T) {
		iter := iter
		iter.Next()
		iter.Next()
		require.Equal(t, 7, iter.Value(), "basic functionality must work")
		iter.Reset()
		assert.False(t, iter.Prev(), "Prev() after Reset() should fail")
		assert.True(t, iter.Next(), "Next() after Reset() should succeed")
		assert.Equal(t, 4, iter.Value(), "after reset first value should be returned after Next()")
	})
	t.Run("skip", func(t *testing.T) {
		iter := iter
		assert.True(t, iter.Skip(2), "Skip(2) from begin should succeed")
		assert.Equal(t, 11, iter.Value(), "Value() after Skip(2) should return third value")
		assert.True(t, iter.Skip(-2), "Skip(-2) from third should succeed")
		assert.Equal(t, 4, iter.Value(), "Value() after Skip(-2) should return first value")
		assert.False(t, iter.Skip(3), "Skip() further than length should fail")
	})
}
