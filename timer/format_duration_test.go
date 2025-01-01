package timer

import (
	"testing"
	"time"
)

func TestFormatDuration(test *testing.T) {
	unitTests := []struct {
		duration  time.Duration
		precision int
		expected  string
	}{
		{time.Nanosecond, 0, "1ns"},
		{time.Microsecond, 0, "1µs"},
		{time.Millisecond, 0, "1ms"},
		{time.Second, 0, "1s"},
		{time.Minute, 0, "1min"},
		{time.Hour, 0, "1h"},
		{time.Nanosecond, 3, "1.000ns"},
		{time.Microsecond, 3, "1.000µs"},
		{time.Millisecond, 3, "1.000ms"},
		{time.Second, 3, "1.000s"},
		{time.Minute, 3, "1.000min"},
		{time.Hour, 3, "1.000h"},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.expected, func(test *testing.T) {
			result := formatDuration(unitTest.duration, unitTest.precision)
			if result != unitTest.expected {
				test.Errorf("Expected %s, got %s", unitTest.expected, result)
			}
		})
	}
}
