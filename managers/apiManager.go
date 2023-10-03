package managers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/HugoJBello/calendar_manager_golang_ui/models"
)

const version string = "/v1"

const GetDateRoute = version + "/date/last"
const CreateDateRoute = version + "/date/new"
const UpdateDateRoute = version + "/date/save"
const DeleteDateRoute = version + "/date/delete"

type ApiManager struct {
	Url string
}

func (m *ApiManager) GetDates() (*[]models.Date, error) {

	currentUrl := m.Url + GetDateRoute + "?limit=100&skip=0"

	resp, err := http.Get(currentUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}

func (m *ApiManager) GetDatesWeek(week int) (*[]models.Date, error) {

	currentUrl := m.Url + GetDateRoute + "?week=" + strconv.Itoa(week)
	resp, err := http.Get(currentUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}

func (m *ApiManager) GetDatesMonth(month int) (*[]models.Date, error) {

	currentUrl := m.Url + GetDateRoute + "?month=" + strconv.Itoa(month)

	resp, err := http.Get(currentUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}

func (m *ApiManager) GetDatesDay(day int) (*[]models.Date, error) {

	currentUrl := m.Url + GetDateRoute + "?day=" + strconv.Itoa(day)

	resp, err := http.Get(currentUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}

func (m *ApiManager) CreateDateStructFromDate(date models.Date) models.CreateDate {
	return models.CreateDate{DateTitle: date.DateTitle, DateId: date.DateId, DateBody: date.DateBody, Tags: date.Tags, Type: date.Type,
		Starts: date.Starts, Ends: date.Ends, AllDay: date.AllDay, CreatedBy: date.CreatedBy}
}

func (m *ApiManager) CreateDate(date models.CreateDate) (*[]models.Date, error) {

	currentUrl := m.Url + CreateDateRoute
	jsonBody, _ := json.Marshal(date)

	resp, err := http.Post(currentUrl, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}
func (m *ApiManager) UpdateDate(date models.CreateDate) (*[]models.Date, error) {

	currentUrl := m.Url + UpdateDateRoute
	jsonBody, _ := json.Marshal(date)

	resp, err := http.Post(currentUrl, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}

func (m *ApiManager) DeleteDate(id string) (*[]models.Date, error) {
	fmt.Println("deleting")
	currentUrl := m.Url + DeleteDateRoute + "?id=" + id

	req, err := http.NewRequest(http.MethodDelete, currentUrl, nil)
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("error in api")
	}

	bodyGetResp, err := ioutil.ReadAll(resp.Body)
	var dateRes models.DateResponse

	json.Unmarshal(bodyGetResp, &dateRes)

	return &dateRes.Data, nil
}
