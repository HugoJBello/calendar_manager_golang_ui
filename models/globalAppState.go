package models

import (
	"time"
)

type GlobalAppState struct {
	CurrentTime           time.Time
	DisplayCurrentTime    string
	RefreshApp            *chan string
	RefreshBlocked        bool
	SelectedDate          *Date
	MultipleSelectedDate  *[]Date
	MultipleSelectedIndex int
}

func (g *GlobalAppState) UpdateDisplayTime() {
	g.DisplayCurrentTime = g.CurrentTime.Format("Mon Jan 02 15:04:05 -0700 2006")
}
