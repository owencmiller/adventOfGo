package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}

type Instruction struct{
	position int
	code int
	params map[int]int
}

func create_instruction(opcode int, position int) (*Instruction){
	code := opcode % 10
	opcode /= 10
	code = code + (10 * (opcode % 10))
	opcode /= 10
	mode1 := opcode % 10
	opcode /= 10
	mode2 := opcode % 10
	opcode /= 10
	mode3 := opcode % 10

	params := map[int]int{
		1: mode1,
		2: mode2,
		3: mode3,
	}

	instruction := Instruction{
		position: position,
		code:  code,
		params: params,
	}

	return &instruction
}


func opcode_1(in *Instruction, data []int, in_pointer int)int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	p2 := get_param_value(in, 2, data[in.position+2], data)
	p3 := data[in.position+3]

	data[p3] = p1 + p2

	return in_pointer + 4
}

func opcode_2(in *Instruction, data []int, in_pointer int)int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	p2 := get_param_value(in, 2, data[in.position+2], data)
	p3 := data[in.position+3]

	data[p3] = p1 * p2

	return in_pointer + 4
}

func opcode_3(in *Instruction, data []int, in_pointer int) int {
	var input int
	fmt.Print("Input required: ")
	_, err := fmt.Scanf("%d", &input)
	check(err)

	p1 := data[in.position+1]
	data[p1] = input

	return in_pointer + 2
}

func opcode_4(in *Instruction, data []int, in_pointer int) int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	fmt.Print(p1)
	return in_pointer + 2
}

func opcode_5(in *Instruction, data []int, in_pointer int)int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	p2 := get_param_value(in, 2, data[in.position+2], data)

	if p1 != 0{
		return p2
	}

	return in_pointer + 3
}

func opcode_6(in *Instruction, data[]int, in_pointer int)int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	p2 := get_param_value(in, 2, data[in.position+2], data)

	if p1 == 0{
		return p2
	}

	return in_pointer + 3
}

func opcode_7(in *Instruction, data[]int, in_pointer int)int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	p2 := get_param_value(in, 2, data[in.position+2], data)
	p3 := data[in.position+3]
	if p1 < p2{
		data[p3] = 1
	}else{
		data[p3] = 0
	}
	return in_pointer + 4
}

func opcode_8(in *Instruction, data[]int, in_pointer int)int{
	p1 := get_param_value(in, 1, data[in.position+1], data)
	p2 := get_param_value(in, 2, data[in.position+2], data)
	p3 := data[in.position+3]
	if p1 == p2{
		data[p3] = 1
	}else{
		data[p3] = 0
	}
	return in_pointer + 4
}


func get_param_value(in *Instruction, param_pos int, param_val int, data[]int)int{
	switch in.params[param_pos]{
	case 0:
		return data[param_val]
	case 1:
		return param_val
	default:
		fmt.Print("INVALID POSITION MODE")
	}
	return 0
}

func execute_instruction(in *Instruction, data []int, in_pointer int) int{
	switch in.code{
	case 1:
		return opcode_1(in, data, in_pointer)
	case 2:
		return opcode_2(in, data, in_pointer)
	case 3:
		return opcode_3(in, data, in_pointer)
	case 4:
		return opcode_4(in, data, in_pointer)
	case 5:
		return opcode_5(in, data, in_pointer)
	case 6:
		return opcode_6(in, data, in_pointer)
	case 7:
		return opcode_7(in, data, in_pointer)
	case 8:
		return opcode_8(in, data, in_pointer)
	case 99:
		return -1
	}
	return -1
}


func execute_code(data []int){
	for i := 0; i < len(data);{
		in := create_instruction(data[i], i)
		i = execute_instruction(in, data, i)
		if i == -1 {
			break
		}
	}
}

func day5(){
	data := load_data()

	execute_code(data)

}

func main(){
	day5()
}

func load_data() []int{
	file, err := os.Open("./2019/day5/day5input")
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