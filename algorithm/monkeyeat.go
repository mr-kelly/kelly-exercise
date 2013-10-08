// 猴子分食桃子 : 递推法

// 五只猴子采得一堆桃子，猴子彼此约定隔天早起后再分食。
// 不过，就在半夜里，一只猴子偷偷起来，把桃子均分成五堆后，发现还多一个，它吃掉这桃子，并拿走了其中一堆。
// 第二只猴子醒来，又把桃子均分成五堆后，还是多了一个，它也吃掉这个桃子，并拿走了其中一堆。
// 第三只，第四只，第五只猴子都依次如此分食桃子。那么桃子数最少应该有几个呢？
// 怎样编程呢？先要找一下第N只猴子和其面前桃子数的关系。如果从第1只开始往第5只找，不好找，但如果思路一变，从第N到第1去，可得出下面的推导式：

// 第N只猴   第N只猴前桃子数目

// 5          s5=x

// 4          s4=s5*5/4+1

// 3          s3=s4*5/4+1

// 2          s2=s3*5/4+1

// 1          s1=s2*5/4+1

// s1即为所求。上面的规律中只要将s1-s5的下标去掉：

// s=x

// s=s*5/4+1

// s=s*5/4+1

// s=s*5/4+1

// s=s*5/4+1

// 所以可以用循环语句加以解决。
package algorithm

import (
	"fmt"
)

const (
	MONKEY_COUNT = 5
	PER_EAT      = 1 // 每次吃一个
)

func CalcBeach2() int {
	n := 1 // 计算最后一只猴子，吃之前有多少个桃子的便宜量
	var count int = 0

Loop:
	for {
		count = 5*n + 1
		n++                                       // 从6开始
		for i := (MONKEY_COUNT - 1); i > 0; i-- { // -1是因为，第五只时，已知最小6个

			if count%4 == 0 { // 能整除
				count = count / 4
				count = count * 5 // count * 5 / 4 + 1
				count = count + 1
				// fmt.Printf("Now: %d\n", count)
			} else {
				// 不能整出，跳出内for
				continue Loop // 下一个大循环
			}

		}
		return count // 完成使命
	}
	return count // 一般不会走到这里
}

// 错误方法，minCount + offset这一步，不满足最后一只猴子
func CalcBeach() int {

	var count int = 0                         // 桃子个数
	var minCount int = MONKEY_COUNT + PER_EAT // 最小桃子个数 = 5个 + 最后被吃的1个 = 6
	var offset = 0                            // 最后是6个假设
Loop:
	for { // 最开始，吃掉一个后，还能分5堆
		count = minCount + offset // 6 开始   //
		offset++                  // 用于下一次遍历
		fmt.Printf("Start Count : %d\n", count)
		for i := (MONKEY_COUNT - 1); i > 0; i-- { // -1是因为，第五只时，已知最小6个

			if count%4 == 0 { // 能整除
				count = count / 4
				count = count * 5 // count * 5 / 4 + 1
				count = count + 1
				// fmt.Printf("Now: %d\n", count)
			} else {
				// 不能整出，跳出内for
				continue Loop // 下一个大循环
			}

		}
		return count // 完成使命

		// if (count)
	}

	return count

}
