/*
* @Author: Administrator
* @Date:   2017-07-27 08:47:31
* @Last Modified by:   Administrator
* @Last Modified time: 2017-07-27 09:13:17
*/

package main

import (
	"fmt"
	/*
	https://golang.org/pkg/math/rand/
	 */
	"math/rand"
	/*
	https://golang.org/pkg/sync/
	 */
	"sync"
	/*
	
	 https://golang.org/pkg/time/
	 
	*/
	"time"
)

// WaitGroup 用于等待结束  goroutines
var wg sync.WaitGroup

/*
加法表
 */
func addTable() {

	//  结束时把 goroutine 的计数器减一
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("加法表:", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("\t%d+%d=%d\t", i, j, i+j)
		}
		fmt.Println("\n")
	}
}

/*
乘法表
 */
func multiTable() {

	//  结束时把 goroutine 的计数器减一
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("乘法表:", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("\t%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println("\n")
	}
}

func main() {

	// WaitGroup 设置为2, 表示有两个goroutine。
	wg.Add(2)
	fmt.Println("goroutines 开始")
	
	go addTable()
	go multiTable()


	fmt.Println("等待结束")
	wg.Wait()
	fmt.Println("\n结束了")


}