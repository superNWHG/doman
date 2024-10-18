package data

import (
	"bytes"
	"encoding/json"
)

func decodeJson(data []byte) (map[string]*json.RawMessage, error) {
	var obj map[string]*json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func encodeJson(data map[string]interface{}) ([]byte, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func jsonToMap(obj map[string]*json.RawMessage) ([]string, []string, error) {
	keys := []string{}
	values := []string{}

	for key, rawMsg := range obj {
		var value string
		if err := json.Unmarshal(*rawMsg, &value); err != nil {
			return nil, nil, err
		}

		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values, nil
}

func formatJson(data []byte) ([]byte, error) {
	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "	"); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
