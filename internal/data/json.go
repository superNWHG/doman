package data

import "encoding/json"

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
