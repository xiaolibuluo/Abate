/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"encoding/json"
	"fmt"
)

// Marshal renders one sandbox configuration as canonical JSON.
//
// The encoding is deterministic: the standard library writes the struct fields
// in declaration order and the allow lists are sorted by Canonical, so the
// same configuration always yields the same bytes.
func Marshal(config Config) ([]byte, error) {
	data, err := json.Marshal(config.Canonical())
	if err != nil {
		return nil, fmt.Errorf("sandbox: encode configuration: %w", err)
	}
	return data, nil
}

// Unmarshal decodes one canonical JSON sandbox configuration.
func Unmarshal(data []byte) (Config, error) {
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("sandbox: decode configuration: %w", err)
	}
	return config, nil
}

// MarshalIndent renders one sandbox configuration as indented canonical JSON.
//
// The report form keeps the same field order as Marshal, so both encodings
// describe the same configuration.
func MarshalIndent(config Config) ([]byte, error) {
	data, err := json.MarshalIndent(config.Canonical(), "", "  ")
	if err != nil {
		return nil, fmt.Errorf("sandbox: encode configuration: %w", err)
	}
	return data, nil
}
