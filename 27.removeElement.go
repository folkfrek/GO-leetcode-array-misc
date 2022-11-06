func removeElement(nums []int, val int) int {
    index := 0
    for _,num := range nums {
        if num != val {
            nums[index] = num
            index++
        }
    }
    return index
}
// Time Complexity: O(n)
// Space Complexity: O(n)
