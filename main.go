package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"advent-of-code-go/aoc2020"
)

func day25() {
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

	fmt.Printf("Encryption key: %v\n", rfidCtx.CalculateEncryptionKey())
}

func main() {
	day25()
}
