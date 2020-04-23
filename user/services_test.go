package user

import (
	"fmt"
	"testing"
)

// TestGetAllSubs is used to tests getAllSubs
func TestGetAllSubs(t *testing.T) {
	resp := getAllSubs("abhiy13")
	if resp == nil {
		t.Error("Bad Resp\n")
	}
	fmt.Println(resp[:10])
}
