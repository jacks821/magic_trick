package main

import (
	"strings"
	"fmt"
	"strconv"
	"os"
	"bufio"
	"log"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Card(i int, firstrow []int, secondrow []int) {
	fmt.Printf("Case #%d: ", i)
	var r []int
	for _, num1 := range firstrow {
		for _, num2 := range secondrow {
			if num1 == num2 {
				r = append(r, num1)
			}
		}
	}
	if len(r) == 1 {
		fmt.Println(r[0])
	} else if len(r) == 0 {
		fmt.Println("Volunteer cheated!")
	} else {
		fmt.Println("Bad magician!")
	}
}

func main() {
	argsWithoutProgram := os.Args[1]
	lines := GrabLines(argsWithoutProgram)
	cases,_ := strconv.Atoi(lines[0])
	index := 1
	for i := 1; i <= cases; i++{
		firstindex,_ := strconv.Atoi(lines[index])
		var firstrow []int
		var secondrow []int
		row := strings.Split(lines[index+firstindex], " ")
		firstrow = GetRow(row)
		index += 5
		secondindex,_ := strconv.Atoi(lines[index])
		j := strings.Split(lines[index+secondindex], " ")
		secondrow = GetRow(j)
		Card(i, firstrow, secondrow)
		index += 5
	}
}

func GetRow(row []string) []int{
	var firstrow []int
	for _, num := range row {
		intnum, _ := strconv.Atoi(num)
		firstrow = append(firstrow, intnum)
	}
	return firstrow
}
