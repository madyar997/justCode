package main

func main() {
	nums := []int{1, 2, 4, 5, 6}
	target := 9
	twoSum(nums, target)

}

// 2,7,11,15
//O(n2)
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//
//	return []int{}
//}

// a + b = target
// b = target - a
// 3 + 5 = 8
// a + (target -a) = target

// 1 2 4 5 6 target = 9
//При наличии ключа значение дополнительной переменной ok будет равно true,  в противном случае — false
// O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		potentialMatch := target - nums[i]
		idx, ok := m[potentialMatch]
		if ok {
			return []int{i, idx}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}
