func daysBetweenDates(date1 string, date2 string) int {
	const format = "2006-01-02"
	day1, _ := time.Parse(format, date1)
	day2, _ := time.Parse(format, date2)

	if day1.After(day2) {
		day1, day2 = day2, day1
	}
	diff := day2.Sub(day1)
	result := int(math.Abs(diff.Hours() / 24))
	return result
}