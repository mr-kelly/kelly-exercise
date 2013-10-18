// 插入排序
package sort

type _T interface{}

// 插入排序, 轮个替换
// 默认 {34, 8, 64, 51, 32, 21}
// 第1趟 {8, 34, 64, 51, 32, 21}
// 第2趟 {8, 34, 64, 51, 32, 21}
// 第3趟 {8, 34, 51, 64, 32, 21}
// 第4趟 {8, 34, 32, 51, 64, 21}
// 第5趟 {8, 21, 32, 34, 51, 64}
func InsertionSort(slice *[]_T) {
	// 排序p趟，每趟前面都要有序
	// foreach()
	var j int = 0

	for i := 1; i < len(*slice); i++ {
		var tmp _T = (*slice)[i] // 第1趟   // 保证位置0到i，是顺序的

		for j = i; j > 0 && tmp.(int) < ((*slice)[j-1]).(int); j-- { // 如果遇到j比j-1小，交换，否则到下一阶段
			// swap , 获取, 对比，小于，则交换
			(*slice)[j] = (*slice)[j-1]
		}

		(*slice)[j] = tmp // j = 0
	}
}
