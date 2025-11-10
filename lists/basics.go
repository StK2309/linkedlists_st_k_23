package lists

import (
	"fmt"
	"linkedlists/elements"
)

// New erstellt eine neue leere Liste.
func New() *List {
	return &List{head: elements.NewEmpty()}
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
	// HINWEIS:
	// Delegiere die Arbeit an das Kopf-Element der Liste.
	// D.d. rufe die Append-Methode des Kopf-Elements auf.
	// begin:solution
	l.head.Append(value)
	// end:solution
}

// Length gibt die Anzahl der Elemente in der Liste zurück.
func (l *List) Length() int {
	// HINWEIS:
	// Delegiere die Arbeit an das Kopf-Element der Liste.
	// begin:solution
	return l.head.Length()
	// end:solution
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (l *List) Contains(value int) bool {
	// HINWEIS:
	// Delegiere die Arbeit an das Kopf-Element der Liste.
	// begin:solution
	return l.head.Contains(value)
	// end:solution
}
