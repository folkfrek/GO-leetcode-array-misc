func findDisappearedNumbers(nums []int) []int {
    numbers := make(map[int]bool)
    for i:=1; i <= len(nums); i++{
        numbers[i]=true
    }
    for _,num := range nums {
        if _, ok := numbers[num]; ok {
            delete(numbers, num)
        }
    }
    var output []int
    for key,_ := range numbers {
        output = append(output,key)
    }
    return output
}
// Time Complexity = O(n)
// Space Complexity = O(n)
