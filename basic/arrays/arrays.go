package arrays

import "fmt"

func main() {
	var ar1 [3]int
	ar1[0] = 1
	ar1[1] = 2
	ar1[2] = 3
	fmt.Println(ar1)
	sl1 := []int{1, 2, 3}
	sl1 = append(sl1, 1)
}
