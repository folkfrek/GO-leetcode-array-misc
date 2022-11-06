func isAnagram(s string, t string) bool {
  var countAlphabet [26]int
	for i := 0; i < len(s); i++ {
		countAlphabet[s[i] - 'a']++
	}
	for i := 0; i < len(t); i++ {
		countAlphabet[t[i] - 'a']--
	}
	for i := 0; i < 26; i++ {
		if countAlphabet[i] != 0{return false}
	}
	return true
}
// Time Complexity = O(2n + 26) = O(n) 
// Space Complexity = O(26) = O(1)
