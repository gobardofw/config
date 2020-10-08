package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

type jsonConfig struct {
	Files []string
	json  string
	data  map[string]interface{}
}

func (c *jsonConfig) fetch(key string) (interface{}, bool, bool) {
	if v, ok := c.data[key]; ok {
		return v, false, true
	}

	if val := gjson.Get(c.json, key); val.Exists() {
		return val, true, true
	}

	return nil, false, false
}

// Load configurations
func (c *jsonConfig) Load() bool {
	c.data = make(map[string]interface{})
	contents := make([]string, 0)
	for _, f := range c.Files {
		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return false
		}
		content := string(bytes)
		if !gjson.Valid(content) {
			return false
		}

		fileName := filepath.Base(f)
		fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

		if len(c.Files) > 1 {
			contents = append(contents, `"`+fileName+`":`+content)
		} else {
			contents = append(contents, content)
		}

	}
	if len(c.Files) > 1 {
		c.json = "{" + strings.Join(contents, ",") + "}"
	} else {
		if !strings.HasPrefix(contents[0], "{") {
			contents[0] = "{" + contents[0] + "}"
		}
		c.json = contents[0]
	}
	return true
}

// Set configuration
// return false if driver not support set or error happend
func (c *jsonConfig) Set(key string, value interface{}) bool {
	c.data[key] = value
	return true
}

// Get configuration
func (c *jsonConfig) Get(key string) (interface{}, bool) {
	if v, isJSON, exists := c.fetch(key); !exists {
		return nil, false
	} else if isJSON {
		return v.(gjson.Result).Value(), true
	} else {
		return v, true
	}
}

// Exists check if env variable exists
func (c *jsonConfig) Exists(key string) bool {
	_, _, exists := c.fetch(key)
	return exists
}

// Bool parse dependency as boolean
func (c *jsonConfig) Bool(key string, fallback bool) (bool, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.False || val.Type == gjson.True {
				return val.Bool(), true
			}
		} else {
			if val, ok := v.(bool); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Int parse dependency as int
func (c *jsonConfig) Int(key string, fallback int) (int, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int(val.Int()), true
			}
		} else {
			if val, ok := v.(int); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Int8 parse dependency as int8
func (c *jsonConfig) Int8(key string, fallback int8) (int8, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int8(val.Int()), true
			}
		} else {
			if val, ok := v.(int8); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Int16 parse dependency as int16
func (c *jsonConfig) Int16(key string, fallback int16) (int16, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int16(val.Int()), true
			}
		} else {
			if val, ok := v.(int16); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Int32 parse dependency as int32
func (c *jsonConfig) Int32(key string, fallback int32) (int32, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int32(val.Int()), true
			}
		} else {
			if val, ok := v.(int32); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Int64 parse dependency as int64
func (c *jsonConfig) Int64(key string, fallback int64) (int64, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int64(val.Int()), true
			}
		} else {
			if val, ok := v.(int64); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// UInt parse dependency as uint
func (c *jsonConfig) UInt(key string, fallback uint) (uint, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint(val.Uint()), true
			}
		} else {
			if val, ok := v.(uint); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// UInt8 parse dependency as uint8
func (c *jsonConfig) UInt8(key string, fallback uint8) (uint8, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint8(val.Uint()), true
			}
		} else {
			if val, ok := v.(uint8); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// UInt16 parse dependency as uint16
func (c *jsonConfig) UInt16(key string, fallback uint16) (uint16, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint16(val.Uint()), true
			}
		} else {
			if val, ok := v.(uint16); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// UInt32 parse dependency as uint32
func (c *jsonConfig) UInt32(key string, fallback uint32) (uint32, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint32(val.Uint()), true
			}
		} else {
			if val, ok := v.(uint32); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// UInt64 parse dependency as uint64
func (c *jsonConfig) UInt64(key string, fallback uint64) (uint64, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint64(val.Uint()), true
			}
		} else {
			if val, ok := v.(uint64); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Float32 parse dependency as float64
func (c *jsonConfig) Float32(key string, fallback float32) (float32, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return float32(val.Float()), true
			}
		} else {
			if val, ok := v.(float32); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// Float64 parse dependency as float64
func (c *jsonConfig) Float64(key string, fallback float64) (float64, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return val.Float(), true
			}
		} else {
			if val, ok := v.(float64); ok {
				return val, true
			}
		}
	}
	return fallback, false
}

// String parse dependency as string
func (c *jsonConfig) String(key string, fallback string) (string, bool) {
	if v, isJSON, exists := c.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.String {
				return val.String(), true
			}
		} else {
			if val, ok := v.(string); ok {
				return val, true
			}
		}
	}
	return fallback, false
}
