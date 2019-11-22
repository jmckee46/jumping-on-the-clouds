package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	cLength := int32(len(c))
	var jumps int32
	var position int32

	for position < cLength-1 {
		if position+2 < cLength && c[position+2] == 0 {
			jumps++
			position = position + 2
		} else if c[position+1] == 0 {
			jumps++
			position = position + 1
		} else {
			fmt.Println("something is wrong")
		}
	}
	return jumps
}

func main() {

}

// c := []int32{0, 0, 1, 0, 0, 1, 0}
