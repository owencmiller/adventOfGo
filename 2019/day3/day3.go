package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type point struct{
	x, y int
}

type line struct{
	start, end *point
}

func New_point(x, y int) *point{
	p := point{x:x, y:y}
	return &p
}

func New_line(start, end *point) *line{
	l := line{
		start: start,
		end:   end,
	}
	return &l
}

func (l *line) is_vertical() bool{
	if l.start.x == l.end.x{
		return true
	}
	return false
}

func (l *line) is_horizontal() bool{
	if l.start.y == l.end.y{
		return true
	}
	return false
}

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}


func load_input() ([]string, []string){
	file, err := os.Open("./2019/day3/day3input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data, data2 []string

	r := csv.NewReader(file)
	record, err := r.Read()
	data = record
	record, err = r.Read()
	data2 = record

	if err != nil {
		//log.Fatal(err)
	}
	return data, data2
}

func execute_move(p *point, move string) *point{
	steps, err := strconv.Atoi(move[1:])
	if err != nil{
		panic(err)
	}
	var new_p *point
	switch move[0]{
	case 'U':
		new_p = New_point(p.x, p.y + steps)
	case 'D':
		new_p = New_point(p.x, p.y - steps)
	case 'R':
		new_p = New_point(p.x + steps, p.y)
	case 'L':
		new_p = New_point(p.x-steps, p.y)
	default:
		new_p = New_point(0,0)
	}
	return new_p
}

func in_between(val, end1, end2 int) bool{
	if (val <= end1 && val >= end2) || (val >= end1 && val <= end2){
		return true
	}
	return false
}

func distance(point1, point2 *point) int{
	return int(math.Abs(float64(point1.y-point2.y)) + math.Abs(float64(point1.x-point2.x)))
}

func intersection(line1, line2 *line) *point{
	var intersect *point
	if line1.is_vertical(){
		if line2.is_vertical(){
			return nil
		}else if line2.is_horizontal(){
			if in_between(line1.start.x, line2.start.x, line2.end.x){
				if in_between(line2.start.y, line1.start.y, line1.end.y){
					intersect = New_point(line1.start.x, line2.start.y)
					return intersect
				}
			}
		}else{
			panic("Line does not follow taxi-cab geometry!!!")
		}
	}else if line1.is_horizontal(){
		if line2.is_horizontal(){
			return nil
		}else if line2.is_vertical(){
			if in_between(line2.start.x, line1.start.x, line1.end.x){
				if in_between(line1.start.y, line2.start.y, line2.end.y){
					intersect = New_point(line2.start.x, line1.start.y)
					return intersect
				}
			}
		}else{
			panic("Line does not follow taxi-cab geometry!!!")
		}
	}else{
		panic("Line does not follow taxi-cab geometry!!!")
	}
	return nil
}

func day3(){
	wire1, wire2 := load_input()
	intersections := make([]*point, 0)
	start := New_point(0,0)
	wire1_segment := New_line(nil, start)
	wire1_counter := 0
	var min_distance_sum int

	for i := 0; i < len(wire1); i++{
		wire1_segment = New_line(wire1_segment.end, execute_move(wire1_segment.end, wire1[i]))

		wire2_segement := New_line(nil, start)
		wire2_counter := 0
		for j := 0; j < len(wire2); j++{
			wire2_segement = New_line(wire2_segement.end, execute_move(wire2_segement.end, wire2[j]))

			intersect := intersection(wire1_segment, wire2_segement)
			if intersect != nil{
				if intersect.y == 0 && intersect.x == 0{
					continue
				}
				// Part 2
 				if min_distance_sum == 0 {
 					fmt.Print(wire1_counter, ":", wire2_counter, ":", distance(wire1_segment.start, intersect), ":", distance(wire2_segement.start, intersect), "\n")
					min_distance_sum = distance(wire2_segement.start, intersect) + distance(wire1_segment.start, intersect) + wire1_counter + wire2_counter
				}

				// Part One
				intersections = append(intersections, intersect)
			}

			wire2_temp, err := strconv.Atoi(wire2[j][1:])
			check(err)
			wire2_counter += wire2_temp
		}

		wire1_temp, err := strconv.Atoi(wire1[i][1:])
		check(err)
		wire1_counter += wire1_temp
	}

	// Part One
	intersections = append(intersections[:0], intersections[0+1:]...)
	min := math.Abs(float64(intersections[0].y)) + math.Abs(float64(intersections[0].x))
	for _, i := range intersections{
		dist := math.Abs(float64(i.y)) + math.Abs(float64(i.x))
		if dist < min{
			min = dist
		}
	}
	fmt.Printf("Part 1: %v\n" , min)
	fmt.Printf("Part 2: %v\n" , min_distance_sum)
}


func main(){
	day3()
}
