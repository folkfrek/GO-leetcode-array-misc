func majorityElement(nums []int) int {
    mid := len(nums) / 2
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
        if freq[num] > mid{return num}
    }
    return -1
}
// Time Complexity: O(n)
// Space Complexity: O(n)
