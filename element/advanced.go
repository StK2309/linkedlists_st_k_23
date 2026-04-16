package element

import "fmt"

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (e *Element) Append(value int) {
	// HINWEIS:
	// Wenn das Element leer ist, initialisieren Sie es mit dem gegebenen Wert
	// (mit SetValue).
	// Ansonsten rufen Sie Append rekursiv auf dem Nachfolger-Element auf.
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
	// Wenn das Element leer ist, geben Sie 0 zurück.
	// Ansonsten geben Sie 1 plus die Länge des Nachfolger-Elements zurück.
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
	// Wenn das Element leer ist, geben Sie false zurück.
	// Wenn der Wert des Elements dem gesuchten Wert entspricht, geben Sie true zurück.
	// Ansonsten rufen Sie Contains rekursiv auf das Nachfolger-Element auf.
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
	// Wenn das Element leer ist, geben Sie 0 zurück.
	// Wenn der Wert des Elements dem gesuchten Wert entspricht, geben Sie 1 plus die Anzahl im Nachfolger-Element zurück.
	// Ansonsten geben Sie die Anzahl im Nachfolger-Element zurück.
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
	// Wenn das Element leer ist, geben Sie 0 zurück.
	// Ansonsten geben Sie den Wert des Elements plus die Summe des Nachfolger-Elements zurück.
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
	// Wenn das Nachfolger-Element leer ist, geben Sie den Wert des Elements zurück.
	// Ansonsten geben Sie den kleineren Wert von Wert des Elements und Minimum des Nachfolger-Elements zurück.
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
	// Wenn das Nachfolger-Element leer ist, geben Sie das aktuelle Element zurück.
	// Ansonsten rufen Sie Last rekursiv auf das Nachfolger-Element auf.
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
	// Wenn die Position 0 ist, geben Sie das aktuelle Element zurück.
	// Ansonsten rufen Sie At rekursiv auf das Nachfolger-Element mit position - 1 auf.
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
	// Wenn das Element leer ist, geben Sie einen leeren String zurück.
	// Wenn das Nachfolger-Element leer ist, geben Sie den Wert des Elements als String zurück.
	// Ansonsten geben Sie den Wert des Elements gefolgt von " -> " und der String-Repräsentation des Nachfolger-Elements zurück.
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

// Swap vertauscht die beiden Elemente an den Stellen i und j.
// Falls eine der Positionen nicht existiert, wird eine panic ausgelöst.
// Die Funktion liefert den neuen Kopf der Liste zurück.

// Es sollen nicht die Werte der Elemente vertauscht werden, sondern die Elemente selbst.
func (e *Element) Swap(i, j int) *Element {
	// HINWEIS:
	// - Rufen Sie At auf, um die Elemente an den Positionen i-1 und j-1 zu erhalten.
	// - Bestimmen Sie dann auch die Elemente an den Positionen i, i+1, j und j+1.
	// - Tauschen Sie dann die next-Pointer der Elemente an den Positionen
	//   i-1, i, j-1 und j entsprechend aus.

	// ZUSATZHINWEISE:
	// - Wenn i und j gleich sind, müssen Sie nichts tun.
	// - Wenn i größer als j ist, können Sie einfach i und j vertauschen, um die Logik zu vereinfachen.
	// - Wenn i = 0 ist, dann ist das ein Sonderfall, da es kein Element an Position i-1 gibt.
	//   In diesem Fall müssen Sie den Kopf der Liste aktualisieren, damit er auf das neue erste Element zeigt.
	// - Falls i und j direkt benachbart sind, ist auch das ein Sonderfall.

	// begin:solution
	if i == j {
		return e
	}
	if i > j {
		i, j = j, i
	}

	if i == 0 {
		// Visualisierung der Liste:
		// A -> B -> ... -> D -> E -> F -> ...
		// ^                     ^
		// i                     j
		A := e
		B := A.Next()
		D := e.At(j - 1)
		E := D.Next()
		F := E.Next()

		D.next = A
		A.next = F
		E.next = B
		return E
	}
	if j == i+1 {
		// Visualisierung der Liste:
		// ... -> A -> B -> C -> D -> ...
		//             ^    ^
		//             i    j
		A := e.At(i - 1)
		B := A.Next()
		C := B.Next()
		D := C.Next()

		A.next = C
		C.next = B
		B.next = D
		return e
	}

	// Visualisierung der Liste:
	// ... -> A -> B -> C -> ... -> D -> E -> F -> ...
	//             ^                     ^
	//             i                     j
	A := e.At(i - 1)
	B := A.Next()
	C := B.Next()
	D := e.At(j - 1)
	E := D.Next()
	F := E.Next()

	A.next = E
	E.next = C
	D.next = B
	B.next = F
	return e

	// end:solution
}
