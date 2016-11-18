// tempconv.go Convert between temperature scales.

package tempconv

import "fmt"

// Celsius Defines float64 type for Celsius unit of temperature
type Celsius float64

// Fahrenheit Defines float64 type for Fahrenheit unit of temperature
type Fahrenheit float64

const (
	// AbsoluteZeroC Absolute Zero in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC Freezing point of water in Celsius
	FreezingC Celsius = 0
	// BoilingC Boiling point of water in Celsius
	BoilingC Celsius = 100
)

// CToF Convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Define how value is printed as a string
func (c Celsius) String() string    { return fmt.Sprintf("%.5g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.5g°F", f) }
