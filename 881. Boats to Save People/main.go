func numRescueBoats(people []int, limit int) int {
    res := 0
    sort.Ints(people)
    i, j := 0, len(people)-1
    for i <= j {
        if people[i] + people[j] <= limit {
            j--
            i++
        } else {
            j--
        }
        res++
    }
    return res
}