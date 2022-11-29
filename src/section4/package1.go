// 패키지(1)
package main

// 선언 방법1
/*
import "fmt"
import "os"
*/

// 선언 방법2
import (
	"fmt"
	"os"
)

func main() {
	var name string

	fmt.Println("이름은? : ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
