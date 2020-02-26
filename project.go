package redmine

type ProjectResult struct {
	Project Project `json:"project"`
}

type ProjectsResult struct {
	Projects   []*Project `json:"projects"`
	TotalCount int        `json:"total_count"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
}

type Project struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Identifier      string   `json:"identifier"`
	Description     string   `json:"description"`
	Parent          IdName   `json:"parent"`
	Homepage        string   `json:"homepage"`
	Status          int      `json:"status"`
	IsPublic        bool     `json:"is_public"`
	Trackers        []IdName `json:"trackers"`
	IssueCategories []IdName `json:"issue_categories"`
	EnabledModules  []IdName `json:"enabled_modules"`
	CreatedOn       string   `json:"created_on"`
	UpdatedOn       string   `json:"updated_on"`
}


