func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    var count int
    for _, value := range items {
        if ruleKey=="type"{
            if value[0]==ruleValue {
                count++
            }
        }
        if ruleKey=="color"{
            if value[1]==ruleValue {
                count++
            }
        }
        if ruleKey=="name"{
            if value[2]==ruleValue {
                count++
            }
        }
    }
    return count
}
