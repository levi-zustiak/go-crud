package entities

type Book struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
}
