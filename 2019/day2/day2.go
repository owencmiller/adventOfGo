package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func load_data() []int{
	file, err := os.Open("./2019/day2/day2input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range record{
			d, _ := strconv.Atoi(r)
			data = append(data, d)
		}
	}
	return data
}

func opcode_1(pos1, pos2, storage int, data []int){
	data[storage] = data[pos1] + data[pos2]
}

func opcode_2(pos1, pos2, storage int, data []int){
	data[storage] = data[pos1] * data[pos2]
}

func execute_code(data []int) int{
	for i := 0; i < len(data); i+=4{
		if data[i] == 99{
			break
		}else if data[i] == 1{
			opcode_1(data[i+1], data[i+2], data[i+3], data)
		}else if data[i] == 2{
			opcode_2(data[i+1], data[i+2], data[i+3], data)
		}else{
			fmt.Printf("Something went terribly terribly wrong!")
			break;
		}
	}
	return data[0]
}

func day2(){
	data := load_data()

	done := false
	for noun := 0; noun < 100; noun++{
		for verb := 0; verb < 100; verb++{
			test := make([]int, len(data))
			copy(test, data)
			test[1] = noun
			test[2] = verb
			output := execute_code(test)
			if output == 19690720{
				fmt.Printf("Noun: %v, Verb: %v, Ans: %v", noun, verb, (100*noun)+verb)
				done = true
				break;
			}
		}
		if done{
			break
		}
	}

}

func main(){
	day2()
}
