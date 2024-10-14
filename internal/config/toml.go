package config

import "github.com/BurntSushi/toml"

func encodeToml(values *config) ([]byte, error) {
	encodedToml, err := toml.Marshal(values)
	if err != nil {
		return nil, err
	}

	return encodedToml, nil
}

func decodeToml(data []byte, configStruct config) (*config, error) {
	if err := toml.Unmarshal(data, configStruct); err != nil {
		return nil, err
	}

	return &configStruct, nil
}
