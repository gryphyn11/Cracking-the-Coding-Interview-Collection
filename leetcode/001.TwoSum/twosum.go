package leetcode

func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	for index, val := range nums {
		if idx, ok := memo[target-val]; ok {
			return []int{idx, index}
		}
		memo[val] = index
	}
	return nil
}
