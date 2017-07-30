/*
* @Author: suifengtec
* @Date:   2017-07-30 13:30:31
* @Last Modified by:   suifengtec
* @Last Modified time: 2017-07-30 13:37:39
*/
/*
go build -o b.exe multi-table.go && b
*/
package main

import "fmt"

func multiTable_v1() {

	var (
		i int
		j int
	)

	for i = 1; i <= 9; i++ {

		for j = 1; j <= 9; j++ {

			fmt.Printf("%d*%d=%2d\t", i, j, i*j)
		}

		fmt.Printf("\n")
	}

}

func multiTable_v2() {

	var (
		i int
		j int
	)

	for i = 1; i <= 9; i++ {

		for j = 1; j <= 9; j++ {

			if j < i {

				fmt.Printf("        ")

			} else {

				fmt.Printf("%d*%d=%2d\t", i, j, i*j)
			}

		}

		fmt.Printf("\n")
	}

}

func multiTable_v3() {

	var (
		i int
		j int
	)

	for i = 1; i <= 9; i++ {

		for j = 1; j <= 9; j++ {

			if j > i {

				fmt.Printf("        ")
			} else {

				fmt.Printf("%d*%d=%2d\t", i, j, i*j)
			}

		}

		fmt.Printf("\n")
	}

}

func multiTable_v4() {

	var (
		i int
		j int
		n int
	)

	for i = 1; i <= 9; i++ {
		for n = 1; n <= 9-i; n++ {

			fmt.Printf("        ")
		}

		for j = 1; j <= 9; j++ {

			fmt.Printf("%d*%d=%2d\t", i, j, i*j)

		}

		fmt.Printf("\n")
	}

}

func main() {

	fmt.Printf("\nVer1\n")
	fmt.Printf("\n********\n")

	multiTable_v1()

	fmt.Printf("\nVer2\n")
	fmt.Printf("\n********\n")

	multiTable_v2()

	fmt.Printf("\nVer3\n")
	fmt.Printf("\n********\n")
	multiTable_v3()

	fmt.Printf("\nVer 4\n")
	fmt.Printf("\n********\n")

	multiTable_v4()

}