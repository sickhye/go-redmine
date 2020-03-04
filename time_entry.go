package redmine

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

func (c *Client) CreateTimeEntry(t TimeEntry) error {
}
