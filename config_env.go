package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type envConfig struct {
	Files []string
}

// Load configurations
func (c *envConfig) Load() bool {
	if godotenv.Overload(c.Files...) != nil {
		return false
	}
	return true
}

// Get configuration
func (c *envConfig) Get(key string) (interface{}, bool) {
	if v, ok := os.LookupEnv(key); ok {
		return v, true
	}
	return nil, false
}

// Exists check if env variable exists
func (c *envConfig) Exists(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

// Bool parse dependency as boolean
func (c *envConfig) Bool(key string, fallback bool) (bool, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseBool(val); err == nil {
			return res, true
		}
	}
	return fallback, false
}

// Int parse dependency as int
func (c *envConfig) Int(key string, fallback int) (int, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int(res), true
		}
	}
	return fallback, false
}

// Int8 parse dependency as int8
func (c *envConfig) Int8(key string, fallback int8) (int8, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int8(res), true
		}
	}
	return fallback, false
}

// Int16 parse dependency as int16
func (c *envConfig) Int16(key string, fallback int16) (int16, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int16(res), true
		}
	}
	return fallback, false
}

// Int32 parse dependency as int32
func (c *envConfig) Int32(key string, fallback int32) (int32, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int32(res), true
		}
	}
	return fallback, false
}

// Int64 parse dependency as int64
func (c *envConfig) Int64(key string, fallback int64) (int64, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int64(res), true
		}
	}
	return fallback, false
}

// UInt parse dependency as uint
func (c *envConfig) UInt(key string, fallback uint) (uint, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint(res), true
		}
	}
	return fallback, false
}

// UInt8 parse dependency as uint8
func (c *envConfig) UInt8(key string, fallback uint8) (uint8, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint8(res), true
		}
	}
	return fallback, false
}

// UInt16 parse dependency as uint16
func (c *envConfig) UInt16(key string, fallback uint16) (uint16, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint16(res), true
		}
	}
	return fallback, false
}

// UInt32 parse dependency as uint32
func (c *envConfig) UInt32(key string, fallback uint32) (uint32, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint32(res), true
		}
	}
	return fallback, false
}

// UInt64 parse dependency as uint64
func (c *envConfig) UInt64(key string, fallback uint64) (uint64, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint64(res), true
		}
	}
	return fallback, false
}

// Float32 parse dependency as float64
func (c *envConfig) Float32(key string, fallback float32) (float32, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseFloat(val, 64); err == nil {
			return float32(res), true
		}
	}
	return fallback, false
}

// Float64 parse dependency as float64
func (c *envConfig) Float64(key string, fallback float64) (float64, bool) {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseFloat(val, 64); err == nil {
			return res, true
		}
	}
	return fallback, false
}

// String parse dependency as string
func (c *envConfig) String(key string, fallback string) (string, bool) {
	if val, ok := os.LookupEnv(key); ok {
		return val, true
	}
	return fallback, false
}
