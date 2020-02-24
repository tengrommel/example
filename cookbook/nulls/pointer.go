package nulls

import (
	"encoding/json"
	"fmt"
)

// ExamplePointer is the same, but uses a *Int
type ExamplePointer struct {
	Age  *int   `json:"age, omitempty"`
	Name string `json:"name"`
}

// PointerEncoding shows methods for dealing with nil/omitted values
func PointerEncoding() error {
	// note that no arg = nil age
	e := ExamplePointer{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Marshal, with no age: %+v\n", e)
	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with no age:", string(value))
	if err := json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	value, err = json.Marshal(&e)
	fmt.Println("Pointer Marshal, with age = 0:", string(value))
	return nil
}
