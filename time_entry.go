package redmine

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type timeEntryResult struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

type TimeEntry struct {
	Id        int     `json:"id"`
	Project   IdName  `json:"project"`
	Issue     Id      `json:"issue"`
	User      IdName  `json:"user"`
	Activity  IdName  `json:"activity"`
	Hours     float32 `json:"hours"` //作業時間
	Comments  string  `json:"comments"`
	SpentOn   string  `json:"spent_on"` //作業日
	CreatedOn string  `json:"created_on"`
	UpdatedOn string  `json:"updated_on"`
}

type TimeEntryRequest struct {
	ProjectId  int     `json:"project_id,omitempty"`
	IssueId    int     `json:"issue_id,omitempty"`
	SpentOn    string  `json:"spent_on,omitempty"`
	Hours      float32 `json:"hours,omitempty"`
	ActivityId int     `json:"activity_id,omitempty"`
	Comments   string  `json:"comments,omitempty"`
}

type timeEntryRequest struct {
	TimeEntry TimeEntryRequest `json:"time_entry"`
}

func (c *Client) CreateTimeEntry(t TimeEntryRequest) (*TimeEntry, error) {
	var item timeEntryRequest
	item.TimeEntry = t
	s, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endpoint+"/time_entries.json", strings.NewReader(string(s)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Redmine-API-Key", c.apikey)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	var r timeEntryResult
	if res.StatusCode != 201 {
		var er errorResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.TimeEntry, nil
}
