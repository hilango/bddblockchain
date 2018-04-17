package main


import(
"testing"
. "github.com/smartystreets/goconvey/convey"
)


func TestSomething(t *testing.T) {
	t.Parallel()
	Convey("1 should equal 1", t, func() {
		So(1, ShouldEqual, 1)
	})

	Convey("Comparing two variables", t, func() {
		myVar := "Hello,world!"

		Convey(`"Asdf" should Not equal "qwrty"`, func() {
			So(myVar, ShouldNotEqual, "qwerty")
		})

		Convey("myVar should not bi nil", func() {
			So(myVar, ShouldNotBeNil)
		})

		Convey("This isn't yet implemented", nil)

	})

}

