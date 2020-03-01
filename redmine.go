package redmine

type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Id struct {
	Id int `json:"id"`
}

type errorResult struct {
	Errors []string `json:"errors"`
}
