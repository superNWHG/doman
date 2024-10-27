package config

import "github.com/BurntSushi/toml"

func encodeToml(values interface{}) ([]byte, error) {
	encodedToml, err := toml.Marshal(values)
	if err != nil {
		return nil, err
	}

	return encodedToml, nil
}

func decodeToml(data []byte, configStruct interface{}) (interface{}, error) {
	if err := toml.Unmarshal(data, configStruct); err != nil {
		return nil, err
	}

	return configStruct, nil
}

func decodeTomlAny(data []byte, configStruct any) (any, error) {
	if err := toml.Unmarshal(data, configStruct); err != nil {
		return nil, err
	}

	return configStruct, nil
}
