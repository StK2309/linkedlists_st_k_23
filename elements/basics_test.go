package elements

import "fmt"

func ExampleElement_IsEmpty() {
	e1 := NewEmpty()
	e2 := NewWithValue(42)

	fmt.Println(e1.IsEmpty())
	fmt.Println(e2.IsEmpty())

	// Output:
	// true
	// false
}

func ExampleElement_Value_nonempty() {
	e := NewWithValue(42)

	fmt.Println(e.Value())

	// Output:
	// 42
}

func ExampleElement_Value_empty() {
	e := NewEmpty()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println(e.Value())

	// Output:
	// value for empty element requested
}

func ExampleElement_Next_nonempty() {
	e1 := NewWithValue(42)
	e1.next = NewWithValue(23)

	fmt.Println(e1.Next().value)
	fmt.Println(e1.Next().IsEmpty())

	// Output:
	// 23
	// false
}

func ExampleElement_SetValue() {
	e := NewEmpty()

	e.SetValue(42)
	fmt.Println(e.Value())
	fmt.Println(e.IsEmpty())

	e.SetValue(23)
	fmt.Println(e.Value())
	fmt.Println(e.IsEmpty())

	// Output:
	// 42
	// false
	// 23
	// false
}
