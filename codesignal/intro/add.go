package intro

import "fmt"

func add(param1 int, param2 int) int {
	return param1 + param2
}

func main() {
	fmt.Println(add(4, 5))
}
