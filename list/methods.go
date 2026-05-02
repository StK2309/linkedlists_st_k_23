package list

import (
	"fmt"
	"linkedlists/element"
)

// Empty erstellt eine neue leere Liste.
func Empty() *List {
	return &List{head: element.Empty()}
}

// IsEmpty gibt an, ob die Liste leer ist.
func (l *List) IsEmpty() bool {
	return l.head.IsEmpty()
}

// String liefert eine menschlich lesbare Darstellung der Liste.
func (l *List) String() string {
	return fmt.Sprintf("[%v]", l.head)
}

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (l *List) Append(value int) {
	l.head.Append(value)
}

// Length gibt die Anzahl der Elemente in der Liste zurück.
func (l *List) Length() int {
	return l.head.Length()
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (l *List) Contains(value int) bool {
	return l.head.Contains(value)
}

// Swap vertauscht die Elemente an den gegebenen Positionen in der Liste.
func (l *List) Swap(pos1, pos2 int) {
	l.head = l.head.Swap(pos1, pos2)
}
