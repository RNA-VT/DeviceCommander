package utilities

import "encoding/json"

//LabelString -
func LabelString(key string, value string) string {
	return "\n\t[" + key + "]:     " + value
}

//JSON Marshals an object into a byte array
func JSON(obj interface{}) (out []byte, err error) {
	out, err = json.Marshal(obj)
	if err != nil {
		return
	}
	return
}

//StringJSON returns a stringified of a json string representing obj
func StringJSON(obj interface{}) (out string, err error) {
	bytes, err := JSON(obj)
	if err != nil {
		out = "ObjectFailedToParse"
		return
	}
	return string(bytes), err
}
