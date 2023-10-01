package managers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type WeekViewManager struct {
	ApiManager ApiManager
}

func (m *WeekViewManager) LoadWeekView(globalAppState *models.GlobalAppState) tview.Table {
	timeNow := time.Now()
	_, week := timeNow.ISOWeek()

	dates, _ := m.ApiManager.GetDatesWeek(week)
	fmt.Println(dates)

	table := tview.NewTable().
		SetBorders(true)

	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	//hours := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	hours := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 1, 2, 3, 4, 5, 6, 7}

	datesByWeek := m.organizeByWeekdays(*dates)

	fmt.Println(datesByWeek)

	cols := len(weekDays)
	rows := len(hours)

	for c := 0; c < cols; c++ {
		weekday := weekDays[c]
		datesByHour := make(map[int][]models.Date)
		_, ok := datesByWeek[weekday]
		if ok {
			datesByHour = m.organizeHours(datesByWeek[weekday], hours)
		} else {
			datesByHour = m.fillEmptyHours(hours)
		}

		for r := 0; r < rows; r++ {
			hour := hours[r]

			color := tcell.ColorWhite

			if c < 1 || r < 1 {
				color = tcell.ColorYellow
			}

			if r == 0 {
				table.SetCell(r, c+1,
					tview.NewTableCell(weekday).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else if r > 0 && c == 0 {
				table.SetCell(r, c,
					tview.NewTableCell(strconv.Itoa(hour)+"h").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else if r > 0 && c > 0 {
				datesInHour := datesByHour[hour]
				var datesText = ""
				for _, date := range datesInHour {
					datesText = datesText + date.DateTitle + " \\ "
					table.SetCell(r, c+1,
						tview.NewTableCell(datesText).
							SetTextColor(color).
							SetAlign(tview.AlignCenter))
				}
			}

		}
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			//app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	return *table

}

func (m *WeekViewManager) organizeByWeekdays(dates []models.Date) map[string][]models.Date {
	result := make(map[string][]models.Date)
	for _, date := range dates {
		value, ok := result[date.DayOfWeek]
		if !ok {
			result[date.DayOfWeek] = []models.Date{date}
		} else {
			result[date.DayOfWeek] = append(value, date)
		}
	}
	return result
}

func (m *WeekViewManager) organizeHours(dates []models.Date, hours []int) map[int][]models.Date {
	result := make(map[int][]models.Date)

	for index, _ := range hours {
		hour := hours[index]

		currentDate := time.Date(dates[0].Starts.Year(), dates[0].Starts.Month(), dates[0].Starts.Day(), hour, 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())
		//currentNext := time.Date(dates[0].Starts.Year(), dates[0].Starts.Month(), dates[0].Starts.Day(), hour+1, 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())

		for index, _ := range dates {
			date := dates[index]
			starts := time.Date(date.Starts.Year(), date.Starts.Month(), date.Starts.Day(), date.Starts.Hour(), 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())
			ends := time.Date(date.Ends.Year(), date.Ends.Month(), date.Ends.Day(), date.Ends.Hour(), 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())

			if m.TimeIsBetween(currentDate, starts, ends) {

				value, ok := result[date.Ends.Hour()]
				if !ok {
					result[currentDate.Hour()] = []models.Date{date}
				} else {
					result[currentDate.Hour()] = append(value, date)
				}
			}
		}

	}
	return result
}

func (m *WeekViewManager) fillEmptyHours(hours []int) map[int][]models.Date {
	result := make(map[int][]models.Date)

	for index, _ := range hours {
		hour := hours[index]

		result[hour] = []models.Date{}

	}
	return result
}

func (m *WeekViewManager) TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}
