func productExceptSelf(nums []int) []int {
    pre, n := 1, len(nums)
    output := make([]int, n)
    for i:=0; i < n; i++ {
        output[i] = 1 * pre
        pre *= nums[i]
    }
    post := 1
    for i := n - 1; i >= 0; i-- {
        output[i] *= post
        post *= nums[i]
    }
    return output
}
// Time Complexity = O(n)
// Space Complexity = O(n) -- building output array
