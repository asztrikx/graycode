package main

import (
	"fmt"
	"log"
	"sort"
)

func grayFill(array []uint, bitindex int, up bool) {
	if bitindex == -1 {
		return
	}

	top := array[0 : len(array)/2]
	bottom := array[len(array)/2:]

	var fill []uint
	if up {
		fill = top
	} else {
		fill = bottom
	}

	for i := range fill {
		fill[i] |= (1 << bitindex)
	}

	grayFill(top, bitindex-1, false)
	grayFill(bottom, bitindex-1, true)
}

func grayCreate(bitwidth int) []uint {
	array := make([]uint, 1<<(bitwidth))
	grayFill(array, bitwidth-1, false)

	return array
}

func permutationNext(array []int) bool {
	greaterIndexes := []int{
		len(array) - 1,
	}
	max := array[greaterIndexes[0]]

	//find first non increasing
	for i := len(array) - 1 - 1; i >= 0; i-- {
		if array[i] > max {
			max = array[i]
			greaterIndexes = append(greaterIndexes, i)
		} else if array[i] < max {
			maxMinIndex := len(greaterIndexes) - 1
			for maxMinIndex >= 0 && array[greaterIndexes[maxMinIndex]] > array[i] {
				maxMinIndex--
			}
			maxMinIndex++

			array[i], array[greaterIndexes[maxMinIndex]] = array[greaterIndexes[maxMinIndex]], array[i]
			sort.Slice(array[i+1:], func(index1, index2 int) bool {
				return array[i+1+index1] < array[i+1+index2]
			})

			return true
		} else {
			log.Fatal("same value in permutation")
		}
	}

	//all increasing
	return false
}

func main() {
	digit := []uint{
		0b000, //0
		0b011, //3
		0b110, //6
		0b010, //2
		0b100, //4
		0b111, //7
		0b001, //1
		0b101, //5
	}

	bitwidth := 3
	array := grayCreate(bitwidth)
	fmt.Println("Default gray code:")
	for _, elem := range array {
		for i := bitwidth - 1; i >= 0; i-- {
			if (elem & (1 << i)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	//gray code does not change if we rotate rows
	min := -1
	for rowOffset := 0; rowOffset < len(array); rowOffset++ {
		//gray code does not change if we swap columns
		var colIndexes []int
		for i := 0; i < bitwidth; i++ {
			colIndexes = append(colIndexes, i)
		}

		for {
			//diff between digit and gray by bit
			diff := 0

			//walk through applying rotate and swap
			for row := 0; row < len(array); row++ {
				for i, colIndex := range colIndexes {
					//gray bit value
					iszeroGray := (array[(rowOffset+row)%len(array)] & (1 << (bitwidth - 1 - colIndex))) == 0
					if iszeroGray {
						fmt.Print("0")
					} else {
						fmt.Print("1")
					}

					//digit bit value
					iszeroDigit := (digit[row] & (1 << (bitwidth - 1 - i))) == 0

					//diff
					if iszeroGray != iszeroDigit {
						diff++
					}
				}
				fmt.Println()
			}

			fmt.Printf("%d diff\n", diff)
			fmt.Println()

			//min
			if min == -1 || min > diff {
				min = diff
			}

			if !permutationNext(colIndexes) {
				break
			}
		}

		fmt.Println()
	}

	fmt.Printf("Min diff: %d\n", min)
}
