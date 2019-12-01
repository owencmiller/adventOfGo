package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calc_fuel(mass float64) int{
	return int(math.Floor(mass/3) - 2)
}

func calc_part2(mass float64) int{
	sum := 0
	temp := mass
	for{
		temp = float64(calc_fuel(temp))
		if temp <= 0{
			break;
		}else{
			sum += int(temp)
		}
	}
	return sum
}

func day1(){
	file, err := os.Open("./2019/day1/day1input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		val, e := strconv.Atoi(scanner.Text())
		sum += calc_part2(float64(val))
		if e != nil{
			log.Fatal(e)
		}
	}

	fmt.Printf("Total Sum: %v ", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func main(){
	day1()
}
