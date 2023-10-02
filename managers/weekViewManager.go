package managers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/helpers"
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type WeekViewManager struct {
	ApiManager               ApiManager
	EditDateViewManager      EditDateViewManager
	ActionsOnDateViewManager ActionsOnDateViewManager
}

func (m *WeekViewManager) LoadWeekView(pages *tview.Pages, globalAppState *models.GlobalAppState) tview.Table {
	timeNow := time.Now()
	_, week := timeNow.ISOWeek()

	dates, _ := m.ApiManager.GetDatesWeek(week)
	fmt.Println(dates)

	table := tview.NewTable().
		SetBorders(true).SetSelectable(true, true)

	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	//hours := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	//hours := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 1, 2, 3, 4, 5, 6, 7}
	hours := []string{"8:00", "9:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00", "1:00", "2:00", "3:00", "4:00", "5:00", "6:00", "7:00"}

	var datesByWeek map[string][]models.Date = make(map[string][]models.Date)
	if dates != nil {
		datesByWeek = m.organizeByWeekdays(*dates)
	}

	fmt.Println(datesByWeek)

	cols := len(weekDays)
	rows := len(hours)
	datesRowsCols := make(map[string][]models.Date)

	for c := 0; c < cols; c++ {
		weekday := weekDays[c]
		datesByHour := make(map[string][]models.Date)
		_, ok := datesByWeek[weekday]
		if ok {
			datesByHour = m.organizeHours(datesByWeek[weekday], hours)
		} else {
			datesByHour = helpers.FillEmptyHours(hours, weekday)
		}

		for r := 0; r < rows; r++ {
			hour := hours[r]

			if r == 0 {
				table.SetCell(0, c+1,
					tview.NewTableCell(weekday).
						SetTextColor(tcell.ColorYellow).
						SetAlign(tview.AlignCenter))
			}
			if c == 0 {
				table.SetCell(r+1, 0,
					tview.NewTableCell(hour+"h").
						SetTextColor(tcell.ColorYellow).
						SetAlign(tview.AlignCenter))
			}

			datesInHour := datesByHour[hour]
			datesRowsCols[strconv.Itoa(r+1)+"-"+strconv.Itoa(c+1)] = datesInHour

			var datesText = ""
			for index, _ := range datesInHour {
				date := datesInHour[index]
				if datesText == "" {
					datesText = date.DateTitle
				} else {
					datesText = datesText + " \\ " + date.DateTitle
				}
				table.SetCell(r+1, c+1,
					tview.NewTableCell(datesText).
						SetTextColor(tcell.ColorWhite).
						SetAlign(tview.AlignCenter))
			}

			cell := table.GetCell(r+1, c+1)
			if cell == nil || cell.Text == "" {
				table.SetCell(r+1, c+1,
					tview.NewTableCell("").
						SetTextColor(tcell.ColorWhite).
						SetAlign(tview.AlignCenter))
			}

		}
	}

	sr, sc := getSelectedFromCurrentDate(hours)
	table.Select(sr, sc).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			//app.Stop()
		}
		if key == tcell.KeyEnter {
			//table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
	})

	table.SetSelectedFunc(func(row, column int) {
		key := strconv.Itoa(row) + "-" + strconv.Itoa(column)

		selecteds, ok := datesRowsCols[key]
		if ok {
			if len(selecteds) > 1 {
				globalAppState.MultipleSelectedDate = &selecteds
				globalAppState.SelectedDate = &selecteds[0]
			} else if len(selecteds) == 1 {
				globalAppState.SelectedDate = &selecteds[0]
			} else {
				hour := hours[row-1]
				week := helpers.Weekdays[column-1]
				date := helpers.CreateDateInThisWeek(hour, week)
				globalAppState.SelectedDate = &date
			}
			actionsFrame := m.ActionsOnDateViewManager.AddActionsPage(pages, globalAppState)
			pages.RemovePage("actions-on-date")
			pages.AddPage("actions-on-date", actionsFrame, true, true)
			pages.SwitchToPage("actions-on-date")
		}
	})

	return *table

}

func getSelectedFromCurrentDate(hours []string) (r int, c int) {
	r = 0
	dateNow := time.Now()
	hour := dateNow.Hour()
	initialHourStr := hours[0]
	initialHour, _ := strconv.Atoi(strings.Split(initialHourStr, ":")[0])
	if hour > initialHour {
		r = hour - initialHour + 1
	} else {
		r = 24 - hour + 1
	}

	c = helpers.WeekDayIntMap[dateNow.Weekday().String()]
	return r, c
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

func (m *WeekViewManager) organizeHours(dates []models.Date, hours []string) map[string][]models.Date {
	result := make(map[string][]models.Date)

	for index, _ := range hours {
		hour, _ := strconv.Atoi(strings.Split(hours[index], ":")[0])
		minute, _ := strconv.Atoi(strings.Split(hours[index], ":")[1])
		keyStr := hours[index]
		currentDate := time.Date(dates[0].Starts.Year(), dates[0].Starts.Month(), dates[0].Starts.Day(), hour, minute, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())
		//currentNext := time.Date(dates[0].Starts.Year(), dates[0].Starts.Month(), dates[0].Starts.Day(), hour+1, 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())

		for index, _ := range dates {
			date := dates[index]
			starts := time.Date(date.Starts.Year(), date.Starts.Month(), date.Starts.Day(), date.Starts.Hour(), 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())
			ends := time.Date(date.Ends.Year(), date.Ends.Month(), date.Ends.Day(), date.Ends.Hour(), 0, 0, dates[0].Starts.Nanosecond(), dates[0].Starts.Location())

			if helpers.TimeIsBetween(currentDate, starts, ends) {

				value, ok := result[keyStr]
				if !ok {
					result[keyStr] = []models.Date{date}
				} else {
					result[keyStr] = append(value, date)
				}
			}
		}

	}
	return result
}
