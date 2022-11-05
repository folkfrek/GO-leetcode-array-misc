func replaceElements(arr []int) []int {
    var n int = len(arr)
    var maxLeft int = arr[len(arr) - 1]
    var res []int =  make([]int, n)
    for i := n-1; i >= 0; i-- {
        res[i] = maxLeft
        if arr[i] > maxLeft{
            maxLeft = arr[i]
        }
    }
    res[n - 1] = -1
    return res
}
// Time Complexity: O(n)
// Space Complexity: O(n) because of res Array
