package check_sort_test

import (
	"testing"

	check_sort "github.com/JlucKant/BalloonsSorter/internal/usecase"
)

type TestCase struct {
	name       string
	containers [][]uint32
	expected   bool
}

func TestCheckSort(t *testing.T) {
	tests := []TestCase{
		{
			name: "Example 1",
			containers: [][]uint32{
				{1, 2},
				{2, 1},
			},
			expected: true,
		},
		{
			name: "Example 2",
			containers: [][]uint32{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
			expected: false,
		},
		{
			name: "Example 3: With only zeros",
			containers: [][]uint32{
				{0, 0},
				{0, 0},
			},
			expected: true,
		},
		{
			name: "Example 4: Already sorted containers",
			containers: [][]uint32{
				{4, 0},
				{0, 4},
			},
			expected: true,
		},
		{
			name: "Example 5: Containers with only one colors",
			containers: [][]uint32{
				{5, 0},
				{5, 0},
			},
			expected: false,
		},
	}

	sorter := check_sort.InitContainerSorter()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := sorter.CheckSort(testCase.containers)
			if result != testCase.expected {
				t.Errorf("Function CheckSort() on %s failed: expected %v, got %v", testCase.name, testCase, result)
			}
		})
	}

}
