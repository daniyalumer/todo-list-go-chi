package rq

type TodoCreate struct {
	Description string `json:"description"`
}

type TodoUpdate struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
