package model

type Link struct {
	Start string
	Next  string
}

type Field struct {
	Type string
	Id   string
}

type Record struct {
	Id                  int    `json:"_id"`
	Type                string `json:"type_of_course"`
	Rank                float32
	Sex                 string
	Number_of_graduates string `json:"no_of_graduates"`
	Year                string
	FullCount           string `json:"_full_count"`
}

type Result struct {
	Resource_id string
	Fields      []Field
	Records     []Record
	Links       []Link
	Limit       int
	Total       int
}

type Data struct {
	Help    string
	Success bool
	Result  Result
}
