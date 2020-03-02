package redmine

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type issueResult struct {
	Issue Issue `json:"issue" yaml:"issue"`
}

type issuesResult struct {
	Issues     []Issue `json:"issues"`
	TotalCount int     `json:"total_count"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
}

type Issue struct {
	Id                  int         `json:"id"`
	Project             *IdName     `json:"project"`
	Tracker             *IdName     `json:"tracker"`
	Status              *IdName     `json:"status"`
	Priority            *IdName     `json:"priority"`
	Author              *IdName     `json:"author"`
	AssignedTo          *IdName     `json:"assigned_to"`
	Parent              *Id         `json:"parent"`
	Subject             string      `json:"subject"`
	Description         string      `json:"description"`
	StartDate           string      `json:"start_date"`
	DueDate             string      `json:"due_date"`
	DoneRatio           int         `json:"done_date"`
	IsPrivate           bool        `json:"is_private"`
	EstimatedHours      float32     `json:"estimated_hours"`
	TotalEstimatedHours float32     `json:"total_estimated_hours"`
	CreatedOn           string      `json:"created_on"`
	UpdatedOn           string      `json:"updated_on"`
	ClosedOn            string      `json:"closed_on"`
	Relations           []*Relation `json:"relations"`
	Journals            []*Journal  `json:"journals"`
}

type Relation struct {
	Id           int    `json:"id"`
	IssueId      int    `json:"issue_id"`
	IssueToId    int    `json:"issue_to_id"`
	RelationType string `json:"relation_type"`
	Delay        string `json:"delay"`
}

type Journal struct {
	Id           int              `json:"id"`
	User         IdName           `json:"user"`
	Notes        string           `json:"notes"`
	CreatedOn    string           `json:"created_on"`
	PrivateNotes bool             `json:"private_notes"`
	Details      []*JournalDetail `json:"details"`
}

type JournalDetail struct {
	Property string `json:"property"`
	Name     string `json:"name"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

type IssueFilter struct {
	IssueId      string
	ProjectId    string
	SubProjectId string
	TrackerId    string
	StatusId     string
	AssignedToId string
	ParentId     string
}

func (c *Client) GetIssueById(issueId int) (*Issue, error) {
	res, err := c.Get(c.endpoint + "/issues/" + strconv.Itoa(issueId) + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var i issueResult
	if res.Status != "200" {
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

	return &i.Issue, nil
}

func (c *Client) GetIssues() ([]Issue, error) {
	res, err := c.Get(c.endpoint + "/issues.json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var i issuesResult
	if res.Status != "200" {
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

	return i.Issues, nil
}

func (c *Client) GetIssueByFilter(f *IssueFilter) ([]Issue, error) {

}

func getIssueUrlQueryString(filter *IssueFilter) string {
	if filter == nil {
		return ""
	}

	var q string
	if filter.IssueId != "" {
		q = q + fmt.Sprintf("&issue_id=%v", filter.IssueId)
	}

	if filter.ProjectId != "" {
		q = q + fmt.Sprintf("&project_id=%v", filter.ProjectId)
	}

	if filter.SubProjectId != "" {
		q = q + fmt.Sprintf("&subproject_id=%v", filter.SubProjectId)
	}

	if filter.TrackerId != "" {
		q = q + fmt.Sprintf("&tracker_id=%v", filter.TrackerId)
	}

	if filter.StatusId != "" {
		q = q + fmt.Sprintf("&tracker_id=%v", filter.StatusId)
	}

	if filter.AssignedToId != "" {
		q = q + fmt.Sprintf("&tracker_id=%v", filter.AssignedToId)
	}

	if filter.ParentId != "" {
		q = q + fmt.Sprintf("&tracker_id=%v", filter.ParentId)
	}

	return q
}
