package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*var grid [9][9] int
	print_grid(grid)*/

	//creating a 2D array for the grid
	//grid=[[0 for x in range(9)]for y in range(9)]

	//assigning values to the grid
	grid := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0}}

	//if sucess print the grid
	if solve_sudoku(grid) {
		print_grid(grid)
	} else {
		//print "No solution exists"
		fmt.Println("No solution exists")
	}

}

func print_grid(arr [][]int) {
	i := 0
	j := 0
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			fmt.Printf(strconv.Itoa(arr[i][j]))
		}
		fmt.Println()
	}
}

func find_empty_location(arr [][]int, list []int) bool {
	i := 0
	j := 0
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if arr[i][j] == 0 {
				list[0] = i
				list[1] = j
				return true
			}
		}

	}
	return false
}

func used_in_row(arr [][]int, row int, num int) bool {
	j := 0
	for j = 0; j < 9; j++ {
		if arr[row][j] == num {
			return true
		}
	}
	return false
}

func used_in_col(arr [][]int, col int, num int) bool {
	i := 0
	for i = 0; i < 9; i++ {
		if arr[i][col] == num {
			return true
		}
	}
	return false
}

func used_in_box(arr [][]int, row int, col int, num int) bool {
	i := 0
	j := 3
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if arr[row+i][col+j] == num {
				return true
			}
		}
	}
	return false
}

func check_location_is_safe(arr [][]int, row int, col int, num int) bool {
	return !used_in_row(arr, row, num) && !used_in_col(arr, col, num) && !used_in_box(arr, row-row%3, col-col%3, num)
}

func solve_sudoku(arr [][]int) bool {

	//keep track of row and col in find_empty_location function
	lst := []int{0, 0}

	//if there is no empty location, we are finished
	if !find_empty_location(arr, lst) {
		return true
	}

	//Assigning list values to row and col that we got from the above Function
	row := lst[0]
	col := lst[1]

	num := 1
	for num = 1; num < 10; num++ {
		if check_location_is_safe(arr, row, col, num) {
			arr[row][col] = num

			if solve_sudoku(arr) {
				return true
			}

			arr[row][col] = 0
		}
	}
	return false

}
