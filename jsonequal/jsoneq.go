package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	helloraj := []byte(`{"hello":"raj","world":"testing"}`)
	var interface1, interface2 interface{}
	json.Unmarshal(helloraj, &interface1)
	hellorajinanotherwayr := []byte(`{"world":"testing","hello":"raj"}`)
	json.Unmarshal(hellorajinanotherwayr, &interface2)
	fmt.Println(interface2)
	fmt.Println(interface1)
	fmt.Println(isJSONEqual(interface1, interface2))
}

//Below code was copied from on of the lib
func isJSONEqual(expectedVal, actualVal interface{}) bool {
	if reflect.TypeOf(expectedVal) != reflect.TypeOf(actualVal) {
		return false
	}
	switch x := expectedVal.(type) {
	case map[string]interface{}:
		y := actualVal.(map[string]interface{})
		if len(x) != len(y) {
			return false
		}

		for k, v := range x {
			val2 := y[k]
			if (v == nil) != (val2 == nil) {
				return false
			}
			if !isJSONEqual(v, val2) {
				return false
			}
		}
		return true

	case []interface{}:
		y := actualVal.([]interface{})
		if len(x) != len(y) {
			return false
		}
		var matches int
		flagged := make([]bool, len(y))
		for _, v := range x {
			for i, v2 := range y {
				if isJSONEqual(v, v2) && !flagged[i] {
					matches++
					flagged[i] = true
					break
				}
			}
		}
		return matches == len(x)
	default:
		return expectedVal == actualVal
	}
}
