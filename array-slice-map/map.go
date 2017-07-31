/*
* @Author: suifengtec
* @Date:   2017-07-31 13:25:57
* @Last Modified by:   suifengtec
* @Last Modified time: 2017-07-31 13:42:20
 */

/*
和其它语言中的字典或映射类似, Go 中的 map 也是不关心元素的顺序的,但是键不能重复。
*/
package main

import (
	"fmt"
)

func main() {

	/*
		使用 Go 内置函数 make 来初始化字典 m1 为: 索引是字符串形式,值是整数形式

	*/
	m1 := make(map[string]int)
	/*
		使用字面量来初始化一个字典 m2
	*/
	m2 := map[string]string{

		"一": "壹",
		"二": "贰",
		"三": "叁",
		"四": "肆",
		"五": "伍",
		"六": "陆",
	}

	fmt.Println(m1)
	fmt.Println(m2)

	/*
		判断一个键是否存在于字典中
	*/
	six, exists := m2["六"]
	if exists {
		fmt.Println(six)
	} else {
		fmt.Println("m2不存在键‘六’")
	}
	seven, exists := m2["七"]
	if exists {
		fmt.Println(seven)
	} else {
		fmt.Println("m2不存在键‘七’")
	}

	fmt.Println("迭代字典 m2")

	for k, v := range m2 {
		fmt.Println(k + "=>" + v)
	}
	fmt.Println("从字典 m2 中删除 键为 '四' 的元素后")

	delete(m2, "四")
	fmt.Println(m2)
}
