package redmine

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type projectResult struct {
	Project Project `json:"project"`
}

type projectsResult struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
}

type Project struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Identifier      string    `json:"identifier"`
	Description     string    `json:"description"`
	Parent          *IdName   `json:"parent"`
	Homepage        string    `json:"homepage"`
	Status          int       `json:"status"`
	IsPublic        bool      `json:"is_public"`
	Trackers        []*IdName `json:"trackers"`
	IssueCategories []*IdName `json:"issue_categories"`
	EnabledModules  []*IdName `j	qson:"enabled_modules"`
	CreatedOn       string    `json:"created_on"`
	UpdatedOn       string    `json:"updated_on"`
}

func (c *Client) GetProject(project string) (*Project, error) {
	res, err := c.Get(c.endpoint + "/projects/" + project + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r projectResult
	if res.StatusCode != http.StatusOK {
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

	return &r.Project, nil
}

func (c *Client) GetProjects() ([]Project, error) {
	// Todo: Limit, Offset settings
	res, err := c.Get(c.endpoint + "/projects.json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r projectsResult

	if res.StatusCode != http.StatusOK {
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

	return r.Projects, nil

}
