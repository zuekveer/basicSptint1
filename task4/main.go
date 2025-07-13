/*
Что произойдёт при изменении среза, созданного из массива?
Объясни вывод следующей программы:

	func main() {
		arr := [5]int{1, 2, 3, 4, 5}
		sl := arr[:]
		sl[0] = 3
		fmt.Println(arr)
		fmt.Println(sl)
	}
*/
package task4

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:]
	sl[0] = 3
	fmt.Println(arr)
	fmt.Println(sl)
}

// [3 2 3 4 5]
// [3 2 3 4 5]
