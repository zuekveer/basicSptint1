/*
	Что произойдёт при нескольких append() из одного исходного среза?

Что выведет этот код и почему?

		func main() {
		var x []int
		x = append(x, 0)
		x = append(x, 1)
		x = append(x, 2)
		y := append(x, 3)
		z := append(x, 4)
		fmt.Println(y)
		fmt.Println(z)
	}
*/
package task5

import "fmt"

func main() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y)
	fmt.Println(z)
}

// [0 1 2 4]
// [0 1 2 4]
