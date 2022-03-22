package utils

import (
	"encoding/json"
)

// ToString, which takes any interface as argument and return a json string from it
func ToString(i interface{}) (str string, err error) {
	b, err := json.Marshal(i)
	str = string(b)
	return
}

// ToStruct, which takes any json string as argument and return a mapped interface{} from it
func ToStruct(jsonStr string, mappedStruct interface{})(mapped interface{}, err error) {
	err = json.Unmarshal([]byte(jsonStr), &mappedStruct)
	if err != nil {
		err = NewError("ToStruct unmarshal error for the following json: %s, error: %s: ", jsonStr, err.Error())
	}
	mapped = mappedStruct
	return
}

// ToStructForBytes, which takes any json string as argument and return a mapped interface{} from it
func ToStructForBytes(jsonBytes []byte, mappedStruct interface{})(mapped interface{}, err error) {
	err = json.Unmarshal(jsonBytes, &mappedStruct)
	if err != nil {
		err = NewError("ToStructForBytes unmarshal error for the following json: %s, error: %s: ", string(jsonBytes), err.Error())
	}
	mapped = mappedStruct
	return
}