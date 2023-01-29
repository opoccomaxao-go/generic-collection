package gmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	// float
	assert.Equal(t, float64(3), Max[float64](0, 1, 2, 3))
	assert.Equal(t, float64(3), Max[float64](3, 2, 1, 0))

	// int
	assert.Equal(t, int(3), Max(0, 1, 2, 3))
	assert.Equal(t, int(3), Max(3, 2, 1, 0))

	// string
	assert.Equal(t, "3", Max("1", "2", "3"))
	assert.Equal(t, "3", Max("3", "2", "1"))

	// single value
	assert.Equal(t, int(1), Max(1))

	// two values
	assert.Equal(t, int(1), Max(1, 0))
}

func TestMin(t *testing.T) {
	// float
	assert.Equal(t, float64(0), Min[float64](0, 1, 2, 3))
	assert.Equal(t, float64(0), Min[float64](3, 2, 1, 0))

	// int
	assert.Equal(t, int(0), Min(0, 1, 2, 3))
	assert.Equal(t, int(0), Min(3, 2, 1, 0))

	// string
	assert.Equal(t, "1", Min("1", "2", "3"))
	assert.Equal(t, "1", Min("3", "2", "1"))

	// single value
	assert.Equal(t, int(1), Min(1))

	// two values
	assert.Equal(t, int(0), Min(1, 0))
}

func TestAbs(t *testing.T) {
	// float
	assert.Equal(t, float64(1), Abs[float64](1))
	assert.Equal(t, float64(1), Abs[float64](-1))

	// int
	assert.Equal(t, int(1), Abs(1))
	assert.Equal(t, int(1), Abs(-1))
}

func TestSum(t *testing.T) {
	assert.Equal(t, int(15), Sum(1, 2, 3, 4, 5))
	assert.Equal(t, float64(15), Sum[float64](1, 2, 3, 4, 5))
}
