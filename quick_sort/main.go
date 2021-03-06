package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快速排序实现，由于要二分递归，所以要传入每次排序的起始和结束的索引
func quickSort(list []int, start, end int) {
	if end-start > 1 {
		// get the pivot
		mid := partition(list, start, end)
		// left sort
		quickSort(list, start, mid)

		// right sort
		quickSort(list, mid+1, end)
	}
}

func partition(list []int, begin, end int) (i int) {
	fmt.Println("begin:", begin, "end:", end, "list:", list)
	cValue := list[begin]
	i = begin
	for j := i + 1; j < end; j++ {
		if list[j] < cValue {
			list[j], list[i] = list[i], list[j]
			i++ // 这里一定要先加1后在交换值，因为支点现在不必交换，现在i 和 j（小于支点和大于支点）正在划分边界
		}
	}
	/* 到此，i和j的边界已经划分完成，把i对应的值和支点对应的值交换后，就符合了快分的要求：i左边对应的值都小于等于且右边的都大于支点对应的值
	   此时i的值就是新的支点, 对应的值就是新的主元
	*/
	//fmt.Println("i:", i, "list:", list)
	//list[i], list[begin] = list[begin], list[i]

	return i
}
func main() {
	rand.Seed(time.Now().Unix())
	data := rand.Perm(20)
	fmt.Println(data)
	quickSort(data, 0, len(data))
	//fmt.Println(data)
}
