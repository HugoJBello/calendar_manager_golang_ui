package managers

import (
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenusManager struct {
	WeekViewManager      WeekViewManager
	ButtonBarViewManager ButtonBarViewManager
	EditDateViewManager  EditDateViewManager
}

func (m *MenusManager) LoadMenus(app *tview.Application, globalAppState *models.GlobalAppState) {

	pagesMainMenus := tview.NewPages()

	weekTable := m.WeekViewManager.LoadWeekView(app, pagesMainMenus, globalAppState)
	pagesMainMenus.AddPage("week-view", weekTable, true, true)

	buttonBar := m.ButtonBarViewManager.CreateButtonBarWithPoints(globalAppState)
	lowerBarFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	lowerBarFlex.AddItem(pagesMainMenus, 0, 1, true)
	lowerBarFlex.AddItem(buttonBar, 2, 0, false)

	newSDateFrame, _ := m.EditDateViewManager.LoadNewDateView(app, pagesMainMenus, globalAppState)
	pagesMainMenus.AddPage("new-date-view", newSDateFrame, true, false)

	lowerBarFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			newSDateFrame, _ := m.EditDateViewManager.LoadNewDateView(app, pagesMainMenus, globalAppState)
			pagesMainMenus.RemovePage("new-date-view")
			pagesMainMenus.AddPage("new-date-view", newSDateFrame, true, true)

		} else if event.Key() == tcell.KeyCtrlH {

			globalAppState.SelectedWeek = globalAppState.SelectedWeek - 1
			weekTable := m.WeekViewManager.LoadWeekView(app, pagesMainMenus, globalAppState)
			pagesMainMenus.RemovePage("week-view")
			pagesMainMenus.AddPage("week-view", weekTable, true, true)

		} else if event.Key() == tcell.KeyCtrlL {
			globalAppState.SelectedWeek = globalAppState.SelectedWeek + 1
			weekTable := m.WeekViewManager.LoadWeekView(app, pagesMainMenus, globalAppState)
			pagesMainMenus.RemovePage("week-view")
			pagesMainMenus.AddPage("week-view", weekTable, true, true)
		}
		return event

	})

	if err := app.SetRoot(lowerBarFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
