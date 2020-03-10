package redmine

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type TimeEntryActivity struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	Active    bool   `json:"active"`
}

type IssuePrioritiy struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	Active    bool   `json:"active"`
}

type issuePriorities struct {
	IssuePrioritiesr []IssuePrioritiy `json:"issue_priorities"`
}

type timeEntryActivity struct {
	TimeEntryActivities []TimeEntryActivity `json:"time_entry_activities"`
}

func (c *Client) GetTimeEntryActivities() ([]TimeEntryActivity, error) {
	res, err := c.Get(c.endpoint + "/enumerations/time_entry_activities.json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var t timeEntryActivity
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

	return t.TimeEntryActivities, nil
}

func (c *Client) GetIssuePriorities() ([]IssuePrioritiy, error) {
	res, err := c.Get(c.endpoint + "/enumerations/issue_priorities.json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var i issuePriorities
	if res.StatusCode != http.StatusOK {
		var er errorResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&i)
	}
	if err != nil {
		return nil, err
	}

	return i.IssuePrioritiesr, nil
}
