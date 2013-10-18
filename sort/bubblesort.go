package sort

// 冒泡排序
// 默认 {34, 8, 64, 51, 32, 21}
// 第1.1趟 {34, 8, 64, 51, 21, 32}
// 第1.2趟 {34, 8, 64, 21, 51, 32}
// 第1.3趟 {34, 8, 21, 64, 51, 32}
// 第1.5趟 {8, 34, 21, 64, 51, 32}
// 第2.1趟 {8, 32, 64, 51, 21, 34}
// 第3.1趟 {8, 32, 64, 51, 21, 34}
// ...
func BubbleSort(slice *[]int) {
	for i := 0; i < len(*slice)-1; i++ { // 从小到大  共走5趟，最后一个不走
		for j := len(*slice) - 1; j > 0; j-- { //开始冒泡，遇到比自己小的交换
			if (*slice)[j] < (*slice)[j-1] {
				var tmp int = (*slice)[j]
				(*slice)[j] = (*slice)[j-1]
				(*slice)[j-1] = tmp
			}
		}
	}
}
