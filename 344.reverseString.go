func reverseString(s []byte)  {
    left, right := 0, len(s) - 1
    for left < right {
        charLeft, charRight := s[left], s[right]
        s[right], s[left] = charLeft, charRight
        left++
        right--
    }
}
// Time Compelexity: O(n)
// Space Complexity: O(1)
