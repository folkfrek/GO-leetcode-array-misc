func moveZeroes(nums []int)  {
    nonZeroIndex := 0
    for i, num := range nums {
        if num != 0 {
            nums[i] = 0
            nums[nonZeroIndex] = num
            nonZeroIndex++
        }
    }
}
//  Time Complexity = O(n)
// Space Complexity = O(1)
