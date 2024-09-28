package models

type Page struct {
	Products    []Product
	CurrentPage int
	TotalPages  int
	Pages       []int
}
