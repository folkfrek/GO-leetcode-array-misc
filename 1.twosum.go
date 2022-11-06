func twoSum(nums []int, target int) []int {
    difference := make(map[int]int)
    for i,num := range nums {
        if index, ok := difference[target - num]; ok {
            return []int{index, i }
        }
        difference[num] = i
    }
    return []int{0,0}
}
// Time Complexity = O(n)
// Space Complexity = O(n)
