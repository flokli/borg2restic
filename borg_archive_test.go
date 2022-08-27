package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseBorgTimestamps(t *testing.T) {
	tExpected := time.Date(
		2022,
		8,
		12,
		21,
		19,
		37,
		878526000,
		time.Local,
	)

	t.Run("new timestamp", func(t *testing.T) {
		str := "2022-08-12T21:19:37.878526+02:00"
		ts, err := ParseBorgTimestamp(str)

		if assert.NoError(t, err) {
			assert.Equal(t, tExpected, *ts)
		}
	})

	t.Run("old timestamp", func(t *testing.T) {
		str := "2022-08-12T21:19:37.878526"
		ts, err := ParseBorgTimestamp(str)

		if assert.NoError(t, err) {
			assert.Equal(t, tExpected, *ts)
		}
	})
}
