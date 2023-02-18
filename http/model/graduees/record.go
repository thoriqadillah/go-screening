package graduees

type Record struct {
	Id                  int    `json:"_id"`
	Type                string `json:"type_of_course"`
	Rank                float32
	Sex                 string
	Number_of_graduates string `json:"no_of_graduates"`
	Year                string
	FullCount           string `json:"_full_count"`
}
