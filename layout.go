package now

import "strings"

const DayFmt = "yyyy-MM-dd"
const TimeFmt = "HH:mm:ss"
const DayTimeFmt = DayFmt + " " + TimeFmt
const DayTimeMillisFmt = DayTimeFmt + ".SSS"

func ConvertLayout(layout string) string {
	lo := layout
	lo = strings.Replace(lo, "yyyy", "2006", -1)
	lo = strings.Replace(lo, "yy", "06", -1)
	lo = strings.Replace(lo, "MM", "01", -1)
	lo = strings.Replace(lo, "dd", "02", -1)
	lo = strings.Replace(lo, "HH", "15", -1)
	lo = strings.Replace(lo, "mm", "04", -1)
	lo = strings.Replace(lo, "ss", "05", -1)
	lo = strings.Replace(lo, "SSS", "000", -1)
	return lo
}
