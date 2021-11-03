package types

import (
	"time"
)

// MustParseRFC3339 parses string time to time in RFC3339 format.
func MustParseRFC3339(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func DateRageOverlap(startTimeA, endTimeA, startTimeB, endTimeB time.Time) bool {
	return !startTimeA.After(endTimeB) && !endTimeA.Before(startTimeB)
}
