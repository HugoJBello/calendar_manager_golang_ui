package managers

import (
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type MenusManager struct {
	WeekViewManager      WeekViewManager
	ButtonBarViewManager ButtonBarViewManager
}

func (m *MenusManager) LoadMenus(app *tview.Application, globalAppState *models.GlobalAppState) {

	pagesMainMenus := tview.NewPages()

	weekTable := m.WeekViewManager.LoadWeekView(globalAppState)

	pagesMainMenus.AddPage("week-view", &weekTable, true, true)

	buttonBar := m.ButtonBarViewManager.CreateButtonBarWithPoints(globalAppState)
	lowerBarFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	lowerBarFlex.AddItem(pagesMainMenus, 0, 1, true)
	lowerBarFlex.AddItem(buttonBar, 2, 0, false)

	if err := app.SetRoot(lowerBarFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
