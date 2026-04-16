package element

// IsEmpty gibt an, ob das Element leer ist.
// Ein Element ist leer, wenn es kein Nachfolger-Element hat.
// D.h. wenn der Nachfolger-Pointer nil ist.
func (e *Element) IsEmpty() bool {
	// HINWEIS:
	// Überprüfen Sie, ob der next-Pointer nil ist.
	return e.next == nil
}

// Value liefert den Wert des Elements.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Value() int {
	// HINWEIS:
	// Überprüfen Sie, ob das Element leer ist (mit IsEmpty).
	// Wenn ja, lösen Sie eine panic aus.
	// Ansonsten geben Sie den value zurück.

	// TODO
	return 0
}

// Next liefert das Nachfolger-Element.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Next() *Element {
	// HINWEIS:
	// Überprüfen Sie, ob das Element leer ist (mit IsEmpty).
	// Wenn ja, lösen Sie eine panic aus.
	// Ansonsten geben Sie den next-Pointer zurück.

	// TODO
	return nil
}

// SetValue setzt den Wert des Elements.
// Falls das Element bisher leer war, wird es mit dem gegebenen Wert initialisiert.
func (e *Element) SetValue(value int) {
	// HINWEIS:
	// Wenn das Element leer ist, initialisieren Sie es:
	//   Erstellen Sie ein neues leeres Element und weise es dem next-Pointer zu.
	//   Setzen Sie den value auf den gegebenen Wert.
	// Ansonsten setzen Sie einfach den value auf den gegebenen Wert.

	// TODO
}
