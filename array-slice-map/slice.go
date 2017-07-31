/*
* @Author: suifengtec
* @Date:   2017-07-31 12:29:45
* @Last Modified by:   suifengtec
* @Last Modified time: 2017-07-31 13:25:36
 */

package main

import (
	"fmt"
	"strconv"
)

/*
切片对底层数组进行了抽象，它有三个字段: 底层数组指针,长度,容量,其长度不应小于容量。

在 64位机器上,除了切片所对应的底层数组,切片自身使用24个字节表示。


*/
func main() {
	/*
		字面量方式声明一个字符串类型的切片
	*/
	s0 := []string{"aa", "bb", "cc", "dd", "ee"}
	/*
		使用 Go 内置函数 make 声明一个长度和容量都是3的字符串型切片
	*/
	s1 := make([]string, 3)
	s1[0] = "第001个"
	s1[1] = "第002个"
	s1[2] = "第003个"
	/*
		使用 Go 内置函数 make 声明一个长度为2,容量为3的字符串型切片
	*/
	s2 := make([]string, 2, 3)

	s2[0] = "第1个"
	s2[1] = "第2个"
	/*
		s2 现在是
		[第001个 第002个 第003个]
	*/

	/*
		基于切片 s0,创建一个从 s0 的第二个元素(索引为1)开始的,长度为2,容量为4的切片


		基于容量为k 的切片 s0, 从 s0 的索引为 i 的元素开始, 创建新切片 s3 := s0[i:j] ,
		新切片 s3 的长度为 j-i, 容量为 k-i

		本例中,s3 的
		长度 len(s3) : 3-1=2
		容量为 cap(s3) : 5-1=4
	*/
	s3 := s0[1:3]

	l3 := len(s3)
	c3 := cap(s3)
	/*
		因为切片是引用类型的数据类型,所以切片s3 的索引为1的元素被以字面量赋值的方式修改后,s0的第3个项目也会被修改

	*/
	s3[1] = "hehehe"
	/*

		s3 现在是
		[bb hehehe]
		s1 现在是
		[aa bb hehehe dd ee]


		由于改切片的容量是 4,所以最多可以再追加2个元素

		s3[5] = "ss"
		或者


		的话,运行时将会给出越界错误
		panic: runtime error: index out of range

		但如果给 s3 追加元素

		s3 = append(s3, "haha", "houhou","ss")

		的话,是可以的,但追加的元素不会影响到 s0


	*/
	s3 = append(s3, "haha", "houhou", "ss", "tt")

	/*
		s3 现在是
		[bb hehehe haha houhou ss tt]

		基于 s3 , 从 s3 的索引为2的元素开始,创建一个长度为1,容量为2的切片 s4

		s3[2] 是haha
	*/
	s4 := s3[2:3:4]

	/*
		把 s0 合并到 s4 中

	*/

	s4 = append(s4, s0...)
	/*
		s4 现在是
		[haha aa bb hehehe dd ee]
	*/
	/*
		长度和容量为4,第4个元素为指定值
	*/

	s5 := []string{3: "这是第4个"}

	/*
		创建一个整形切片的切片

		每个元素都是一个切片,第一个元素使用一个整数8来初始化,第二个元素使用两个整数32,64来初始化
	*/
	s6 := [][]int{{8}, {32, 64}}
	/*
		给第一个元素追加上一个整数 16
		切片现在是
		[[8 16] [32 64]]
	*/
	s6[0] = append(s6[0], 16)

	fmt.Println("s0")
	fmt.Printf("%v\n", s0)
	fmt.Println("s1")
	fmt.Printf("%v\n", s1)
	fmt.Println("s2")
	fmt.Printf("%v\n", s2)
	fmt.Println("s3")
	fmt.Printf("%v\n", s3)

	fmt.Println("切片s3的长度为:" + strconv.Itoa(l3))
	fmt.Println("切片s3的容量为:" + strconv.Itoa(c3))

	fmt.Println("s4")
	fmt.Printf("%v\n", s4)

	fmt.Println("切片s4的迭代")
	/*
		//或者
		for _, v := range s4 {
			fmt.Println(v)
		}
	*/
	for i, v := range s4 {

		/*
			//或者
			fmt.Printf("%d\t=>\t%s\n", i, v)
		*/
		fmt.Println(strconv.Itoa(i) + "=>" + v)

	}
	fmt.Println("s5")
	fmt.Printf("%v\n", s5)

	fmt.Println("多维切片s6")
	fmt.Printf("%v\n", s6)

}
