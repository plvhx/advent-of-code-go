package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"advent-of-code-go/aoc2020"
)

func day1_part1() {
	f, err := os.Open("./input/day1.txt")

	if err != nil {
		fmt.Errorf("[Error] %s\n", err)
		return
	}

	defer f.Close()

	arr := make([]uint32, 0)
	reader := bufio.NewReader(f)

	for {
		lbuf, _, err := reader.ReadLine()

		if err != nil || len(lbuf) == 0 {
			break
		}

		num, err := strconv.Atoi(string(lbuf))

		if err != nil {
			break
		}

		arr = append(arr, uint32(num))
	}

	fmt.Printf("[Day 1, Part 1] multiplied: %v\n", aoc2020.ReportRepair(arr))
}

func day19() {
	f, err := os.Open("./input/day19.txt")

	if err != nil {
		fmt.Errorf("[Error] %s\n", err)
		return
	}

	defer f.Close()

	msgCtx := aoc2020.NewMessage()
	reader := bufio.NewReader(f)

	for {
		lbuf, _, err := reader.ReadLine()

		if err != nil || len(lbuf) == 0 {
			break
		}

		msgCtx.BuildRuleTable(string(lbuf))
	}

	msgCtx.TopDownMatchTraversal("gandung", 0)
}

func day25_part1() {
	f, err := os.Open("./input/day25.txt")

	if err != nil {
		fmt.Errorf("[Error] %s\n", err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	cardPkeyBuf, _, _ := reader.ReadLine()
	doorPkeyBuf, _, _ := reader.ReadLine()

	cardPkey, _ := strconv.Atoi(string(cardPkeyBuf))
	doorPkey, _ := strconv.Atoi(string(doorPkeyBuf))
	rfidCtx := aoc2020.NewRfidContext(uint32(cardPkey), uint32(doorPkey))

	fmt.Printf("[Day 25, Part 1] Encryption key: %v\n", rfidCtx.CalculateEncryptionKey())
}

func main() {
	day1_part1()
	day19()
	day25_part1()
}
