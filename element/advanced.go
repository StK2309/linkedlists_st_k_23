package element

import "fmt"

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (e *Element) Append(value int) {
	// HINWEIS:
	// Wenn das Element leer ist, initialisiere es mit dem gegebenen Wert (mit SetValue).
	// Ansonsten rufe Append rekursiv auf das Nachfolger-Element auf.
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
	// HINWEIS:
	// Wenn das Element leer ist, gib 0 zurück.
	// Ansonsten gib 1 plus die Länge des Nachfolger-Elements zurück.
	// begin:solution
	if e.IsEmpty() {
		return 0
	}
	return 1 + e.Next().Length()
	// end:solution
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (e *Element) Contains(value int) bool {
	// HINWEIS:
	// Wenn das Element leer ist, gib false zurück.
	// Wenn der Wert des Elements dem gesuchten Wert entspricht, gib true zurück.
	// Ansonsten rufe Contains rekursiv auf das Nachfolger-Element auf.
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
	// HINWEIS:
	// Wenn das Element leer ist, gib 0 zurück.
	// Wenn der Wert des Elements dem gesuchten Wert entspricht, gib 1 plus die Anzahl im Nachfolger-Element zurück.
	// Ansonsten gib die Anzahl im Nachfolger-Element zurück.
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
	// HINWEIS:
	// Wenn das Element leer ist, gib 0 zurück.
	// Ansonsten gib den Wert des Elements plus die Summe des Nachfolger-Elements zurück.
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
	// HINWEIS:
	// Wenn das Element leer ist, löse eine panic aus.
	// Wenn das Nachfolger-Element leer ist, gib den Wert des Elements zurück.
	// Ansonsten gib den kleineren Wert von Wert des Elements und Minimum des Nachfolger-Elements zurück.
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
	// HINWEIS:
	// Wenn das Element leer ist, löse eine panic aus.
	// Wenn das Nachfolger-Element leer ist, gib das aktuelle Element zurück.
	// Ansonsten rufe Last rekursiv auf das Nachfolger-Element auf.
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
	// HINWEIS:
	// Wenn das Element leer ist, löse eine panic aus.
	// Wenn die Position 0 ist, gib das aktuelle Element zurück.
	// Ansonsten rufe At rekursiv auf das Nachfolger-Element mit position - 1 auf.
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
	// HINWEIS:
	// Wenn das Element leer ist, gib einen leeren String zurück.
	// Wenn das Nachfolger-Element leer ist, gib den Wert des Elements als String zurück.
	// Ansonsten gib den Wert des Elements gefolgt von " -> " und der String-Repräsentation des Nachfolger-Elements zurück.
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
