package simple

import "testing"

func TestGroup(t *testing.T) {
	s := Default()
	s.Group("/v1")
}
