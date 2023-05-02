package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		url := "https://github.com/October003?tab=repositories"
		got := Get(url)
		convey.So(got, convey.ShouldEqual, true) // 断言
		convey.ShouldBeTrue(got)
	})
	convey.Convey("url请求不通的示例", t, func() {
		url := "/October003?tab=repositories"
		got := Get(url)
		// 断言
		convey.ShouldBeFalse(got)
	})
}
