package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func get_bounds()(int, int){
	data1, err := ioutil.ReadFile("./2019/day4/day4input")
	check(err)
	data := string(data1)
	var lowerbound int
	var upperbound int

	for ind, val := range data{
		if val == '-'{
			lowerbound, err = strconv.Atoi(data[0:ind])
			upperbound, err = strconv.Atoi(data[ind+1:])
			break
		}
	}
	return lowerbound, upperbound
}

func check_doubles(num string) bool{
	counter := 0
	for i := 0; i <= len(num)-2; i++{
		counter = 0
		if num[i] == num[i+1]{
			counter = 2
			i += 2
			for {
				if i >= len(num){
					break
				}
				if num[i] == num[i-1]{
					counter++
				}else{
					i--
					break
				}
				i++
			}
		}
		if counter == 2{
			return true
		}
	}

	return false
}

func check_increasing(num string) bool{
	increasing := true
	for i := 0; i < len(num)-1; i++{
		if num[i] > num[i+1]{
			increasing = false
		}
	}
	return increasing
}


func validate(val int)bool{
	num := strconv.Itoa(val)
	if check_doubles(num) && check_increasing(num){
		return true
	}
	return false
}

func day4(){
	counter := 0
	lower, upper := get_bounds()
	for i := lower; i < upper; i++{
		if validate(i){
			counter++
		}
	}
	fmt.Print(counter)
}

func main(){
	day4()
}