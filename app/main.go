package main

import (
	"fmt"
	"os"

	parser "github.com/JlucKant/BalloonsSorter/internal/interface"
	check_sort "github.com/JlucKant/BalloonsSorter/internal/usecase"
)

func main() {
	containers, err := parser.InputParser(os.Stdin)

	if err != nil {
		fmt.Println("Parse Error:", err)
		os.Exit(1)
	}

	sorter := check_sort.InitContainerSorter()

	isSortAvailable := sorter.CheckSort(containers)

	if isSortAvailable {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
