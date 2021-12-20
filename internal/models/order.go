package models

type Order struct {
	ID         uint64
	ProductIDs []uint64
	OrderedAt  string
}
