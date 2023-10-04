package managers

import (
	"strconv"
	"strings"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type EditDateViewManager struct {
	ApiManager ApiManager
}

func (m *EditDateViewManager) LoadNewDateView(app *tview.Application, pages *tview.Pages, globalAppState *models.GlobalAppState) (*tview.Frame, error) {
	globalAppState.RefreshBlocked = true

	var dateNow = time.Now()

	if globalAppState.SelectedDate == nil {
		globalAppState.SelectedDate = &models.Date{DateId: "", DateTitle: "", DateBody: "", Tags: "", Starts: &dateNow}
	}

	var repeats = ""
	var numIter = "0"
	var repeatUntilDate = globalAppState.SelectedDate.Starts.Format("2006-01-02")

	if globalAppState.SelectedDate.Starts == nil {
		globalAppState.SelectedDate.Starts = &dateNow
	}
	if globalAppState.SelectedDate.Ends == nil {
		globalAppState.SelectedDate.Ends = &dateNow
	}

	form := tview.NewForm().
		AddInputField("title", globalAppState.SelectedDate.DateTitle, 20, nil, func(text string) {
			globalAppState.SelectedDate.DateTitle = text
		}).
		AddInputField("Starts", globalAppState.SelectedDate.Starts.Format("2006-01-02 15:04:05"), 20, nil, func(text string) {
			formatted, _ := time.Parse("2006-01-02 15:04:05", text)
			globalAppState.SelectedDate.Starts = &formatted
		}).
		AddInputField("Ends", globalAppState.SelectedDate.Ends.Format("2006-01-02 15:04:05"), 20, nil, func(text string) {
			formatted, _ := time.Parse("2006-01-02 15:04:05", text)
			globalAppState.SelectedDate.Ends = &formatted
		}).
		AddCheckbox("All day", globalAppState.SelectedDate.AllDay == "true", func(b bool) {
			globalAppState.SelectedDate.AllDay = strconv.FormatBool(b)
		}).
		AddTextArea("Body", globalAppState.SelectedDate.DateBody, 40, 0, 0, func(text string) {
			globalAppState.SelectedDate.DateBody = text
		})

	if globalAppState.SelectedDate.DateId == "" {
		form.AddInputField("Repeats weekly (example: '1, 2')", repeats, 20, nil, func(text string) {
			repeats = text
		}).
			AddInputField("Number of repetitions", numIter, 20, nil, func(text string) {
				numIter = text
			}).
			AddInputField("Repeat until date", repeatUntilDate, 20, nil, func(text string) {
				repeatUntilDate = text
			})
	}

	form.AddButton("Save", func() {

		createDate := m.ApiManager.CreateDateStructFromDate(*globalAppState.SelectedDate)
		createDate.NumberOfIterations, _ = strconv.Atoi(numIter)
		createDate.RepetitionsWeekdays = repeats
		rep, _ := time.Parse("2006-01-02", repeatUntilDate)
		createDate.RepeatUntilDate = &rep

		if globalAppState.SelectedDate.DateId != "" {
			m.ApiManager.UpdateDate(createDate)
		} else {
			m.ApiManager.CreateDate(createDate)
		}
		globalAppState.RefreshBlocked = false
		globalAppState.MultipleSelectedDate = nil
		globalAppState.SelectedDate = nil

		go func() {
			app.Stop()
			*globalAppState.RefreshApp <- "refresh"
		}()
		pages.RemovePage("actions-on-date")
		pages.SwitchToPage("week-view")

	}).
		AddButton("Quit", func() {
			pages.RemovePage("actions-on-date")
			pages.SwitchToPage("week-view")
			globalAppState.RefreshBlocked = false
			globalAppState.MultipleSelectedDate = nil
			globalAppState.SelectedDate = nil
		})

	frame := tview.NewFrame(form).SetBorders(2, 2, 2, 2, 4, 4)
	return frame, nil
}

func formatRepetitions(repeatsTxt string) []int {
	result := []int{}
	if repeatsTxt == "" {
		return result
	}
	if strings.Contains(repeatsTxt, ",") {
		days := strings.Split(repeatsTxt, ",")
		for _, d := range days {
			dInt, _ := strconv.Atoi(d)
			result = append(result, dInt)
		}
	} else {
		dInt, _ := strconv.Atoi(repeatsTxt)
		result = append(result, dInt)
	}
	return result

}
