package builddate

import "C"
import (
	"time"
)

func DateString() string {
	return C.__DATE__
}

func TimeString() string {
	return C.__TIME__
}

func BuildString() string {
	return DateString() + " " + TimeString()
}

func TimeLayout() string {
	return "Jan 2 2006 15:04:05"
}

func BuildTime() (time.Time, error) {
	return time.Parse(TimeLayout(), BuildString())
}
