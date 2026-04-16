package element

// IsEmpty gibt an, ob das Element leer ist.
// Ein Element ist leer, wenn es kein Nachfolger-Element hat.
// D.h. wenn der Nachfolger-Pointer nil ist.
func (e *Element) IsEmpty() bool {
	// HINWEIS:
	// Überprüfe, ob der next-Pointer nil ist.
	// begin:solution
	return e.next == nil
	// end:solution
}

// Value liefert den Wert des Elements.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Value() int {
	// HINWEIS:
	// Überprüfe, ob das Element leer ist (mit IsEmpty).
	// Wenn ja, löse eine panic aus.
	// Ansonsten gib den value zurück.
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
	// HINWEIS:
	// Überprüfe, ob das Element leer ist (mit IsEmpty).
	// Wenn ja, löse eine panic aus.
	// Ansonsten gib den next-Pointer zurück.
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
	// HINWEIS:
	// Wenn das Element leer ist, initialisiere es:
	//   Erstelle ein neues leeres Element und weise es dem next-Pointer zu.
	//   Setze den value auf den gegebenen Wert.
	// Ansonsten setze einfach den value auf den gegebenen Wert.
	// begin:solution
	if e.IsEmpty() {
		e.next = Empty()
	}
	e.value = value
	// end:solution
}
