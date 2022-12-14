package main

import "fmt"

func main(){
	// 예제 1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1,2,3,4,5} // 배열의 개수와 데이터 타입 모두 일치해야 넣을 수 있다.
	var arr3 = [5]int{1,2,3,4,5}
	arr4 := [5]int{1,2,3,4,5}
	arr5 := [5]int{1,2,3}// 나머지는 0으로 초기화
	arr6 := [...]int{1,2,3,4,5} // 배열 크기 자동 맞춤
	arr7 := [5][5]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
	}

	arr1[2] = 5 
	fmt.Printf("%-5T %d %v\n", arr1,len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2,len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3,len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4,len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5,len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6,len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7,len(arr7), arr7)

	// 예제 2
	arr8 := [5]int{1,2,3,4,5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim","lee","park"}
	fmt.Printf("%-5T %d %v\n", arr8,len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9,len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10,len(arr10), arr10)
}