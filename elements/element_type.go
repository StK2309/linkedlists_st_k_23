package elements

// Ein Element einer einfach verketteten Liste.
type Element struct {
	value int
	next  *Element
}

// NewEmpty liefert ein neues leeres Element.
func NewEmpty() *Element {
	return &Element{value: 0, next: nil}
}

// NewWithValue liefert ein neues Element mit dem gegebenen Wert.
func NewWithValue(value int) *Element {
	return &Element{value: value, next: NewEmpty()}
}
