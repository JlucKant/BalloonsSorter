package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func InputParser(reader io.Reader) ([][]uint32, error) {
	scanner := bufio.NewScanner(reader)

	if isScan := scanner.Scan(); !isScan {
		return nil, fmt.Errorf("Missing size value!")
	}

	containerSize, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return nil, fmt.Errorf("Incorrect type of size value!")
	}

	containers := make([][]uint32, containerSize)

	for countOfContainers := 0; countOfContainers < containerSize; countOfContainers++ {
		scanner.Scan()

		containerColors := strings.Fields(scanner.Text())

		if len(containerColors) != containerSize {
			return nil, fmt.Errorf("The number of entered values does not correspond to the number of colors")
		}

		container := make([]uint32, containerSize)
		for indexOfColor, coloredBalloons := range containerColors {
			coloredBalloonsAmount, err := strconv.Atoi(coloredBalloons)

			if err != nil {
				return nil, fmt.Errorf("Incorrect type of amount of balloons in container!")
			}

			if coloredBalloonsAmount < 0 {
				return nil, fmt.Errorf("The amount of balloons must be positive!")
			}

			container[indexOfColor] = uint32(coloredBalloonsAmount)
		}

		containers[countOfContainers] = container
	}

	return containers, nil
}
