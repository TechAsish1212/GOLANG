	// Structs---->A struct (short for "structure") is a composite data type in Go that groups together zero or more fields with varying data types. It allows you to create complex data structures that can represent real-world entities or concepts. Structs are defined using the `type` keyword, and you can create instances of a struct to store and manipulate data.





In Go, when you define a method with a value receiver (like `func (o Order) changeStatus(status string)`), it operates on a copy of the struct. This means that any changes made to the struct's fields within the method will not affect the original struct outside the method. In this case, when you call `myOrder1.changeStatus("new status")`, it will change the status of the copy of `myOrder1` inside the method, but the original `myOrder1` will remain unchanged. To modify the original struct, you would need to use a pointer receiver (e.g., `func (o *Order) changeStatus(status string)`), which allows you to modify the original struct through its memory address.


	// syntax for defining a struct
	// type StructName struct {
	// 	Field1 DataType1
	// 	Field2 DataType2
	// 	...
	// }