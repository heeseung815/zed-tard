package vo

import (
	"strconv"
	"strings"
	"time"
)

type UnixTimeMillis time.Time

func (t UnixTimeMillis) MarshalJSON() ([]byte, error) {
	epoch := time.Time(t)
	tick := epoch.UnixNano() / 1000000
	return []byte(strconv.FormatInt(tick, 10)), nil
}

func (t *UnixTimeMillis) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(q/1000, (q%1000)*int64(time.Millisecond))
	return
}

func (t UnixTimeMillis) Time() time.Time {
	return time.Time(t)
}

func (t UnixTimeMillis) String() string {
	return time.Time(t).Format("15:04:05.000")
}
