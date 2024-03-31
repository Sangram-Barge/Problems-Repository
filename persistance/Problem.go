package persistance

type Problem struct {
	Id          int    `json:"id"`
	Problem     string `json:"problem"`
	Platform    string `json:"platform"`
	Description string `json:"desc"`
	Intiution   string `json:"intiution"`
	Link        string `json:"link"`
}

func NewProblem() Problem {
	return Problem{}
}
