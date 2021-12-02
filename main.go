package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)


func main() {
	fmt.Println("Advent of Code.")
	fmt.Printf("Day One A: %d\n", dayOneA("input1.txt"))
	fmt.Printf("Day One B: %d\n", dayOneB("input1.txt"))
	fmt.Printf("Day Two A: %d\n", dayTwoA("input2.txt"))
	fmt.Printf("Day Two B: %d\n", dayTwoB("input2.txt"))
	fmt.Println("Done.")
}

func dayOneA(filename string) int {
	fmt.Println("~~Day One - A")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var pre, cur, count int
	pre = math.MaxInt32
	count = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cur, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if cur > pre {
			count++
		}
		pre = cur
	}
	return count
}

func dayOneB(filename string) int {
	fmt.Println("~~Day One - B")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var min1, min2, min3, cur, count int
	min1, min2, min3, count = 0, 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cur, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if min3 > 0 {
			if cur > min3 {
				count++
			}
		}
		// Shift
		min3 = min2
		min2 = min1
		min1 = cur
	}
	return count
}

func dayTwoA(filename string) int {
		fmt.Println("~~Day Two - A")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var x, y int
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse
		line := strings.Fields(scanner.Text())
		if len(line) != 2 {
			log.Fatal()
		}
		val, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case "forward":
			x = x + val
		case "down":
			y = y + val
		case "up":
			y = y - val
		default:
			log.Fatal("Unknown command")
		}
	}
	return x*y
}

func dayTwoB(filename string) int {
		fmt.Println("~~Day Two - B")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var x, y, aim int
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse
		line := strings.Fields(scanner.Text())
		if len(line) != 2 {
			log.Fatal()
		}
		val, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case "forward":
			x = x + val
			y = y + aim*val
		case "down":
			aim = aim + val
		case "up":
			aim = aim - val
		default:
			log.Fatal("Unknown command")
		}
	}
	return x*y
}