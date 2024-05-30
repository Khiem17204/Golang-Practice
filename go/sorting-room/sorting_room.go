package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (f FancyNumber) Value() string {
	return f.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if fancyNumber, ok := fnb.(FancyNumber); ok {
		num, err := strconv.Atoi(fancyNumber.Value())
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if _, ok := fnb.(FancyNumber); ok {
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
	}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", 0.0)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch val := i.(type) {
	case float64:
		return DescribeNumber(val)
	case int:
		return DescribeNumber(float64(val))
	case NumberBox:
		return DescribeNumberBox(val)
	case FancyNumberBox:
		return DescribeFancyNumberBox(val)
	default:
		return "Return to sender"
	}
}
