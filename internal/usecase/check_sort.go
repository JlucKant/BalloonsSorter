package check_sort

import (
	"slices"

	"github.com/JlucKant/BalloonsSorter/domain"
)

type ContainerSorterByColor struct{}

func InitContainerSorter() domain.SortChecker {
	return &ContainerSorterByColor{}
}

func (containerSorter *ContainerSorterByColor) sorterBySums(containers [][]uint32) ([]uint, []uint) {
	var (
		containerLength = len(containers)

		containerBalloonsSums = make([]uint, containerLength)
		colorBalloonsSums     = make([]uint, containerLength)
	)

	for containerIndex := 0; containerIndex < containerLength; containerIndex++ {
		for colorIndex := 0; colorIndex < containerLength; colorIndex++ {
			balloon := containers[containerIndex][colorIndex]

			containerBalloonsSums[containerIndex] += uint(balloon)
			colorBalloonsSums[colorIndex] += uint(balloon)
		}
	}

	slices.Sort(containerBalloonsSums)
	slices.Sort(colorBalloonsSums)

	return containerBalloonsSums, colorBalloonsSums
}

func (containerSorter *ContainerSorterByColor) CheckSort(containers [][]uint32) bool {
	var (
		containerSortedByColors, containerSortedByBalloons = containerSorter.sorterBySums(containers)

		lengthOfSortedContainers = len(containers)
	)

	for indexOfSortedSlice := 0; indexOfSortedSlice < lengthOfSortedContainers; indexOfSortedSlice++ {
		if containerSortedByBalloons[indexOfSortedSlice] != containerSortedByColors[indexOfSortedSlice] {
			return false
		}
	}

	return true
}
