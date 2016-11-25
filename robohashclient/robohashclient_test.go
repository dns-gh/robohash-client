package robohashclient

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// go test ...robohashclient -gocheck.vv -test.v -gocheck.f TestNAME
func (s *MySuite) TestURIAssembly(c *C) {
	client := MakeRobohashClient(200, 100, 1, 1)
	uri := client.makeURI("test")
	c.Assert(uri, Equals, "https://robohash.org/test?size=200x100&set=1&bgset=1")
	client2 := MakeRobohashClient(100, 100, 0, 0)
	uri = client2.makeURI("test")
	c.Assert(uri, Equals, "https://robohash.org/test?size=100x100")
}

func (s *MySuite) TestFetch(c *C) {
	client := MakeRobohashClient(200, 100, 1, 1)
	img, err := client.Fetch("test")
	c.Assert(err, IsNil)
	c.Assert(len(img) != 0, Equals, true)
}
