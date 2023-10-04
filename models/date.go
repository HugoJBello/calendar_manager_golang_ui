package models

import "time"

type Date struct {
	Id                 uint       `json:"id" bson:"id" gorm:"primary_key"`
	DateId             string     `json:"dateId" bson:"dateId"`
	RepeatedFromDateID string     `json:"repeatedFromDateID" bson:"repeatedFromDateID"`
	DateTitle          string     `json:"dateTitle" bson:"dateTitle"`
	DateBody           string     `json:"dateBody" bson:"dateBody"`
	Tags               string     `json:"tags" bson:"tags"`
	DayOfWeek          string     `json:"dayOfWeek" bson:"dayOfWeek"`
	Month              int        `json:"month" bson:"month"`
	Week               int        `json:"week" bson:"week"`
	Day                int        `json:"day" bson:"day"`
	Type               *string    `json:"type" bson:"type"`
	CreatedAt          *time.Time `json:"createdAt" bson:"createdAt"`
	EditedAt           *time.Time `json:"editedAt" bson:"editedAt"`
	Starts             *time.Time `json:"starts" bson:"starts"`
	Ends               *time.Time `json:"ends" bson:"ends"`
	AllDay             string     `json:"allDay" bson:"allDay"`
	CreatedBy          string     `json:"createdBy" bson:"createdBy"`
}

type CreateDate struct {
	DateId              string     `json:"dateId" bson:"dateId"`
	DateTitle           string     `json:"dateTitle" bson:"dateTitle"`
	DateBody            string     `json:"dateBody" bson:"dateBody"`
	Tags                string     `json:"tags" bson:"tags"`
	Type                *string    `json:"type" bson:"type"`
	Starts              *time.Time `json:"starts" bson:"starts"`
	Ends                *time.Time `json:"ends" bson:"ends"`
	AllDay              string     `json:"allDay" bson:"allDay"`
	NumberOfIterations  int        `json:"numberOfIterations" bson:"numberOfIterations"`
	RepetitionsWeekdays string     `json:"repetitionsWeekdays" bson:"repetitionsWeekdays"`
	RepeatUntilDate     *time.Time `json:"repeatUntilDate" bson:"repeatUntilDate"`
	CreatedBy           string     `json:"createdBy" bson:"createdBy"`
}

type DateResponse struct {
	message string `json:"message" bson:"message"`
	Data    []Date `json:"data" bson:"data"`
}
