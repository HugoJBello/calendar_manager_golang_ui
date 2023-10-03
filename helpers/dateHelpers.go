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

func RepeatDate(date models.Date, weekdayInt int, times int) []models.Date {
	result := []models.Date{date}
	for i := 1; i < times+1; i++ {
		newDate := result[len(result)-1]
		currentWeekDay := newDate.Starts.Weekday().String()
		currentWeekDayInt := WeekDayIntMap[currentWeekDay]
		daysAdd := 0
		if (weekdayInt - currentWeekDayInt) > 0 {
			daysAdd = weekdayInt - currentWeekDayInt
		} else {
			daysAdd = 7 + weekdayInt - currentWeekDayInt
		}
		newStarts := newDate.Starts.AddDate(0, 0, daysAdd)
		newEnds := newDate.Ends.AddDate(0, 0, daysAdd)
		newDate.Starts = &newStarts
		newDate.Ends = &newEnds
		result = append(result, newDate)
	}
	return result
}

func RepeatDateUntil(date models.Date, weekdayInt int, limit time.Time) []models.Date {
	result := []models.Date{date}
	var currentDate = *date.Starts
	for currentDate.Before(limit) || currentDate.Equal(limit) {

		newDate := result[len(result)-1]
		currentWeekDay := newDate.Starts.Weekday().String()
		currentWeekDayInt := WeekDayIntMap[currentWeekDay]
		daysAdd := 0
		if (weekdayInt - currentWeekDayInt) > 0 {
			daysAdd = weekdayInt - currentWeekDayInt
		} else {
			daysAdd = 7 + weekdayInt - currentWeekDayInt
		}
		newStarts := newDate.Starts.AddDate(0, 0, daysAdd)
		newEnds := newDate.Ends.AddDate(0, 0, daysAdd)
		newDate.Starts = &newStarts
		newDate.Ends = &newEnds
		currentDate = currentDate.AddDate(0, 0, daysAdd)
		if newDate.Starts.Before(limit) {
			result = append(result, newDate)
		}
	}

	return result
}

func FillEmptyHours(hours []string, weekday string, selectedWeek int) map[string][]models.Date {
	result := make(map[string][]models.Date)

	for index, _ := range hours {
		hour := hours[index]
		result[hour] = []models.Date{CreateDateInThisWeek(hour, weekday, selectedWeek)}
	}
	return result
}

func CreateDateInThisWeek(hourStr, weekday string, selectedWeek int) models.Date {
	hour, _ := strconv.Atoi(strings.Split(hourStr, ":")[0])
	minutes, _ := strconv.Atoi(strings.Split(hourStr, ":")[1])

	dateNow := time.Now()
	currentWeekDay := dateNow.Weekday().String()
	_, currentWeekNum := dateNow.ISOWeek()

	weekdayInt := WeekDayIntMap[weekday]
	currentWeekDayInt := WeekDayIntMap[currentWeekDay]

	timeResult := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day()+(weekdayInt-currentWeekDayInt), hour, minutes, 0, dateNow.Nanosecond(), dateNow.Location())

	timeResult = timeResult.AddDate(0, 0, 7*(selectedWeek-currentWeekNum))
	return models.Date{DateTitle: "", DateBody: "", Starts: &timeResult, Ends: &timeResult}
}

func TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) && t.Equal(max)) || (t.Equal(min) || t.After(min)) && (t.Before(max))
}
