package main

import (
	"fmt"
	"strconv"
)

func main() {
/*	i:= 1
	for i=1; i <100 ;i++  {
		if i % 3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i % 5 == 0 {
			fmt.Println("Buzz")
			continue
		} else {
		    fmt.Println(strconv.Itoa(i))
		}

	}*/
	var grid [9][9] int
	print_grid(grid)
}

func print_grid( arr [9][9] int)  {
	i := 0
	j := 0
	for i = 0;i < 9 ;  i++{
		for j = 0; j < 9 ;j++  {
			fmt.Printf(strconv.Itoa( arr[i][j] ))
		}
		fmt.Println()
	}
}

func find_empty_location(arr [9][9] int , list [2] int) bool {
	i := 0
	j := 0
	for i = 0;i < 9 ;  i++{
		for j = 0; j < 9 ;j++  {
			if arr[i][j] == 0 {
				list[0] = i
				list[1] = j
				return true
			}
		}

	}
	return false
}

func used_in_row(arr [9][9] int, row int, num int) bool {
	j:= 0
	for j = 0; j < 9 ;j++  {
		if arr[row][j] == num {
			return true
		}
	}
	return false
}

func used_in_col(arr [9][9] int, col int, num int) bool {
	i:= 0
	for i = 0; i < 9 ;i++  {
		if arr[i][col] == num {
			return true
		}
	}
	return false
}

func used_in_box(arr [9][9] int, row int, col int, num int) bool {
	i:= 0
	j:=3
	for i = 0; i < 3 ;i++  {
		for j = 0; j < 3 ;j++  {
			if arr[row + i][col + j] == num {
				return true
			}
		}
	}
	return false
}

func check_location_is_safe(arr [9][9] int, row int, col int, num int) bool {
	return !used_in_row(arr,row,num) && !used_in_col(arr,col,num) && !used_in_box(arr,row,col,num)
}

func solve_sudoku(arr[9][9] int) bool {

	//keep track of row and col in find_empty_location function
	lst := [2]int{0, 0}

	//if there is no empty location, we are finished
	if !find_empty_location(arr, lst) {
		return true
	}

	//Assigning list values to row and col that we got from the above Function
	row := lst[0]
	col := lst[1]

	num := 1
	for num = 1; num < 10 ; num++  {
		if check_location_is_safe(arr, row, col, num) {
			arr[row][col] = num

			if(solve_sudoku(arr)) {
				return true
			}

			arr[row][col] = 0
		}
	}
	return  false

}