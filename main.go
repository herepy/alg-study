/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/2/11 10:46
 */

package main

import (
	"algorithm/sort_alg"
	"fmt"
)

func main() {
	arr := []int{
		8, 3, 9, 6, 4, 1, 5, 2, 10, 7,
	}

	result := sort_alg.Insert(arr)

	fmt.Println(result)
}
