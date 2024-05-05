package utils

import "encoding/json"

func ParseJSON[T any](b []byte) (*T, error) {
	var result T

	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
