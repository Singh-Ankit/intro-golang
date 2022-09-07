package main

import (
	"fmt"
)

/* Use MOD operator | (current_index+no.of rotations)%(N)
* When MOD of any input is caluclated on 'N', output is alway between 0 to N
* so, if here N is taken as the size of the array then any mod with N on input,
* results a value which is within the length of the array i.e 0 to N
* So, an element at any index 'i' if we are suppose to rotate that element by 'k' steps,
* after rotating it's new position should be i+k (ex: i = 2 and k = 3 then new position is 2+3 = 5)
* But, what if the to rotate elements at the end of array by 2 steps, we can't go beyond the size of the array
* we should be continuing rotating from end to start index.
* So, for this will use MOD, which will never allow element to go beyong the length 'N',
* Therefore, our logic will be (i+k)%N
* Here will create a new array to store the rotated results at the calculated index,
* ex: temp[(i+k)%N] = arr[i]
 */
func Rotate(input []int, k int) []int {
	temp_arr := []int{}
	//var temp_arr[len(input)]int
	fmt.Println(len(input))
	for i := 0; i < len(input); i++ {
		fmt.Println(i)
		temp_arr[(i+k)%len(input)] = input[i]
	}
	return temp_arr
}

func main() {

	fmt.Println("Array Rotation")
	slice_array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	result := Rotate(slice_array, k)

	for i, j := range result {
		fmt.Println(i, j)
	}
}
