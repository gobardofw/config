package config

// Config is the interface for configuration manager drivers.
type Config interface {
	// Load configurations
	Load() bool
	// Get configuration
	// return nil, false if value not exists
	Get(key string) (interface{}, bool)
	// Exists check if env variable exists
	Exists(key string) bool
	// Bool parse dependency as boolean
	// return fallback, false if value not exists
	Bool(key string, fallback bool) (bool, bool)
	// Int parse dependency as int
	// return fallback, false if value not exists
	Int(key string, fallback int) (int, bool)
	// Int8 parse dependency as int8
	// return fallback, false if value not exists
	Int8(key string, fallback int8) (int8, bool)
	// Int16 parse dependency as int16
	// return fallback, false if value not exists
	Int16(key string, fallback int16) (int16, bool)
	// Int32 parse dependency as int32
	// return fallback, false if value not exists
	Int32(key string, fallback int32) (int32, bool)
	// Int64 parse dependency as int64
	// return fallback, false if value not exists
	Int64(key string, fallback int64) (int64, bool)
	// UInt parse dependency as uint
	// return fallback, false if value not exists
	UInt(key string, fallback uint) (uint, bool)
	// UInt8 parse dependency as uint8
	// return fallback, false if value not exists
	UInt8(key string, fallback uint8) (uint8, bool)
	// UInt16 parse dependency as uint16
	// return fallback, false if value not exists
	UInt16(key string, fallback uint16) (uint16, bool)
	// UInt32 parse dependency as uint32
	// return fallback, false if value not exists
	UInt32(key string, fallback uint32) (uint32, bool)
	// UInt64 parse dependency as uint64
	// return fallback, false if value not exists
	UInt64(key string, fallback uint64) (uint64, bool)
	// Float32 parse dependency as float64
	// return fallback, false if value not exists
	Float32(key string, fallback float32) (float32, bool)
	// Float64 parse dependency as float64
	// return fallback, false if value not exists
	Float64(key string, fallback float64) (float64, bool)
	// String parse dependency as string
	// return fallback, false if value not exists
	String(key string, fallback string) (string, bool)
}
