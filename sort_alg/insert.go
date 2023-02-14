/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/2/11 17:32
 */

package sort_alg


func Insert(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	//外循环是遍历未排序的元素
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		//内循环遍历已排序元素，找到合适位置插入当前未排序元素
		for j := i-1; j >= 0; j-- {
			//当前未排序元素比当前已排序元素小，则以类似冒泡算法方式，在已排序元素中找到当前未排序元素应该放入的位置（一直对比交换过去）
			if tmp < arr[j] {
				swap(arr, j, j+1)
			} else {
				break
			}
		}
	}

	return arr
}
