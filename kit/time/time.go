package time

import "time"

func FormatTs(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func FormatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
