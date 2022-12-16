package timex

import (
	"strconv"
	"time"
)

/**
 * $(TimeFormatNumber)
 * @Description:
 * @param tp
 * @return int64
 */
func TimeFormatNumber(t time.Time) int64 {

	format := t.Format(`20060102150405`)
	number, _ := strconv.Atoi(format)
	return int64(number)

}

/**
 * $(TimeJST)
 * @Description:
 * @return time.Time
 */
func TimeJST() time.Time {
	cstSh, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(cstSh)
}


/**
 * $(TimeFormatDate)
 * @Description:
 * @param t
 * @return int32
 */
func TimeFormatDate(t time.Time) int32 {

	format := t.Format(`20060102`)
	date, _ := strconv.Atoi(format)
	return int32(date)

}

/**
 * $(DateToTime)
 * @Description:
 * @param date
 * @return time.Time
 */
func DateToTime(date int64) time.Time {
	t := strconv.Itoa(int(date))
	parse, _ := time.Parse(`20060102150405`, t)
	return parse

}


/**
 * $(StrToTime)
 * @Description:
 * @param date
 * @return time.Time
 */
func StrToTime(date string) time.Time {
	parse, _ := time.Parse(`20060102150405`, date)
	return parse

}

/**
 * $(XslDate)
 * @Description:
 * @param date
 * @return time.Time
 */
func XslDate(date string)  time.Time {

	t := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	day, _ := strconv.Atoi(date)
	return t.Add(time.Duration(day*86400) * time.Second)

}
