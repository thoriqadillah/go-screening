package graduees

type Result struct {
	Resource_id string
	Fields      []Field
	Records     []Record
	Links       []Link
	Limit       int
	Total       int
}
