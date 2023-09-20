package main

import "fmt"

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		defer close(nums)
		//считываем данные с канала nums^ возводим в квдрат и перредаем в канал squares
		for num := range nums {
			squares <- num * num
		}
	}()

	go func() {
		//кладем значения в канал nums
		for _, num := range []int{1, 2, 3} {
			nums <- num
		}
		close(nums)
	}()

	for num := range squares {
		fmt.Println(num)
	}
}
