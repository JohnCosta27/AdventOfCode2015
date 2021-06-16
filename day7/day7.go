package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Day 7")
	start := time.Now()

	dat, err := ioutil.ReadFile("day7.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(dat, []byte("\n"))
	values := make(map[string][]byte)

	for _, line := range lines {

		split := strings.Split(string(line), "->")
		split[0] = split[0][:len(split[0])-1]
		split[1] = split[1][1:len(split[1])]

		if len(values[split[1]]) == 0 {
			values[split[1]] = []byte{0, 0}
		}

		instruction := strings.Fields(split[0])

		if len(instruction) == 1 {
			byteNumber := make([]byte, 2)
			intNum, _ := strconv.Atoi(instruction[0])
			binary.BigEndian.PutUint16(byteNumber, uint16(intNum))
			values[split[1]] = byteNumber
		} else if instruction[0] == "NOT" {
			if len(values[instruction[1]]) == 0 {
				values[instruction[1]] = []byte{0, 0}
			}
			values[split[1]] = []byte{^values[instruction[1]][0], ^values[instruction[1]][1]}
		} else {

			if len(values[instruction[0]]) == 0 {
				values[instruction[0]] = []byte{0, 0}
			}
			if len(values[instruction[2]]) == 0 {
				values[instruction[2]] = []byte{0, 0}
			}

			switch instruction[1] {
			case "AND":
				values[split[1]] = []byte{values[instruction[0]][0] & values[instruction[2]][0], values[instruction[0]][1] & values[instruction[2]][1]}
			case "OR":
				values[split[1]] = []byte{values[instruction[0]][0] | values[instruction[2]][0], values[instruction[0]][1] | values[instruction[2]][1]}
			case "LSHIFT":
				expo, _ := strconv.Atoi(instruction[2])
				num := binary.BigEndian.Uint16(values[instruction[0]])
				byteNum := make([]byte, 2)
				binary.BigEndian.PutUint16(byteNum, uint16(num)<<expo)
				values[split[1]] = byteNum
			case "RSHIFT":
				expo, _ := strconv.Atoi(instruction[2])
				num := binary.BigEndian.Uint16(values[instruction[0]])
				byteNum := make([]byte, 2)
				binary.BigEndian.PutUint16(byteNum, uint16(num)>>expo)
				values[split[1]] = byteNum
			}
		}

	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1: %v %v", values["a"], elapsed)

}

func compare(slice1 []byte, slice2 []byte) bool {
	for i := 0; i < len(slice1) && i < len(slice2); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
