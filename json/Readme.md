## JSON Marshal and Unmarshal

In Go, **marshaling** converts a Go value into JSON, while **unmarshaling** converts JSON into a Go value.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Asish", Email: "asish@example.com"}

	// Marshal: Go value -> JSON bytes
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Unmarshal: JSON bytes -> Go value
	var decoded User
	if err := json.Unmarshal(data, &decoded); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", decoded)
}
```

### Notes

- `json.Marshal(value)` returns `([]byte, error)`.
- `json.Unmarshal(data, &value)` requires a pointer so it can update the value.
- Only exported struct fields are processed.
- Use struct tags such as ``json:"name"`` to control JSON field names.
- Use `json:"-"` to omit a field, or `json:",omitempty"` to omit zero values.
