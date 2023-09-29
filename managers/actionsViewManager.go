package managers

import (
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type ActionsViewManager struct {
	ApiManager ApiManager
}

func (m *ActionsViewManager) AddActionsPage(app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) *tview.Frame {

	list := tview.NewList()

	list.AddItem("Create New Task", "adds a new task", 'c', func() {

	})

	list.AddItem("Quit", "Press to exit", 'q', func() {
		pages.SwitchToPage("tasks_board")
	})

	frame := tview.NewFrame(list).SetBorders(2, 2, 2, 2, 4, 4)

	return frame
}
