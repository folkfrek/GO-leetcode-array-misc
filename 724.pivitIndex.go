func pivotIndex(nums []int) int {
    var rightTotal int = 0
    var leftTotal int = 0
    for _, num := range nums {
        rightTotal += num
    }
    for i, num := range nums {
        leftTotal += num
        if leftTotal == rightTotal {return i}
        rightTotal -= num
    }
    return -1
}
// Time Complexity = O(n)
// Space Complexity = O(1)
