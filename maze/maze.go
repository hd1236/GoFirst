package main

import (
	"fmt"
	"os"
)

/**
广度优先搜索，走迷宫
*/
func main() {
	maze := readFile("maze/maze.in")

	steps := travel(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

	next := point{0, 0}
	start := point{0, 0}
	fmt.Println(&next == &start)
	fmt.Println(next == start)

	nextPlus := pointPlus{point{2, 4}, 5, "hand"}
	startPlus := pointPlus{point{2, 4}, 5, "hand"}
	fmt.Println("pointPlus:", nextPlus == startPlus)
	fmt.Println("pointPlus&:", &nextPlus == &startPlus)

	a := make([]int, 5)
	b := make([]int, 5)
	//fmt.Println("[]int:",a == b)slice不能比较本身
	fmt.Println("[]int]&:", &a == &b)
}

type point struct {
	i, j int
}

//结构体比较，会把结构体里的基础数据做对比
type pointPlus struct {
	p point
	i int
	s string
}

var dirs = []point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(k point) point {
	return point{p.i + k.i, p.j + k.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func travel(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze)) //只是一个slice
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//fmt.Print(steps)

	//开始从start点遍历，每个点可以走上左下右
	Q := make([]point, 0) //需要定义一个队列，表示四个方向
	Q = append(Q, start)
	for len(Q) > 0 {
		var curP = Q[0]
		Q = Q[1:]

		if curP == end {
			break
		}

		for _, dir := range dirs {
			next := curP.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start { //结构体比较可以直接==，会对比结构体内的各个参数是否相等
				continue
			}

			curSteps, _ := curP.at(steps)

			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func readFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var row, colum int
	fmt.Fscanf(file, "%d %d", &row, &colum)
	fmt.Println(row, colum)

	maze := make([][]int, row) //只是一个slice
	for i := 0; i < row; i++ {
		col := make([]int, colum)
		for j := 0; j < colum; j++ {
			fmt.Fscanf(file, "%d", &col[j])
		}
		maze[i] = col
	}

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d ", val) //%3d，3个字符位置，默认右对齐，超过长度不截取
		}
		fmt.Println()
	}
	fmt.Println()

	return maze
}
