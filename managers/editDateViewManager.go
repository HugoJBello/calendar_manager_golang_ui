package managers

import (
	"strconv"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type EditDateViewManager struct {
	ApiManager ApiManager
}

func (m *EditDateViewManager) LoadNewDateView(pages *tview.Pages, globalAppState *models.GlobalAppState) (*tview.Frame, error) {
	globalAppState.RefreshBlocked = true

	var dateNow = time.Now()

	if globalAppState.SelectedDate == nil {
		globalAppState.SelectedDate = &models.Date{DateId: "", DateTitle: "", DateBody: "", Tags: ""}
	}

	var repeats = ""
	var numIter = "0"
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
		}).
		AddInputField("Repeats weekly", repeats, 20, nil, func(text string) {
			repeats = text
		}).
		AddInputField("Number of repetitions", numIter, 20, nil, func(text string) {
			numIter = text
		}).
		AddButton("Save", func() {
			createDate := m.ApiManager.CreateDateStructFromDate(*globalAppState.SelectedDate)

			if globalAppState.SelectedDate.DateId != "" {
				m.ApiManager.UpdateDate(createDate)
			} else {
				m.ApiManager.CreateDate(createDate)
			}
			globalAppState.RefreshBlocked = false
			pages.SwitchToPage("week-view")

		}).
		AddButton("Quit", func() {
			pages.SwitchToPage("week-view")
			globalAppState.RefreshBlocked = false
		})

	frame := tview.NewFrame(form).SetBorders(2, 2, 2, 2, 4, 4)
	return frame, nil
}
