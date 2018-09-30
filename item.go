package main

type (
	i_item struct {
		x, y       int
		appearance rune
		name       string
	}
)

func (i *i_item) getType() string {
	return "item"
}
