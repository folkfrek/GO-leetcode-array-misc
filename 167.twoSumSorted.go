func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers) - 1
    sumOfTwoNums := numbers[left] + numbers[right]
    for sumOfTwoNums != target {
        if sumOfTwoNums < target{
            left++
        } else {
            right--
        }
        sumOfTwoNums = numbers[left] + numbers[right]
    }
    return []int{left+1, right+1}
}
// Time Complexity: O(n)
// Space Complexity: O(1)
