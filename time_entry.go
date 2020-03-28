package redmine

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type timeEntryResult struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}

type TimeEntry struct {
	Id        int     `json:"id"`         // TimeEntry ID
	Project   *IdName `json:"project"`    // Project
	Issue     *Id     `json:"issue"`      // IssueID
	User      *IdName `json:"user"`       // ユーザー
	Activity  *IdName `json:"activity"`   // 作業分類
	Hours     float32 `json:"hours"`      // 作業時間
	Comments  string  `json:"comments"`   // コメント
	SpentOn   string  `json:"spent_on"`   // 作業日
	CreatedOn string  `json:"created_on"` // 作成日
	UpdatedOn string  `json:"updated_on"` // 更新日
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
	TimeEntry *TimeEntryRequest `json:"time_entry"`
}

type TimeEntryFilter struct {
	ProjectID string
	UserID    string
	From      string
	To        string
}

func (c *Client) GetTimeEntries() ([]TimeEntry, error) {
	timeEntries, err := getTimeEntries(c, c.endpoint+"/time_entries.json?key="+c.apikey)
	if err != nil {
		return nil, err
	}

	return timeEntries, nil

}

func (c *Client) GetTimeEntriesByFilter(f *TimeEntryFilter) ([]TimeEntry, error) {
	timeEntries, err := getTimeEntries(c, c.endpoint+"/time_entries.json?key="+c.apikey+getTimeEntryUrlQueryString(f))
	if err != nil {
		return nil, err
	}

	return timeEntries, nil
}

func (c *Client) CreateTimeEntry(t *TimeEntryRequest) ([]TimeEntry, error) {
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
	if res.StatusCode != http.StatusCreated {
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
	return r.TimeEntries, nil
}

func getTimeEntries(c *Client, url string) ([]TimeEntry, error) {
	res, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var t timeEntryResult
	if res.StatusCode != http.StatusOK {
		var er errorResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&t)
	}
	if err != nil {
		return nil, err
	}

	return t.TimeEntries, nil
}

func getTimeEntryUrlQueryString(filter *TimeEntryFilter) string {
	if filter == nil {
		return ""
	}

	var q string
	if filter.ProjectID != "" {
		q = q + fmt.Sprintf("&project_id=%v", filter.ProjectID)
	}

	if filter.UserID != "" {
		q = q + fmt.Sprintf("&user_id=%v", filter.UserID)
	}

	if filter.From != "" {
		q = q + fmt.Sprintf("&from=%V", filter.From)
	}

	if filter.To != "" {
		q = q + fmt.Sprintf("&to=%V", filter.To)
	}

	return q

}
