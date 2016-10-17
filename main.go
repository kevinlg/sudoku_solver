package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*grid := [9][9]int{
		{9, 0, 0, 1, 0, 0, 0, 0, 5},
		{0, 0, 5, 0, 9, 0, 2, 0, 1},
		{8, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 0, 0, 7, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 6, 0, 0, 9},
		{2, 0, 0, 3, 0, 0, 0, 0, 6},
		{0, 0, 0, 2, 0, 0, 9, 0, 0},
		{0, 0, 1, 9, 0, 4, 5, 7, 0},
	}*/

	grid2 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0, 0, 5, 3},
		{0, 0, 0, 3, 6, 4, 0, 0, 0},
		{0, 0, 8, 7, 0, 5, 3, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 2, 4, 0, 9, 1, 0, 0},
		{0, 0, 0, 8, 5, 1, 0, 0, 0},
		{5, 8, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	//simplePut(&grid)
	//isValid(&grid, 0)
	//fmt.Println("After resolution")
	//affiche(grid)

	isValid(&grid2, 0)
	affiche(grid2)
}

func affiche(grille [9][9]int) {
	for i := 0; i < 9; i++ {
		fmt.Println("--------------------")
		line := ""
		for j := 0; j < 9; j++ {
			s := strconv.Itoa(grille[i][j])
			line = line + "|" + s
		}
		line = line + "|"
		fmt.Println(line)
	}
	fmt.Println("---------------------")
}

func simplePut(grille *[9][9]int) {
	grille[4][4] = 122222
	grille[5][5] = 122222
	grille[6][6] = 122222
}

// AbsentSurLigne Check if the number k number is present in the grid
// grill at the line i
func absentSurLigne(k int, grille [9][9]int, i int) bool {
	for j := 0; j < 9; j++ {
		if grille[i][j] == k {
			return false
		}
	}
	return true
}

func absentSurColonne(k int, grille [9][9]int, j int) bool {
	for i := 0; i < 9; i++ {
		if grille[i][j] == k {
			return false
		}
	}
	return true
}

func absentSurBloc(k int, grille [9][9]int, r int, l int) bool {
	_i := r - (r % 3)
	_j := l - (l % 3)
	for i := _i; i < _i+3; i++ {
		for j := _j; j < _j+3; j++ {
			if grille[r][l] == k {
				return false
			}
		}
	}
	return true
}

func isValid(grille *[9][9]int, position int) bool {
	if position == 9*9 {
		return true
	}

	i := position / 9
	j := position % 9

	if grille[i][j] != 0 {
		return isValid(grille, position+1)
	}

	for k := 1; k <= 9; k++ {
		fmt.Println("k", k)
		if absentSurLigne(k, *grille, i) && absentSurColonne(k, *grille, j) && absentSurBloc(k, *grille, i, j) {
			grille[i][j] = k
			if isValid(grille, position+1) {
				return true
			}
		}
	}
	grille[i][j] = 0
	return false

}
