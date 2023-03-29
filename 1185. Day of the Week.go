func dayOfTheWeek(day int, month int, year int) string {
	date := day
	today := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
	d := today.Weekday().String()
	return d

}