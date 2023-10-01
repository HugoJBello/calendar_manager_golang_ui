package managers

import (
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenusManager struct {
	WeekViewManager      WeekViewManager
	ButtonBarViewManager ButtonBarViewManager
	NewDateViewManager   NewDateViewManager
}

func (m *MenusManager) LoadMenus(app *tview.Application, globalAppState *models.GlobalAppState) {

	pagesMainMenus := tview.NewPages()

	weekTable := m.WeekViewManager.LoadWeekView(globalAppState)

	pagesMainMenus.AddPage("week-view", &weekTable, true, true)

	buttonBar := m.ButtonBarViewManager.CreateButtonBarWithPoints(globalAppState)
	lowerBarFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	lowerBarFlex.AddItem(pagesMainMenus, 0, 1, true)
	lowerBarFlex.AddItem(buttonBar, 2, 0, false)

	lowerBarFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			newSDateFrame, _ := m.NewDateViewManager.LoadNewDateView(app, pagesMainMenus, globalAppState)
			pagesMainMenus.AddPage("new-date-view", newSDateFrame, true, true)

		}
		return event

	})

	if err := app.SetRoot(lowerBarFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
