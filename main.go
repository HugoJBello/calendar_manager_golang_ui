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

func init() {
	gotenv.Load()
	apiManager = managers.ApiManager{Url: os.Getenv("API_URL")}
	weekViewManager = managers.WeekViewManager{ApiManager: apiManager}
	editDateViewManager = managers.EditDateViewManager{ApiManager: apiManager}
	menusManager = managers.MenusManager{WeekViewManager: weekViewManager, EditDateViewManager: editDateViewManager}
	timerManager = managers.TimerManager{}
}

func main() {

	refreshApp := make(chan string)
	globalAppState := models.GlobalAppState{RefreshApp: &refreshApp, CurrentTime: time.Now(), RefreshBlocked: false}

	app := tview.NewApplication()

	go timerManager.SetTimer(app, &globalAppState)

	menusManager.LoadMenus(app, &globalAppState)

	for refresh := range *globalAppState.RefreshApp {

		fmt.Println(refresh)

		if globalAppState.RefreshBlocked == false {
			menusManager.LoadMenus(app, &globalAppState)
		}
	}

}
