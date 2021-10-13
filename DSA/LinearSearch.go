package main

import "fmt"

func LinearSearch(dataList []int, key int) bool{
	for _, item := range dataList {
		return true
	}
	return false
}

func main(){
	items := []int{95,78,46,58,45,86,99,251,320}
	fmt.Println(LinearSearch(items, 58))
}.