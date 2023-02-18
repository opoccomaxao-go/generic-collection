package slice

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	res := Map([]int{1, 2, 3, 4, 5}, strconv.Itoa)

	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, res)
}

func TestFilter(t *testing.T) {
	res := Filter(
		[]int{1, 2, 3, 4, 5},
		func(item int) bool { return item&1 == 0 },
	)

	assert.Equal(t, []int{2, 4}, res)
}

func TestJoin(t *testing.T) {
	res := Join([]int{1, 2, 3, 4, 5}, ":", strconv.Itoa)

	assert.Equal(t, "1:2:3:4:5", res)
}

func TestReduce(t *testing.T) {
	sumFunc := func(prev int, cur int) int { return prev + cur }

	t.Run("int result", func(t *testing.T) {
		res := Reduce([]int{1, 2, 3, 4, 5}, sumFunc, 0)

		assert.Equal(t, 15, res)
	})

	mapFunc := func(prev map[int]int, cur int) map[int]int {
		prev[cur] = cur

		return prev
	}

	t.Run("map result", func(t *testing.T) {
		res := Reduce([]int{1, 2, 3, 4, 5}, mapFunc, map[int]int{})

		assert.Equal(t, map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
			5: 5,
		}, res)
	})
}

func TestUniqueUnordered(t *testing.T) {
	res := UniqueUnordered([]int{1, 2, 3, 4, 5, 3, 4, 5})

	sort.Ints(res)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
}

func TestFindIndex(t *testing.T) {
	res := FindIndex([]int{1, 2, 3, 4, 5, 3, 4, 5}, func(i int) bool {
		return i == 5
	})

	assert.Equal(t, 4, res)
}

func TestFind(t *testing.T) {
	res := Find([]int{1, 2, 3, 4, 5, 3, 4, 5}, func(i int) bool {
		return i > 4
	})

	assert.Equal(t, 5, res)
}

func TestFirstNonEmpty(t *testing.T) {
	assert.Equal(t, 1, FirstNonEmpty(0, 0, 0, 1, 2, 4, 5))

	assert.Equal(t, "1", FirstNonEmpty("", "", "1"))
}

func TestChunk(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testCases := []struct {
		size   int
		output [][]int
	}{
		{
			size: 10,
			output: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
		{
			size: 5,
			output: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
			},
		},
		{
			size: 4,
			output: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10},
			},
		},
		{
			size: 3,
			output: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10},
			},
		},
		{
			size: 2,
			output: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
				{9, 10},
			},
		},
		{
			size: 1,
			output: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
				{6},
				{7},
				{8},
				{9},
				{10},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(strconv.Itoa(tC.size), func(t *testing.T) {
			output := Chunk(slice, tC.size)
			assert.Equal(t, tC.output, output)
		})
	}
}

func TestConcat(t *testing.T) {
	assert.Equal(t,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Concat(
			[]int{1, 2, 3},
			[]int{4, 5},
			[]int{6},
			[]int{7, 8, 9},
		),
	)
}
