package main

import (
	"fmt"
	"math"
)

func main() {
	// 예제 1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1 : ", n1)
	fmt.Println("ex1 : ", n2)
	fmt.Println("ex1 : ", n3)
	fmt.Println("ex1 : ", n4)
	fmt.Println("ex1 : ", math.MaxInt8)
	fmt.Println("ex1 : ", math.MaxInt16)
	fmt.Println("ex1 : ", math.MaxInt32)
	fmt.Println("ex1 : ", math.MaxInt64)
	fmt.Println("ex1 : ", math.MaxFloat32)
	fmt.Println("ex1 : ", math.MaxFloat64)

	n5 := 100000 // int
	n6 := int16(10000)
	n7 := uint8(100)

	// fmt.Println("ex2 : ", n5+n6) // 에러, 다른 타입끼리는 할 수 없다
	fmt.Println("ex2: ", n5+int(n6))
	// fmt.Println("ex2: ", n6 + n7)
	fmt.Println("ex2: ", n6+int16(n7))
	fmt.Println("ex2: ", n6 > int16(n7))
	fmt.Println("ex2: ", n6-int16(n7) > 5000)

}
