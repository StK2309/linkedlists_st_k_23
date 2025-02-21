package elements

import "fmt"

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (e *Element) Append(value int) {
	// begin:solution
	if e.IsEmpty() {
		e.SetValue(value)
	} else {
		e.Next().Append(value)
	}
	// end:solution
}

// Length gibt die Anzahl der Elemente in der Liste zurück.
func (e *Element) Length() int {
	// begin:solution
	if e.IsEmpty() {
		return 0
	}
	return 1 + e.Next().Length()
	// end:solution
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (e *Element) Contains(value int) bool {
	// begin:solution
	if e.IsEmpty() {
		return false
	}
	if e.Value() == value {
		return true
	}
	return e.Next().Contains(value)
	// end:solution
}

// Count gibt die Anzahl der Elemente in der Liste zurück, die den gegebenen Wert enthalten.
func (e *Element) Count(value int) int {
	// begin:solution
	if e.IsEmpty() {
		return 0
	}

	if e.Value() == value {
		return 1 + e.Next().Count(value)
	}
	return e.Next().Count(value)
	// end:solution
}

// Sum berechnet die Summe der Werte aller Elemente in der Liste.
func (e *Element) Sum() int {
	// begin:solution
	if e.IsEmpty() {
		return 0
	}
	return e.Value() + e.Next().Sum()
	// end:solution
}

// Min gibt den kleinsten Wert aller Elemente in der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Min() int {
	// begin:solution
	if e.IsEmpty() {
		panic("min for empty list requested")
	}
	if e.Next().IsEmpty() {
		return e.Value()
	}
	min := e.Next().Min()
	if e.Value() < min {
		return e.Value()
	}
	return min
	// end:solution
}

// Last gibt das letzte Element der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Last() *Element {
	// begin:solution
	if e.IsEmpty() {
		panic("last for empty list requested")
	}
	if e.Next().IsEmpty() {
		return e
	}
	return e.Next().Last()
	// end:solution
}

// At gibt das Element an der gegebenen Position zurück.
// Falls die Position nicht existiert, wird eine panic ausgelöst.
func (e *Element) At(position int) *Element {
	// begin:solution
	if e.IsEmpty() {
		panic("index out of bounds")
	}
	if position == 0 {
		return e
	}
	return e.Next().At(position - 1)
	// end:solution
}

// String gibt eine textuelle Repräsentation der Liste zurück.
// Die Elemente werden durch " -> " getrennt.
// Falls die Liste leer ist, wird ein leerer String zurückgegeben.
func (e *Element) String() string {
	// begin:solution
	if e.IsEmpty() {
		return ""
	}
	if e.Next().IsEmpty() {
		return fmt.Sprintf("%d", e.Value())
	}
	return fmt.Sprintf("%d -> %s", e.Value(), e.Next().String())
	// end:solution
}
