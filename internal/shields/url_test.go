package shields

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestURLBuild(t *testing.T) {
	Convey("build URL", t, func() {
		Convey("basic badge", func() {
			So(URL(B{
				Subject: "test platform",
				Status:  "this_is--definitely ok",
				Color:   "brightgreen",
			}), ShouldEqual, "https://img.shields.io/badge/test%20platform-this__is----definitely%20ok-brightgreen.svg")
		})
	})
}
