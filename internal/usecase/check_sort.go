package check_sort

import (
	"slices"
)

type ContainerSorterByColor struct{}

func InitContainerSorter() SortChecker {
	return &ContainerSorterByColor{}
}

func (containerSorter *ContainerSorterByColor) sorterBySums(containers [][]uint32) ([]uint32, []uint32) {
	var (
		containerLength = len(containers)

		containerBalloonsSums = make([]uint32, containerLength)
		colorBalloonsSums     = make([]uint32, containerLength)
	)

	for containerIndex := 0; containerIndex < containerLength; containerIndex++ {
		for colorIndex := 0; colorIndex < containerLength; colorIndex++ {
			balloon := containers[containerIndex][colorIndex]

			containerBalloonsSums[containerIndex] += balloon
			colorBalloonsSums[colorIndex] += balloon
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
