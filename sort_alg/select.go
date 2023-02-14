/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/2/11 10:36
 */

package sort_alg


func Select(arr []int) []int {
	//元素少不用排
	if len(arr) <= 1 {
		return arr
	}

	for i := range arr {
		//默认假设该轮循环的第一个数为最小数
		minIndex := i
		//循环未排序元素，获取未排序的最小值下标
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//交换该轮的最小值和未排序的第一个元素
		swap(arr, i, minIndex)
	}

	return arr
}