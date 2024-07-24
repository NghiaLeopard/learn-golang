package func

import "fmt"

type transform func(int) int

func main() {
	valueArr := []int{1, 2, 3, 4}

	doubleArr := transformArr(&valueArr, double)
	tripleArr := transformArr(&valueArr, triple)

	fmt.Println(doubleArr,tripleArr)

}

func transformArr(arr *[]int, fn transform) []int {
	newArr := []int{}
	for _, value := range *arr {
		newArr = append(newArr, fn(value))
	}

	return newArr
}

func double(value int) int {
	return value * 2
}

func triple(value int) int {
	return value * 3
}