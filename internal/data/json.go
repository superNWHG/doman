package data

import "encoding/json"

func decodeJson(data []byte) (error, map[string]*json.RawMessage) {
	var obj map[string]*json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return err, nil
	}

	return nil, obj
}

func encodeJson(data map[string]interface{}) (error, []byte) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err, nil
	}

	return nil, jsonString
}
