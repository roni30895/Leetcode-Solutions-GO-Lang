func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    count :=0
    i,j :=0,len(people)-1
    for i <= j {
        if people[i] + people[j] <= limit {
            i++
            j--
        } else{
            j--
        }
        count++
    }
    return count    
}
