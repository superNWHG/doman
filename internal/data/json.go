package data

import "encoding/json"

func decodeJson(data []byte) (error, map[string]*json.RawMessage) {
	var obj map[string]*json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return err, nil
	}

	return nil, obj
}
