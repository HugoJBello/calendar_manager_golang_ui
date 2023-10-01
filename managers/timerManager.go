package managers

import (
	"fmt"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
)

type TimerManager struct {
}

func (m *TimerManager) SetTimer(app *tview.Application, globalAppState *models.GlobalAppState) {

	for now := range time.Tick(60 * time.Second) {
		if globalAppState.RefreshBlocked == false {
			app.Stop()
			fmt.Println(now)
			*&globalAppState.CurrentTime = now
			globalAppState.UpdateDisplayTime()
			*globalAppState.RefreshApp <- "refresh"
		}
	}

}
