package element

// IsEmpty gibt an, ob das Element leer ist.
// Ein Element ist leer, wenn es kein Nachfolger-Element hat.
// D.h. wenn der Nachfolger-Pointer nil ist.
func (e *Element) IsEmpty() bool {
	if e.next == nil {
		return true
	}
	return false
}

// Value liefert den Wert des Elements.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Value() int {
	if e.IsEmpty() == true {
		panic("value for empty element requested")
	}
	return e.value
}

// Next liefert das Nachfolger-Element.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Next() *Element {
	if e.IsEmpty() == true {
		panic("value for empty element requested")
	}
	return e.next
}

// SetValue setzt den Wert des Elements.
// Falls das Element bisher leer war, wird es mit dem gegebenen Wert initialisiert.
func (e *Element) SetValue(value int) {
	if e.IsEmpty() == true {
		e.value = value
		e.next = &Element{}
	} else {
		e.value = value
	}
}
