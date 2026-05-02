package element

import "fmt"

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (e *Element) Append(value int) {
	if e.IsEmpty() {
		e.SetValue(value)
	} else {
		e.next.Append(value)
	}
	// e.Last().next.SetValue(value)
}

// Length gibt die Anzahl der Elemente in der Liste zurück.
func (e *Element) Length() int {
	if e.IsEmpty() {
		return 0
	}
	return 1 + e.next.Length()
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (e *Element) Contains(value int) bool {
	if e.IsEmpty() {
		return false
	}
	if e.Value() == value {
		return true
	}
	return e.next.Contains(value)
}

// Count gibt die Anzahl der Elemente in der Liste zurück, die den gegebenen Wert enthalten.
func (e *Element) Count(value int) int {
	if e.IsEmpty() {
		return 0
	}
	count := 0
	if e.Value() == value {
		count = 1
	}
	return count + e.next.Count(value)
}

// Sum berechnet die Summe der Werte aller Elemente in der Liste.
func (e *Element) Sum() int {
	if e.IsEmpty() {
		return 0
	}
	return e.Value() + e.next.Sum()
}

// Min gibt den kleinsten Wert aller Elemente in der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Min() int {
	if e.IsEmpty() {
		panic("min for empty element requested")
	}
	min := e.Value()
	if !e.next.IsEmpty() {
		nextMin := e.next.Min()
		if nextMin < min {
			min = nextMin
		}
	}
	return min
}

// Last gibt das letzte Element der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Last() *Element {
	if e.IsEmpty() {
		panic("last for empty element requested")
	}
	if e.next.IsEmpty() {
		return e
	}
	return e.next.Last()
}

// At gibt das Element an der gegebenen Position zurück.
// Falls die Position nicht existiert, wird eine panic ausgelöst.
func (e *Element) At(position int) *Element {
	if position < 0 {
		panic("negative position requested")
	}
	if e.IsEmpty() {
		panic("index out of bounds")
	}
	if position == 0 {
		return e
	}
	return e.next.At(position - 1)
}

// String gibt eine textuelle Repräsentation der Liste zurück.
// Die Elemente werden durch " -> " getrennt.
// Falls die Liste leer ist, wird ein leerer String zurückgegeben.
func (e *Element) String() string {
	if e.IsEmpty() {
		return ""
	}
	if e.next.IsEmpty() {
		return fmt.Sprintf("%d", e.Value())
	}
	return fmt.Sprintf("%d -> %s", e.Value(), e.next.String())
}

// Swap vertauscht die beiden Elemente an den Stellen i und j.
// Falls eine der Positionen nicht existiert, wird eine panic ausgelöst.
// Die Funktion liefert den neuen Kopf der Liste zurück.

// Es sollen nicht die Werte der Elemente vertauscht werden, sondern die Elemente selbst.
func (e *Element) Swap(i, j int) *Element {
	if i < 0 || j < 0 {
		panic("negative position requested")
	}
	if e.IsEmpty() {
		panic("index out of bounds")
	}
	if i == j {
		return e
	}
	var prevI, prevJ, nodeI, nodeJ *Element = nil, nil, nil, nil
	current := e
	position := 0

	for current != nil {
		if position == i {
			nodeI = current
		}
		if position == j {
			nodeJ = current
		}
		if nodeI == nil {
			prevI = current
		}
		if nodeJ == nil {
			prevJ = current
		}
		current = current.next
		position++
	}
	if nodeI == nil || nodeJ == nil {
		panic("index out of bounds")
	}

	if prevI != nil {
		prevI.next = nodeJ
	} else {
		e = nodeJ
	}
	if prevJ != nil {
		prevJ.next = nodeI
	} else {
		e = nodeI
	}

	nodeI.next, nodeJ.next = nodeJ.next, nodeI.next

	return e
}
