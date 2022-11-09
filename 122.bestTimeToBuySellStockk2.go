func maxProfit(prices []int) int {
    maximumProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i - 1] < prices[i]{
            maximumProfit += prices[i] - prices[i - 1]
        }
    }
    return maximumProfit
}
// Time Complexity = O(n)
//  Space Complexity = O(1)
