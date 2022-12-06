package utilities

import (
	"encoding/json"
	"fmt"
)

// LabelString -.
func LabelString(key string, value string) string {
	return "\n\t[" + key + "]:     " + value
}

// JSON Marshals an object into a byte array.
func JSON(obj interface{}) (out []byte, err error) {
	out, err = json.Marshal(obj)
	if err != nil {
		return
	}
	return
}

// StringJSON returns a stringified of a json string representing obj.
func StringJSON(obj interface{}) (out string, err error) {
	bytes, err := JSON(obj)
	if err != nil {
		out = "ObjectFailedToParse"
		return
	}
	return string(bytes), err
}

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func WrapError(outer, inner error) error {
	return fmt.Errorf("%w; "+outer.Error(), inner)
}
