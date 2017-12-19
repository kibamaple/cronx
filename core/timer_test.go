package core

import (
    "testing"
	time "time"
	convey "github.com/smartystreets/goconvey/convey"
)

func Test_Compute(t *testing.T) {

	convey.Convey("test compute", t, func() {
		
		convey.Convey("error spec",func(c convey.C) {
			spec,now := "adasd",getTime("Mon Jan 2 15:04:08 2017")

			_,err := compute(now,spec,1)
			c.So(err,convey.ShouldNotBeNil)
		})

		convey.Convey("zero offset",func(c convey.C) {
			spec,now := "adasd",getTime("Mon Jan 2 15:04:08 2017")

			actual,err := compute(now,spec,0)
			c.So(err,convey.ShouldBeNil)
			c.So(actual,convey.ShouldEqual,now)
		})
		
		convey.Convey("positive offset",func(c convey.C) {
			spec,now := "0-10/5 4 15 * * *",getTime("Mon Jan 2 15:04:03 2017")
			expecteds := []int64{getTime("Mon Jan 2 15:04:05 2017"),getTime("Mon Jan 2 15:04:10 2017"),getTime("Mon Jan 3 15:04:00 2017")}

			for index,expected := range expecteds {
				actual,err := compute(now,spec,int64(index+1))
				c.So(err,convey.ShouldBeNil)
				c.So(actual,convey.ShouldEqual,expected)
			}
		})

		convey.Convey("negative offset",func(c convey.C) {
			spec,now := "0-10/5 4 15 2 1 *",getTime("Mon Jan 2 15:04:03 2017")
			expecteds := []int64{getTime("Mon Jan 2 15:04:00 2017"),getTime("Mon Jan 2 15:04:10 2016"),getTime("Mon Jan 2 15:04:05 2016")}

			for index,expected := range expecteds {
				actual,err := compute(now,spec,-1 * int64(index+1))
				c.So(err,convey.ShouldBeNil)
				c.So(actual,convey.ShouldEqual,expected)
			}
		})

	})
}

func getTime(value string) int64 {

	if value == "" {
		return 0
	}

	const layout = "Mon Jan 2 15:04:05 2006"

	if _time,err := time.Parse(layout,value); err == nil {
		return _time.Unix()
	}
	
	panic("could not parse time value " + value)
}