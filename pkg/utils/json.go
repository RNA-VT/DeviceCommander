package utils

import "encoding/json"

func PrettyPrintJSON(data interface{}) string {
	byteArray, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic("failed to marshal basic json")
	}
	return string(byteArray)
}

func PrintJSON(data interface{}) string {
	byteArray, err := json.Marshal(data)
	if err != nil {
		panic("failed to marshal basic json")
	}
	return string(byteArray)
}
