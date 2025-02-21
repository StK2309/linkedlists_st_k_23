package elements

// IsEmpty gibt an, ob das Element leer ist.
// Ein Element ist leer, wenn es kein Nachfolger-Element hat.
// D.h. wenn der Nachfolger-Pointer nil ist.
func (e *Element) IsEmpty() bool {
	// begin:solution
	return e.next == nil
	// end:solution
}

// Value liefert den Wert des Elements.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Value() int {
	// begin:solution
	if e.IsEmpty() {
		panic("value for empty element requested")
	}
	return e.value
	// end:solution
}

// Next liefert das Nachfolger-Element.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Next() *Element {
	// begin:solution
	if e.IsEmpty() {
		panic("next element for empty element requested")
	}
	return e.next
	// end:solution
}

// SetValue setzt den Wert des Elements.
// Falls das Element bisher leer war, wird es mit dem gegebenen Wert initialisiert.
func (e *Element) SetValue(value int) {
	// begin:solution
	if e.IsEmpty() {
		e.next = NewEmpty()
	}
	e.value = value
	// end:solution
}
