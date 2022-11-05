func containsDuplicate(nums []int) bool {   
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j]{return true}
        }
    }
    return false
}
// Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)


func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		} else {
			seen[n] = true
		}
	}
	return false
}
// Optimized for speed using a HashMap
// Time: O(n)
// Space: O(n)
