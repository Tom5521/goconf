package conf

// Preferences describes the ways that an app can save and load user preferences
type Preferences interface {
	Read(key string) any
	Set(key string, value any)
	// Bool looks up a bool value for the key
	Bool(key string) bool

	// SetBool saves a bool value for the given key
	SetBool(key string, value bool)

	// BoolList looks up a list of bool values for the key
	BoolList(key string) []bool

	// SetBoolList saves a list of bool values for the given key
	SetBoolList(key string, value []bool)

	// Float looks up a float64 value for the key
	Float(key string) float64

	// SetFloat saves a float64 value for the given key
	SetFloat(key string, value float64)

	// FloatList looks up a list of float64 values for the key
	FloatList(key string) []float64

	// SetFloatList saves a list of float64 values for the given key
	SetFloatList(key string, value []float64)

	// Int looks up an integer value for the key
	Int(key string) int

	// SetInt saves an integer value for the given key
	SetInt(key string, value int)

	// IntList looks up a list of int values for the key
	IntList(key string) []int

	// SetIntList saves a list of string values for the given key
	SetIntList(key string, value []int)

	// String looks up a string value for the key
	String(key string) string

	// SetString saves a string value for the given key
	SetString(key string, value string)

	// StringList looks up a list of string values for the key
	StringList(key string) []string

	// SetStringList saves a list of string values for the given key
	SetStringList(key string, value []string)

	// RemoveValue removes a value for the given key
	RemoveValue(key string)
}
