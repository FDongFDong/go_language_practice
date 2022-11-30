package main
import "fmt"

func main(){
	// 예제 1
	var map1 map[string]int = make(map[string]int)// 정석, key는 string, value는 int	
	var map2 = make(map[string]int) // 자료형 생략
	map3 := make(map[string]int) // 리터럴 형

	fmt.Println("ex1 : ",map1)
	fmt.Println("ex1 : ",map2)
	fmt.Println("ex1 : ",map3)
	fmt.Println()

	// 예제 2
	map4 := map[string]int{} // Json 형태
	map4["apple"] =25
	map4["banana"] =40
	map4["orange"] = 33

	map5 := map[string]int {
		"apple": 15,
		"banana":40,
		"orange":23,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex2 : ",map4)
	fmt.Println("ex2 : ",map5)
	fmt.Println("ex2 : ",map6)
	fmt.Println("ex2 : ",map6["orange"])
	fmt.Println("ex2 : ",map6["apple"])
}