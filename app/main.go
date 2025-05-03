package main

import (
	"BalloonsClassifier/internal/check_sort"
	"BalloonsClassifier/internal/parser"
	"fmt"
	"os"
)

func main() {
	containers, err := parser.InputParser(os.Stdin)

	if err != nil {
		fmt.Println("Input Error. Failed input parsing: %s", err)
	}

	sorter := &check_sort.ContainerSorterByColor{}

	isSortAvailable := sorter.CheckSort(sorter, containers)

	if isSortAvailable {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
