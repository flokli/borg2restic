package main

import (
	"fmt"
	"time"
)

type BorgArchive struct {
	Archive  string `json:"archive"`
	BArchive string `json:"barchive"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Start    string `json:"start"`
	Time     string `json:"time"`

	startT time.Time
	timeT  time.Time
}

// ParseBorgTimestamp tries to parse a borg timestamp into a time.Time.
// Both time formats used by borg are supported.
func ParseBorgTimestamp(str string) (*time.Time, error) {
	// borg format seems to be 2022-08-12T21:19:37.878526+02:00
	// (as per https://github.com/borgbackup/borg/pull/6964#issuecomment-1213444009)
	layoutNew := "2006-01-02T15:04:05.000000-07:00"

	// The old format is 2016-06-01T00:00:00.000000
	layoutOld := "2006-01-02T15:04:05.000000"

	// try to parse the new format
	t, err1 := time.ParseInLocation(layoutNew, str, time.Local)
	if err1 != nil {
		// try to parse the old format
		t, err2 := time.ParseInLocation(layoutOld, str, time.Local)
		if err2 != nil {
			return nil, fmt.Errorf("unable to parse in either new or old format: %w", err1)
		}
		return &t, nil
	}
	return &t, nil
}

// ParseTimestamps parses the string timestamps from Start and Time
// into the internal startT and timeT time.Time values,
// and returns an error if there was one during parsing.
func (ba *BorgArchive) ParseTimestamps() error {
	t, err := ParseBorgTimestamp(ba.Start)
	if err != nil {
		return fmt.Errorf("unable to parse Start field: %w", err)
	}
	ba.startT = *t

	t, err = ParseBorgTimestamp(ba.Time)
	if err != nil {
		return fmt.Errorf("unable to parse Time field: %w", err)
	}
	ba.timeT = *t

	return nil
}

func (ba *BorgArchive) GetStartTime() time.Time {
	return ba.startT
}

func (ba *BorgArchive) GetTimeTime() time.Time {
	return ba.timeT
}
