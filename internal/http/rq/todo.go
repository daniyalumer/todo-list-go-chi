package rq

type Todo struct {
	Description string `json:"description"`
}

type TodoUpdate struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
