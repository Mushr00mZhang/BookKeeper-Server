package utils

import "time"

func ParseTime(t string) (*time.Time, error) {
	format := "2006-01-02 15:04:05"
	res, err := time.ParseInLocation(format, t, time.Local)
	if err == nil {
		return &res, nil
	}
	format = "2006-01-02T15:04:05"
	res, err = time.ParseInLocation(format, t, time.Local)
	if err == nil {
		return &res, nil
	}
	format = "2006-01-02T15:04:05+08:00"
	res, err = time.ParseInLocation(format, t, time.Local)
	if err == nil {
		return &res, nil
	}
	return nil, err
}
