package now_test

import (
	"testing"
	"time"

	"github.com/bingoohuang/now"
	"github.com/stretchr/testify/assert"
)

var (
	layout          = "yyyy-MM-dd HH:mm:ss"
	layoutSSS       = "yyyy-MM-dd HH:mm:ss.SSS"
	format          = "2006-01-02 15:04:05.999999999"
	locationCaracas *time.Location
	locationBerlin  *time.Location
	timeCaracas     time.Time
)

func init() {
	var err error
	if locationCaracas, err = time.LoadLocation("America/Caracas"); err != nil {
		panic(err)
	}

	if locationBerlin, err = time.LoadLocation("Europe/Berlin"); err != nil {
		panic(err)
	}

	timeCaracas = time.Date(2016, 1, 1, 12, 10, 0, 0, locationCaracas)
}

func TestMondayAndSunday(t *testing.T) {
	a := assert.New(t).Equal

	n := time.Date(2013, 11, 19, 17, 51, 49, 123456789, time.UTC)
	n2 := time.Date(2013, 11, 24, 17, 51, 49, 123456789, time.UTC)
	nDst := time.Date(2017, 10, 29, 10, 0, 0, 0, locationBerlin)

	a(now.MakeTime(n).Monday().Format(layout), "2013-11-18 00:00:00", "Monday")
	a(now.MakeTime(n2).Monday().Format(layout), "2013-11-18 00:00:00", "Monday")
	a(now.MakeTime(timeCaracas).Monday().Format(layout), "2015-12-28 00:00:00", "Monday Caracas")
	a(now.MakeTime(nDst).Monday().Format(layout), "2017-10-23 00:00:00", "Monday DST")
	a(now.MakeTime(n).Sunday().Format(layout), "2013-11-24 00:00:00", "Sunday")
	a(now.MakeTime(n2).Sunday().Format(layout), "2013-11-24 00:00:00", "Sunday")
	a(now.MakeTime(timeCaracas).Sunday().Format(layout), "2016-01-03 00:00:00", "Sunday Caracas")
	a(now.MakeTime(nDst).Sunday().Format(layout), "2017-10-29 00:00:00", "Sunday DST")
	a(now.MakeTime(n).EndOfSunday().Format(format), "2013-11-24 23:59:59.999999999", "EndOfSunday")
	a(now.MakeTime(timeCaracas).EndOfSunday().Format(format), "2016-01-03 23:59:59.999999999", "EndOfSunday Caracas")
	a(now.MakeTime(nDst).EndOfSunday().Format(format), "2017-10-29 23:59:59.999999999", "EndOfSunday DST")
	a(now.MakeTime(n).BeginningOfWeek(time.Sunday).Format(layout), "2013-11-17 00:00:00", "BeginningOfWeek, FirstDayMonday")
	a(now.MakeTime(n).BeginningOfWeek(time.Monday).Format(layout), "2013-11-18 00:00:00", "BeginningOfWeek, FirstDayMonday")
}

func TestParse(t *testing.T) {
	a := assert.New(t).Equal

	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)
	a(now.MakeTime(n).MustParseAny("2002").Format(layout), "2002-01-01 00:00:00", "Parse 2002")
	a(now.MakeTime(n).MustParseAny("2002-10").Format(layout), "2002-10-01 00:00:00", "Parse 2002-10")
	a(now.MakeTime(n).MustParseAny("2002-10-12").Format(layout), "2002-10-12 00:00:00", "Parse 2002-10-12")
	a(now.MakeTime(n).MustParseAny("2002-10-12 22").Format(layout), "2002-10-12 22:00:00", "Parse 2002-10-12 22")
	a(now.MakeTime(n).MustParseAny("2002-10-12 22:14").Format(layout), "2002-10-12 22:14:00", "Parse 2002-10-12 22:14")
	a(now.MakeTime(n).MustParseAny("2002-10-12 2:4").Format(layout), "2002-10-12 02:04:00", "Parse 2002-10-12 2:4")
	a(now.MakeTime(n).MustParseAny("2002-10-12 02:04").Format(layout), "2002-10-12 02:04:00", "Parse 2002-10-12 02:04")
	a(now.MakeTime(n).MustParseAny("2002-10-12 22:14:56").Format(layout), "2002-10-12 22:14:56", "Parse 2002-10-12 22:14:56")
	a(now.MakeTime(n).MustParseAny("2002-10-12 00:14:56").Format(layout), "2002-10-12 00:14:56", "Parse 2002-10-12 00:14:56")
	a(now.MakeTime(n).MustParseAny("2013-12-19 23:28:09.999999999 +0800 CST").Format(format), "2013-12-19 23:28:09.999999999", "Parse two strings 2013-12-19 23:28:09.999999999 +0800 CST")
	a(now.MakeTime(n).MustParseAny("10-12").Format(layout), "2013-10-12 00:00:00", "Parse 10-12")
	a(now.MakeTime(n).MustParseAny("18").Format(layout), "2013-11-18 18:00:00", "Parse 18 as hour")
	a(now.MakeTime(n).MustParseAny("18:20").Format(layout), "2013-11-18 18:20:00", "Parse 18:20")
	a(now.MakeTime(n).MustParseAny("00:01").Format(layout), "2013-11-18 00:01:00", "Parse 00:01")
	a(now.MakeTime(n).MustParseAny("00:00:00").Format(layout), "2013-11-18 00:00:00", "Parse 00:00:00")
	a(now.MakeTime(n).MustParseAny("18:20:39").Format(layout), "2013-11-18 18:20:39", "Parse 18:20:39")
	a(now.MakeTime(n).MustParseAny("18:20:39", "2011-01-01").Format(layout), "2011-01-01 18:20:39", "Parse two strings 18:20:39, 2011-01-01")
	a(now.MakeTime(n).MustParseAny("2011-1-1", "18:20:39").Format(layout), "2011-01-01 18:20:39", "Parse two strings 2011-01-01, 18:20:39")
	a(now.MakeTime(n).MustParseAny("2011-01-01", "18").Format(layout), "2011-01-01 18:00:00", "Parse two strings 2011-01-01, 18")
	now.Formats = append(now.Formats, "02 Jan 15:04")
	a(now.MakeTime(n).MustParseAny("04 Feb 12:09").Format(layout), "2013-02-04 12:09:00", "Parse 04 Feb 12:09 with specified format")
	a(now.MakeTime(n).MustParseAny("23:28:9 Dec 19, 2013 PST").Format(layout), "2013-12-19 23:28:09", "Parse 23:28:9 Dec 19, 2013 PST")
	if now.MakeTime(n).MustParseAny("23:28:9 Dec 19, 2013 PST").T.Location().String() != "PST" {
		t.Errorf("Parse 23:28:9 Dec 19, 2013 PST shouldn't lose time zone")
	}

	n2 := now.MakeTime(n).MustParseAny("23:28:9 Dec 19, 2013 PST")
	if now.MakeTime(n2.T).MustParseAny("10:20").T.Location().String() != "PST" {
		t.Errorf("Parse 10:20 shouldn't change time zone")
	}

	now.Formats = append(now.Formats, "2006-01-02T15:04:05.0")
	if now.MustParseAnyInLocation(time.UTC, "2018-02-13T15:17:06.0").T.String() != "2018-02-13 15:17:06 +0000 UTC" {
		t.Errorf("ParseInLocation 2018-02-13T15:17:06.0")
	}

	now.Formats = append(now.Formats, "2006-01-02 15:04:05.000")
	a(now.MakeTime(n).MustParseAny("2018-04-20 21:22:23.473").Format(layoutSSS), "2018-04-20 21:22:23.473", "Parse 2018/04/20 21:22:23.473")

	now.Formats = append(now.Formats, "15:04:05.000")
	a(now.MakeTime(n).MustParseAny("13:00:01.365").Format(layoutSSS), "2013-11-18 13:00:01.365", "Parse 13:00:01.365")

	now.Formats = append(now.Formats, "2006-01-02 15:04:05.000000")
	a(now.MakeTime(n).MustParseAny("2010-01-01 07:24:23.131384").Format(format), "2010-01-01 07:24:23.131384", "Parse 2010-01-01 07:24:23.131384")
	a(now.MakeTime(n).MustParseAny("00:00:00.182736").Format(format), "2013-11-18 00:00:00.182736", "Parse 00:00:00.182736")
}
