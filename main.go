// Demo code for the List primitive.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/HugoJBello/calendar_manager_golang_ui/managers"
	"github.com/HugoJBello/calendar_manager_golang_ui/models"
	"github.com/rivo/tview"
	"github.com/subosito/gotenv"
)

var apiManager managers.ApiManager
var menusManager managers.MenusManager
var weekViewManager managers.WeekViewManager
var timerManager managers.TimerManager
var editDateViewManager managers.EditDateViewManager
var actionsOnDateViewManager managers.ActionsOnDateViewManager

func init() {
	gotenv.Load()
	apiManager = managers.ApiManager{Url: os.Getenv("API_URL")}
	editDateViewManager = managers.EditDateViewManager{ApiManager: apiManager}
	actionsOnDateViewManager = managers.ActionsOnDateViewManager{ApiManager: apiManager, EditDateViewManager: editDateViewManager}
	weekViewManager = managers.WeekViewManager{ApiManager: apiManager, EditDateViewManager: editDateViewManager, ActionsOnDateViewManager: actionsOnDateViewManager}
	menusManager = managers.MenusManager{WeekViewManager: weekViewManager, EditDateViewManager: editDateViewManager}
	timerManager = managers.TimerManager{}
}

func main() {

	refreshApp := make(chan string)
	timeNow := time.Now()
	_, weekNumber := timeNow.ISOWeek()
	globalAppState := models.GlobalAppState{RefreshApp: &refreshApp, SelectedWeek: weekNumber,
		CurrentTime: timeNow, RefreshBlocked: false}

	app := tview.NewApplication()

	go timerManager.SetTimer(app, &globalAppState)

	menusManager.LoadMenus(app, &globalAppState)

	for refresh := range *globalAppState.RefreshApp {

		fmt.Println("-----", refresh)

		menusManager.LoadMenus(app, &globalAppState)
	}

}
