/*
* @Author: suifengtec
* @Date:   2017-07-31 13:43:39
* @Last Modified by:   suifengtec
* @Last Modified time: 2017-07-31 14:05:15
 */

package main

import (
	"fmt"
)

func main() {

	/*
		初始化一个有6个元素的数组,数组的各个元素值使用0填充
		[0 0 0 0 0 0]
	*/
	a1 := [6]int{}
	/*
		初始化一个有3个元素的数组,并使用指定的值填充
	*/
	a2 := [3]int{6, 7, 8}
	/*
		初始化一个数组,并使用指定的值填充
	*/
	a3 := [...]int{7, 8, 9, 10}
	/*
		初始化一个有10个元素的数组,并指定某些索引的值
		[0 6 0 0 0 0 8 0 9 0]
	*/
	a4 := [10]int{1: 6, 6: 8, 8: 9}
	/*
		初始化一个有9个元素(最大的索引加1)的数组:

		并指定某些索引的值
	*/
	a5 := [...]int{2: 6, 8: 10}

	/*
		指针数组
	*/

	/*a6 := [5]*int{0: new(int), 1: new(int)}*/
	a6 := [6]*int{0: new(int), 1: new(int)}
	*a6[0] = 16
	*a6[1] = 36

	/*
		类型和元素数量相同的数组之间可以相互赋值
	*/

	var a7 [3]string
	a8 := [3]string{"aa", "bb", "cc"}
	a7 = a8

	a9 := [5][2]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}}

	a9[0] = [...]int{111, 222}

	fmt.Println("a1")
	fmt.Println(a1)
	fmt.Println("a2")
	fmt.Println(a2)
	fmt.Println("a3")
	fmt.Println(a3)
	fmt.Println("a4")
	fmt.Println(a4)
	fmt.Println("a5")
	fmt.Println(a5)

	fmt.Println("a6")
	fmt.Println(a6)
	/*
		指针数组的迭代
	*/
	fmt.Println("指针数组 a6 的迭代")
	for i, v := range a6 {
		if v != nil {
			fmt.Printf("%d=>\t%d\n", i, *v)
		}
	}

	fmt.Println("a7")
	fmt.Println(a7)
	fmt.Println("多维数组a9")
	fmt.Println(a9)
}
