package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Square struct {
	x int64
	y int64
}

func main() {

	fmt.Println("Day 8")
	start := time.Now()

	dat, err := ioutil.ReadFile("day8.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(dat, []byte("\n"))

	part2 := 0

	totalCharacters := 0
	realCharacters := 0

	for _, value := range lines {
		line := string(value)
		totalCharacters += len(line)

		line = line[1 : len(line)-1]
		splitLine := strings.Split(line, `\`)

		currentReal := 0

		if len(splitLine) == 1 {
			currentReal += len(splitLine[0])
		} else {

			for i := 0; i < len(splitLine); i = i + 2 {
				currentReal += len(splitLine[i])

				if string(splitLine[i+1][0]) == "x" {
					currentReal += 1
				} else if string(splitLine[i+1][0]) == `\` {
					currentReal += 1
				} else if string(splitLine[i+1][0]) == `"` {
					currentReal += 1
				}

			}

		}

		fmt.Println(realCharacters)
		realCharacters += currentReal

		fmt.Println(splitLine)
		fmt.Println()

	}

	fmt.Println(totalCharacters)
	fmt.Println(realCharacters)

	elapsed := time.Since(start)
	fmt.Printf("Part 1: %v %v \n", totalCharacters-realCharacters, elapsed)
	fmt.Printf("Part 2: %v %v \n", part2, elapsed)

}

func compare(slice1 []byte, slice2 []byte) bool {
	for i := 0; i < len(slice1) && i < len(slice2); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
