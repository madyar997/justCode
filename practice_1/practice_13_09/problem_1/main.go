package main

//Написать функцию для сравнения слайса с цело-числовыми значениями
//a. сравнивает два слайса
//b. возвращает булево значение: совпало или нет
//c1. порядок важен
//c2. порядок не важен
func main() {
	//a := []int{1, 2, 3, 3, 3, 5, 6, 6, 9}
	//b := []int{1, 2, 3}
	//equal(a, b)
	//equalWithOrder(a, b)
}

// 1, 2, 3
// 1, 2, 3
//порядок важен
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, num := range a {
		if num != b[i] {
			return false
		}
	}

	return false
}

//func equalWithOrder(a, b []int) bool {
//	aMap := make(map[int]int, len(a))
//	bMap := make(map[int]int, len(b))
//
//	for _, aVal := range a {
//		aMap[aVal]++
//	}
//
//	for _, bVal := range b {
//		bMap[bVal]++
//	}
//
//	for aKey, aVal := range aMap {
//		if bMap[aKey] != aVal {
//			return false
//		}
//	}
//
//	return true
//}
