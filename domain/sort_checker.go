package domain

type SortChecker interface {
	CheckSort([][]uint32) bool
}
