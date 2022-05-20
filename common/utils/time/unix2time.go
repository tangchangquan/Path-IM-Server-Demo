package timeUtils

import "time"

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)

}
