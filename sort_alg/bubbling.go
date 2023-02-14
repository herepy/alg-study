/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/2/11 15:37
 */

package sort_alg

//8, 3, 9, 6, 4, 1, 5, 2, 10, 7

func Bubble(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	//外层循环确定第x大（小）的数
	for i := range arr[:len-1] {
		//内层循环找到本轮未排序中最大（小）的数
		for j := 0; j < len - i - 1; j++ {
			//当前数比下一个数大（小）则交换位置（冒泡）
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}
