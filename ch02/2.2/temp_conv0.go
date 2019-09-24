package main

type Celsius float64
type Fahrenheit float64

const (
	// AbsoluteZeroC absolute zero tempreature
	AbsoluteZeroC Celsius = -273.15

	// FreezingC freeze zero tempreature
	FreezingC Celsius = 0

	// BoilingC bloing tempreature
	BoilingC Celsius = 100
)

// CToF convert
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC convert
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
