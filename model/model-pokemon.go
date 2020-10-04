package model

type Pokemon struct {
	ID      string `bson:"_id" json:"id"`
	Name    string `json:"name"`
	Element string `json:"element"`
	Weight  string `json:"weight"`
}
