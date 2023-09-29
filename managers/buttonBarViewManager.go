package managers

import (
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ButtonBarViewManager struct {
}

func (m *ButtonBarViewManager) CreateButtonBarWithPoints(globalAppState *models.GlobalAppState) *tview.Frame {

	globalAppState.UpdateDisplayTime()
	currentDate := globalAppState.DisplayCurrentTime
	lowerBarMenu := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 4, 4).
		AddText(currentDate, true, tview.AlignLeft, tcell.ColorWhite).
		AddText("_", true, tview.AlignCenter, tcell.ColorWhite).
		AddText("_", true, tview.AlignRight, tcell.ColorWhite)

	return lowerBarMenu
}
