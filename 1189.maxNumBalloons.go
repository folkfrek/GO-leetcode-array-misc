func maxNumberOfBalloons(text string) int {
    ballon := map[rune]int{'b':0,'a':0,'l':0,'o':0, 'n':0}
    for _,char := range text{
        if _, ok := ballon[char]; ok {
            ballon[char]++
        }
    }
    ballon['l'] = ballon['l']/ 2
    ballon['o'] = ballon['o']/ 2
    smallest := len(text)
    for _,val := range ballon {
        if smallest > val{smallest = val}
    }
    return smallest
}

// Time Compexity: O(n)
// Space Complexity: O(5) - size of map
