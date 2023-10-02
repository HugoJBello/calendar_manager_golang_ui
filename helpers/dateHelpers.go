package helpers

import (
	"strconv"
	"strings"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
)

var WeekDayIntMap = map[string]int{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    7,
}

var Weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func FillEmptyHours(hours []string, weekday string) map[string][]models.Date {
	result := make(map[string][]models.Date)

	for index, _ := range hours {
		hour := hours[index]
		result[hour] = []models.Date{CreateDateInThisWeek(hour, weekday)}
	}
	return result
}

func CreateDateInThisWeek(hourStr, weekday string) models.Date {
	hour, _ := strconv.Atoi(strings.Split(hourStr, ":")[0])
	minutes, _ := strconv.Atoi(strings.Split(hourStr, ":")[1])

	dateNow := time.Now()
	currentWeekDay := dateNow.Weekday().String()

	weekdayInt := WeekDayIntMap[weekday]
	currentWeekDayInt := WeekDayIntMap[currentWeekDay]

	timeResult := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day()+(weekdayInt-currentWeekDayInt), hour, minutes, 0, dateNow.Nanosecond(), dateNow.Location())

	return models.Date{DateTitle: "", DateBody: "", Starts: &timeResult, Ends: &timeResult}
}

func TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}
