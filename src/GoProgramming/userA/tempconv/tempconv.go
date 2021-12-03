package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC 	  Celsius = 0
	BoilingC	  Celsius = 100
)

/*
*	可以认为是重载(但是这样对吗？)
 */
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°C", f)
}



