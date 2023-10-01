package managers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type EditDateViewManager struct {
	ApiManager ApiManager
}

func (m *EditDateViewManager) LoadNewDateView(app *tview.Application, pages *tview.Pages, globalAppState *models.GlobalAppState) (*tview.Frame, error) {
	globalAppState.RefreshBlocked = true
	var date = globalAppState.SelectedDate

	var dateNow = time.Now()
	var dateNowString = dateNow.Format("2006-01-02 15:04:05")

	if date == nil {
		globalAppState.SelectedDate = &models.Date{DateId: "", DateTitle: "", DateBody: ""}
		date = globalAppState.SelectedDate
	}

	var repeats = ""
	var numIter = "0"
	if date.Starts != nil {
		date.Starts = &dateNow
	}
	if date.Ends != nil {
		date.Ends = &dateNow
	}

	form := tview.NewForm().
		AddInputField("title", date.DateTitle, 20, nil, func(text string) {
			date.DateTitle = text
		}).
		AddInputField("Starts", dateNowString, 20, nil, func(text string) {
			formatted, _ := time.Parse("2006-01-02 15:04:05", text)
			date.Starts = &formatted
		}).
		AddInputField("Ends", dateNowString, 20, nil, func(text string) {
			formatted, _ := time.Parse("2006-01-02 15:04:05", text)
			date.Ends = &formatted
		}).
		AddCheckbox("All day", date.AllDay == "true", func(b bool) {
			date.AllDay = strconv.FormatBool(b)
		}).
		AddTextArea("Body", date.DateBody, 40, 0, 0, func(text string) {
			date.DateBody = text
		}).
		AddInputField("Repeats weekly", repeats, 20, nil, func(text string) {
			repeats = text
		}).
		AddInputField("Number of repetitions", numIter, 20, nil, func(text string) {
			numIter = text
		}).
		AddButton("Save", func() {
			createDate := m.ApiManager.CreateDateStructFromDate(*date)
			if date.DateId != "" {
				m.ApiManager.UpdateDate(createDate)
			} else {
				m.ApiManager.CreateDate(createDate)
			}
			fmt.Println(createDate)
			globalAppState.RefreshBlocked = false
			globalAppState.SelectedDate = date
			pages.SwitchToPage("week-view")

		}).
		AddButton("Quit", func() {
			pages.SwitchToPage("week-view")
			globalAppState.RefreshBlocked = true
		})

	frame := tview.NewFrame(form).SetBorders(2, 2, 2, 2, 4, 4)
	return frame, nil
}
