package managers

import (
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type ActionsOnDateViewManager struct {
	ApiManager          ApiManager
	EditDateViewManager EditDateViewManager
}

func (m *ActionsOnDateViewManager) AddActionsPage(app *tview.Application, pages *tview.Pages, globalAppState *models.GlobalAppState) *tview.Frame {
	globalAppState.RefreshBlocked = true

	list := tview.NewList()

	list.AddItem("Create new date at time ", "adds a new date", 'a', func() {
		if globalAppState.SelectedDate.DateTitle != "" {
			globalAppState.SelectedDate.DateTitle = ""
			globalAppState.SelectedDate.DateId = ""
			globalAppState.SelectedDate.AllDay = "false"
			globalAppState.SelectedDate.DateBody = ""
			globalAppState.SelectedDate.Tags = ""
			globalAppState.SelectedDate.Type = nil
		}
		newSDateFrame, _ := m.EditDateViewManager.LoadNewDateView(app, pages, globalAppState)

		pages.RemovePage("new-date-view")
		pages.AddPage("new-date-view", newSDateFrame, true, true)
		pages.SwitchToPage("new-date-view")
	})

	if globalAppState.MultipleSelectedDate == nil && globalAppState.SelectedDate != nil && globalAppState.SelectedDate.DateTitle != "" {
		list.AddItem("Edit date "+globalAppState.SelectedDate.DateTitle, "edits selected date", 'b', func() {
			newSDateFrame, _ := m.EditDateViewManager.LoadNewDateView(app, pages, globalAppState)
			pages.RemovePage("new-date-view")
			pages.AddPage("new-date-view", newSDateFrame, true, true)
			pages.SwitchToPage("new-date-view")

		})

		list.AddItem("Delete date "+globalAppState.SelectedDate.DateTitle, "removes selected date", 'b', func() {
			m.ApiManager.DeleteDate(globalAppState.SelectedDate.DateId)
			go func() {
				app.Stop()
				*globalAppState.RefreshApp <- "refresh"
			}()
		})

	}

	if globalAppState.MultipleSelectedDate != nil {
		for index, _ := range *globalAppState.MultipleSelectedDate {
			list.AddItem("Edit date "+(*globalAppState.MultipleSelectedDate)[index].DateTitle, "edits selected date", 'a', func() {
				selected := list.GetCurrentItem()

				globalAppState.SelectedDate = &(*globalAppState.MultipleSelectedDate)[selected-1]
				newSDateFrame, _ := m.EditDateViewManager.LoadNewDateView(app, pages, globalAppState)
				pages.RemovePage("new-date-view")
				pages.AddPage("new-date-view", newSDateFrame, true, true)
				pages.SwitchToPage("new-date-view")
			})
		}

		for index, _ := range *globalAppState.MultipleSelectedDate {
			list.AddItem("Delete date "+(*globalAppState.MultipleSelectedDate)[index].DateTitle, "edits selected date", 'a', func() {
				selected := list.GetCurrentItem()

				globalAppState.SelectedDate = &(*globalAppState.MultipleSelectedDate)[selected-len(*globalAppState.MultipleSelectedDate)-1]
				m.ApiManager.DeleteDate(globalAppState.SelectedDate.DateId)
				pages.RemovePage("actions-on-date")
				pages.RemovePage("new-date-view")
				pages.SwitchToPage("week-view")
				go func() {
					app.Stop()
					*globalAppState.RefreshApp <- "refresh"
				}()

			})
		}

	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		pages.SwitchToPage("week-view")
		globalAppState.RefreshBlocked = false
	})

	frame := tview.NewFrame(list).SetBorders(2, 2, 2, 2, 4, 4)

	return frame
}
