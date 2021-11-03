package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/budget/x/budget/types"
)

func TestParseTime(t *testing.T) {
	normalCase := "9999-12-31T00:00:00Z"
	normalRes, err := time.Parse(time.RFC3339, normalCase)
	require.NoError(t, err)
	errorCase := "9999-12-31T00:00:00_ErrorCase"
	_, err = time.Parse(time.RFC3339, errorCase)
	require.PanicsWithError(t, err.Error(), func() { types.ParseTime(errorCase) })
	require.Equal(t, normalRes, types.ParseTime(normalCase))
}

func TestDateRageOverlap(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		startTimeA     time.Time
		endTimeA       time.Time
		startTimeB     time.Time
		endTimeB       time.Time
	}{
		{
			"same range",
			true,
			types.ParseTime("2021-12-31T00:00:00Z"),
			types.ParseTime("2021-12-31T00:00:00Z"),
			types.ParseTime("2021-12-31T00:00:00Z"),
			types.ParseTime("2021-12-31T00:00:00Z"),
		},
		{
			"overlap with start",
			true,
			types.ParseTime("2021-10-05T00:00:00Z"),
			types.ParseTime("2021-12-31T00:00:00Z"),
			types.ParseTime("2021-10-05T00:00:00Z"),
			types.ParseTime("2021-11-10T00:00:00Z"),
		},
		{
			"overlap with start 2",
			true,
			types.ParseTime("2021-10-05T00:00:00Z"),
			types.ParseTime("2021-11-10T00:00:00Z"),
			types.ParseTime("2021-10-05T00:00:00Z"),
			types.ParseTime("2021-12-31T00:00:00Z"),
		},
		{
			"overlap 1 sec",
			true,
			types.ParseTime("2021-10-05T00:00:00Z"),
			types.ParseTime("2021-11-10T00:00:01Z"),
			types.ParseTime("2021-11-10T00:00:00Z"),
			types.ParseTime("2021-12-31T00:00:00Z"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedResult, types.DateRageOverlap(tc.startTimeA, tc.endTimeA, tc.startTimeB, tc.endTimeB))
		})
	}
}
