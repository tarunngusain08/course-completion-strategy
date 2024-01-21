package main

import (
	"errors"
	"fmt"
)

func findOrder(n int, dependencies [][]int) (*[]int, error) {
	inDegrees := make([]int, n)
	adjacencyLists := make(map[int][]int)

	for _, dependency := range dependencies {
		course, prerequisite := dependency[0], dependency[1]
		inDegrees[course]++
		adjacencyLists[prerequisite] = append(adjacencyLists[prerequisite], course)
	}

	queue := make([]int, 0)

	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range adjacencyLists[node] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != n {
		return nil, errors.New("courses cannot be completed due to cycle dependency")
	}

	return &result, nil
}

func main() {
	ans, err := findOrder(2, [][]int{{1, 0}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*ans)
	}

	ans, err = findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*ans)
	}
	ans, err = findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {1, 2}, {1, 3}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*ans)
	}
	ans, err = findOrder(1, [][]int{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*ans)
	}
}
