// Package leap provides utilities for detecting leap years
package leap

// IsLeapYear returns whether an input year is a leap year - A year divisible
// by four but not 100 unless also divisible by 400.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
